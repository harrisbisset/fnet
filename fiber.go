// MIT License

// Copyright (c) 2021 Adrian Hesketh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package fnet

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// ComponentHandler is a http.Handler that renders components.
type ComponentHandler struct {
	Response
	Status         int
	ContentType    string
	ErrorHandler   func(r *http.Request, err error) http.Handler
	StreamResponse bool
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

const componentHandlerErrorMessage = "failed to render template"

func renderHandler(resp Response, c *fiber.Ctx) error {
	handler := adaptor.HTTPHandler(&ComponentHandler{
		Response:    resp,
		ContentType: "text/html; charset=utf-8",
	})
	return handler(c)
}

func ReleaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufferPool.Put(b)
}

func (ch *ComponentHandler) ServeHTTPBuffered(w http.ResponseWriter, r *http.Request) {
	// Since the component may error, write to a buffer first.
	// This prevents partial responses from being written to the client.
	buf := bufferPool.Get().(*bytes.Buffer)
	defer ReleaseBuffer(buf)
	err := ch.Response.Render(r.Context(), buf)
	if err != nil {
		if ch.ErrorHandler != nil {
			w.Header().Set("Content-Type", ch.ContentType)
			ch.ErrorHandler(r, err).ServeHTTP(w, r)
			return
		}
		http.Error(w, componentHandlerErrorMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ch.ContentType)
	if ch.Status != 0 {
		w.WriteHeader(ch.Status)
	}
	// Ignore write error like http.Error() does, because there is
	// no way to recover at this point.
	_, _ = w.Write(buf.Bytes())
}

func (ch *ComponentHandler) ServeHTTPStreamed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ch.ContentType)
	if ch.Status != 0 {
		w.WriteHeader(ch.Status)
	}
	if err := ch.Response.Render(r.Context(), w); err != nil {
		if ch.ErrorHandler != nil {
			w.Header().Set("Content-Type", ch.ContentType)
			ch.ErrorHandler(r, err).ServeHTTP(w, r)
			return
		}
		http.Error(w, componentHandlerErrorMessage, http.StatusInternalServerError)
	}
}

// ServeHTTP implements the http.Handler interface.
func (ch ComponentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ch.StreamResponse {
		ch.ServeHTTPStreamed(w, r)
		return
	}
	ch.ServeHTTPBuffered(w, r)
}
