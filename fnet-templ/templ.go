package fnet

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
)

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

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func releaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufferPool.Put(b)
}

type Component interface {
	Render(c context.Context, w io.Writer) error
}

// ComponentHandler is a http.Handler that renders components.
type ComponentHandler struct {
	Component      Component
	Status         int
	ContentType    string
	ErrorHandler   func(r *http.Request, err error) http.Handler
	StreamResponse bool
}

const componentHandlerErrorMessage = "templ: failed to render template"

func (ch *ComponentHandler) ServeHTTPBuffered(w http.ResponseWriter, r *http.Request) {
	// Since the component may error, write to a buffer first.
	// This prevents partial responses from being written to the client.
	buf := getBuffer()
	defer releaseBuffer(buf)
	err := ch.Component.Render(r.Context(), buf)
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
	if err := ch.Component.Render(r.Context(), w); err != nil {
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

// Handler creates a http.Handler that renders the template.
func Handler(c Component, options ...func(*ComponentHandler)) *ComponentHandler {
	ch := &ComponentHandler{
		Component:   c,
		ContentType: "text/html; charset=utf-8",
	}
	for _, o := range options {
		o(ch)
	}
	return ch
}

// WithStatus sets the HTTP status code returned by the ComponentHandler.
func WithStatus(status int) func(*ComponentHandler) {
	return func(ch *ComponentHandler) {
		ch.Status = status
	}
}

// WithContentType sets the Content-Type header returned by the ComponentHandler.
func WithContentType(contentType string) func(*ComponentHandler) {
	return func(ch *ComponentHandler) {
		ch.ContentType = contentType
	}
}

// WithErrorHandler sets the error handler used if rendering fails.
func WithErrorHandler(eh func(r *http.Request, err error) http.Handler) func(*ComponentHandler) {
	return func(ch *ComponentHandler) {
		ch.ErrorHandler = eh
	}
}

// WithStreaming sets the ComponentHandler to stream the response instead of buffering it.
func WithStreaming() func(*ComponentHandler) {
	return func(ch *ComponentHandler) {
		ch.StreamResponse = true
	}
}
