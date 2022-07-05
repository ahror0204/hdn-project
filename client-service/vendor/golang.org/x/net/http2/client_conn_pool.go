// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code's User connection pooling.

package http2

import (
	"crypto/tls"
	"net/http"
	"sync"
)

// UserConnPool manages a pool of HTTP/2 User connections.
type UserConnPool interface {
	GetUserConn(req *http.Request, addr string) (*UserConn, error)
	MarkDead(*UserConn)
}

// UserConnPoolIdleCloser is the interface implemented by UserConnPool
// implementations which can close their idle connections.
type UserConnPoolIdleCloser interface {
	UserConnPool
	closeIdleConnections()
}

var (
	_ UserConnPoolIdleCloser = (*UserConnPool)(nil)
	_ UserConnPoolIdleCloser = noDialUserConnPool{}
)

// TODO: use singleflight for dialing and addConnCalls?
type UserConnPool struct {
	t *Transport

	mu sync.Mutex // TODO: maybe switch to RWMutex
	// TODO: add support for sharing conns based on cert names
	// (e.g. share conn for googleapis.com and appspot.com)
	conns        map[string][]*UserConn // key is host:port
	dialing      map[string]*dialCall   // currently in-flight dials
	keys         map[*UserConn][]string
	addConnCalls map[string]*addConnCall // in-flight addConnIfNeede calls
}

func (p *UserConnPool) GetUserConn(req *http.Request, addr string) (*UserConn, error) {
	return p.getUserConn(req, addr, dialOnMiss)
}

const (
	dialOnMiss   = true
	noDialOnMiss = false
)

// shouldTraceGetConn reports whether getUserConn should call any
// UserTrace.GetConn hook associated with the http.Request.
//
// This complexity is needed to avoid double calls of the GetConn hook
// during the back-and-forth between net/http and x/net/http2 (when the
// net/http.Transport is upgraded to also speak http2), as well as support
// the case where x/net/http2 is being used directly.
func (p *UserConnPool) shouldTraceGetConn(st UserConnIdleState) bool {
	// If our Transport wasn't made via ConfigureTransport, always
	// trace the GetConn hook if provided, because that means the
	// http2 package is being used directly and it's the one
	// dialing, as opposed to net/http.
	if _, ok := p.t.ConnPool.(noDialUserConnPool); !ok {
		return true
	}
	// Otherwise, only use the GetConn hook if this connection has
	// been used previously for other requests. For fresh
	// connections, the net/http package does the dialing.
	return !st.freshConn
}

func (p *UserConnPool) getUserConn(req *http.Request, addr string, dialOnMiss bool) (*UserConn, error) {
	if isConnectionCloseRequest(req) && dialOnMiss {
		// It gets its own connection.
		traceGetConn(req, addr)
		const singleUse = true
		cc, err := p.t.dialUserConn(addr, singleUse)
		if err != nil {
			return nil, err
		}
		return cc, nil
	}
	p.mu.Lock()
	for _, cc := range p.conns[addr] {
		if st := cc.idleState(); st.canTakeNewRequest {
			if p.shouldTraceGetConn(st) {
				traceGetConn(req, addr)
			}
			p.mu.Unlock()
			return cc, nil
		}
	}
	if !dialOnMiss {
		p.mu.Unlock()
		return nil, ErrNoCachedConn
	}
	traceGetConn(req, addr)
	call := p.getStartDialLocked(addr)
	p.mu.Unlock()
	<-call.done
	return call.res, call.err
}

// dialCall is an in-flight Transport dial call to a host.
type dialCall struct {
	_    incomparable
	p    *UserConnPool
	done chan struct{} // closed when done
	res  *UserConn     // valid after done is closed
	err  error         // valid after done is closed
}

// requires p.mu is held.
func (p *UserConnPool) getStartDialLocked(addr string) *dialCall {
	if call, ok := p.dialing[addr]; ok {
		// A dial is already in-flight. Don't start another.
		return call
	}
	call := &dialCall{p: p, done: make(chan struct{})}
	if p.dialing == nil {
		p.dialing = make(map[string]*dialCall)
	}
	p.dialing[addr] = call
	go call.dial(addr)
	return call
}

