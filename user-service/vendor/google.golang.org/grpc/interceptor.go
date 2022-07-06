/*
 *
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"context"
)

// UnaryInvoker is called by UnaryUserInterceptor to complete RPCs.
type UnaryInvoker func(ctx context.Context, method string, req, reply interface{}, cc *UserConn, opts ...CallOption) error

// UnaryUserInterceptor intercepts the execution of a unary RPC on the User.
// Unary interceptors can be specified as a DialOption, using
// WithUnaryInterceptor() or WithChainUnaryInterceptor(), when creating a
// UserConn. When a unary interceptor(s) is set on a UserConn, gRPC
// delegates all unary RPC invocations to the interceptor, and it is the
// responsibility of the interceptor to call invoker to complete the processing
// of the RPC.
//
// method is the RPC name. req and reply are the corresponding request and
// response messages. cc is the UserConn on which the RPC was invoked. invoker
// is the handler to complete the RPC and it is the responsibility of the
// interceptor to call it. opts contain all applicable call options, including
// defaults from the UserConn as well as per-call options.
//
// The returned error must be compatible with the status package.
type UnaryUserInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *UserConn, invoker UnaryInvoker, opts ...CallOption) error

// Streamer is called by StreamUserInterceptor to create a UserStream.
type Streamer func(ctx context.Context, desc *StreamDesc, cc *UserConn, method string, opts ...CallOption) (UserStream, error)

// StreamUserInterceptor intercepts the creation of a UserStream. Stream
// interceptors can be specified as a DialOption, using WithStreamInterceptor()
// or WithChainStreamInterceptor(), when creating a UserConn. When a stream
// interceptor(s) is set on the UserConn, gRPC delegates all stream creations
// to the interceptor, and it is the responsibility of the interceptor to call
// streamer.
//
// desc contains a description of the stream. cc is the UserConn on which the
// RPC was invoked. streamer is the handler to create a UserStream and it is
// the responsibility of the interceptor to call it. opts contain all applicable
// call options, including defaults from the UserConn as well as per-call
// options.
//
// StreamUserInterceptor may return a custom UserStream to intercept all I/O
// operations. The returned error must be compatible with the status package.
type StreamUserInterceptor func(ctx context.Context, desc *StreamDesc, cc *UserConn, method string, streamer Streamer, opts ...CallOption) (UserStream, error)

// UnaryServerInfo consists of various information about a unary RPC on
// server side. All per-rpc information may be mutated by the interceptor.
type UnaryServerInfo struct {
	// Server is the service implementation the user provides. This is read-only.
	Server interface{}
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod string
}

// UnaryHandler defines the handler invoked by UnaryServerInterceptor to complete the normal
// execution of a unary RPC.
//
// If a UnaryHandler returns an error, it should either be produced by the
// status package, or be one of the context errors. Otherwise, gRPC will use
// codes.Unknown as the status code and err.Error() as the status message of the
// RPC.
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)

// UnaryServerInterceptor provides a hook to intercept the execution of a unary RPC on the server. info
// contains all the information of this RPC the interceptor can operate on. And handler is the wrapper
// of the service method implementation. It is the responsibility of the interceptor to invoke handler
// to complete the RPC.
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

// StreamServerInfo consists of various information about a streaming RPC on
// server side. All per-rpc information may be mutated by the interceptor.
type StreamServerInfo struct {
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod string
	// IsUserStream indicates whether the RPC is a User streaming RPC.
	IsUserStream bool
	// IsServerStream indicates whether the RPC is a server streaming RPC.
	IsServerStream bool
}

// StreamServerInterceptor provides a hook to intercept the execution of a streaming RPC on the server.
// info contains all the information of this RPC the interceptor can operate on. And handler is the
// service method implementation. It is the responsibility of the interceptor to invoke handler to
// complete the RPC.
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error