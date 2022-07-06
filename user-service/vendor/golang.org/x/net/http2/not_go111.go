// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.11
// +build !go1.11

package http2

import (
	"net/http/httptrace"
	"net/textproto"
)

func traceHasWroteHeaderField(trace *httptrace.UserTrace) bool { return false }

func traceWroteHeaderField(trace *httptrace.UserTrace, k, v string) {}

func traceGot1xxResponseFunc(trace *httptrace.UserTrace) func(int, textproto.MIMEHeader) error {
	return nil
}