// run in its own goroutine.
func (c *dialCall) dial(addr string) {
	const singleUse = false // shared conn
	c.res, c.err = c.p.t.dialUserConn(addr, singleUse)
	close(c.done)

	c.p.mu.Lock()
	delete(c.p.dialing, addr)
	if c.err == nil {
		c.p.addConnLocked(addr, c.res)
	}
	c.p.mu.Unlock()
}

// addConnIfNeeded makes a NewUserConn out of c if a connection for key doesn't
// already exist. It coalesces concurrent calls with the same key.
// This is used by the http1 Transport code when it creates a new connection. Because
// the http1 Transport doesn't de-dup TCP dials to outbound hosts (because it doesn't know
// the protocol), it can get into a situation where it has multiple TLS connections.
// This code decides which ones live or die.
// The return value used is whether c was used.
// c is never closed.
func (p *UserConnPool) addConnIfNeeded(key string, t *Transport, c *tls.Conn) (used bool, err error) {
	p.mu.Lock()
	for _, cc := range p.conns[key] {
		if cc.CanTakeNewRequest() {
			p.mu.Unlock()
			return false, nil
		}
	}
	call, dup := p.addConnCalls[key]
	if !dup {
		if p.addConnCalls == nil {
			p.addConnCalls = make(map[string]*addConnCall)
		}
		call = &addConnCall{
			p:    p,
			done: make(chan struct{}),
		}
		p.addConnCalls[key] = call
		go call.run(t, key, c)
	}
	p.mu.Unlock()

	<-call.done
	if call.err != nil {
		return false, call.err
	}
	return !dup, nil
}

type addConnCall struct {
	_    incomparable
	p    *UserConnPool
	done chan struct{} // closed when done
	err  error
}

func (c *addConnCall) run(t *Transport, key string, tc *tls.Conn) {
	cc, err := t.NewUserConn(tc)

	p := c.p
	p.mu.Lock()
	if err != nil {
		c.err = err
	} else {
		p.addConnLocked(key, cc)
	}
	delete(p.addConnCalls, key)
	p.mu.Unlock()
	close(c.done)
}

// p.mu must be held
func (p *UserConnPool) addConnLocked(key string, cc *UserConn) {
	for _, v := range p.conns[key] {
		if v == cc {
			return
		}
	}
	if p.conns == nil {
		p.conns = make(map[string][]*UserConn)
	}
	if p.keys == nil {
		p.keys = make(map[*UserConn][]string)
	}
	p.conns[key] = append(p.conns[key], cc)
	p.keys[cc] = append(p.keys[cc], key)
}

func (p *UserConnPool) MarkDead(cc *UserConn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, key := range p.keys[cc] {
		vv, ok := p.conns[key]
		if !ok {
			continue
		}
		newList := filterOutUserConn(vv, cc)
		if len(newList) > 0 {
			p.conns[key] = newList
		} else {
			delete(p.conns, key)
		}
	}
	delete(p.keys, cc)
}

func (p *UserConnPool) closeIdleConnections() {
	p.mu.Lock()
	defer p.mu.Unlock()
	// TODO: don't close a cc if it was just added to the pool
	// milliseconds ago and has never been used. There's currently
	// a small race window with the HTTP/1 Transport's integration
	// where it can add an idle conn just before using it, and
	// somebody else can concurrently call CloseIdleConns and
	// break some caller's RoundTrip.
	for _, vv := range p.conns {
		for _, cc := range vv {
			cc.closeIfIdle()
		}
	}
}

func filterOutUserConn(in []*UserConn, exclude *UserConn) []*UserConn {
	out := in[:0]
	for _, v := range in {
		if v != exclude {
			out = append(out, v)
		}
	}
	// If we filtered it out, zero out the last item to prevent
	// the GC from seeing it.
	if len(in) != len(out) {
		in[len(in)-1] = nil
	}
	return out
}

// noDialUserConnPool is an implementation of http2.UserConnPool
// which never dials. We let the HTTP/1.1 User dial and use its TLS
// connection instead.
type noDialUserConnPool struct{ *UserConnPool }

func (p noDialUserConnPool) GetUserConn(req *http.Request, addr string) (*UserConn, error) {
	return p.getUserConn(req, addr, noDialOnMiss)
}
