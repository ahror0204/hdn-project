/*
 * Copyright 2021 gRPC authors.
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
 */

package credentials

import (
	"context"
)

// requestInfoKey is a struct to be used as the key to store RequestInfo in a
// context.
type requestInfoKey struct{}

// NewRequestInfoContext creates a context with ri.
func NewRequestInfoContext(ctx context.Context, ri interface{}) context.Context {
	return context.WithValue(ctx, requestInfoKey{}, ri)
}

// RequestInfoFromContext extracts the RequestInfo from ctx.
func RequestInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(requestInfoKey{})
}

// UserHandshakeInfoKey is a struct used as the key to store
// UserHandshakeInfo in a context.
type UserHandshakeInfoKey struct{}

// UserHandshakeInfoFromContext extracts the UserHandshakeInfo from ctx.
func UserHandshakeInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(UserHandshakeInfoKey{})
}

// NewUserHandshakeInfoContext creates a context with chi.
func NewUserHandshakeInfoContext(ctx context.Context, chi interface{}) context.Context {
	return context.WithValue(ctx, UserHandshakeInfoKey{}, chi)
}
