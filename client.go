// Copyright (c) 2017 Jared Patrick <jared.patrick@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package jrpc2

import (
	"net/http"
)

// ClientOptions options that used as configure JSON-RPC web client.
type ClientOptions struct {
	Headers    map[string]string
	HTTPClient *http.Client
}

// Client represents a JSON-RPC 2.0 capable web client.
type Client interface {
	// Call sends a JSON-RPC request to the server endpoint.
	Call(req *RequestObject) (ResponseObject, error)
	// CallBatch sends a list of RPCRequests in a single batch request
	CallBatch(reqs []*RequestObject) ([]*ResponseObject, error)
}
