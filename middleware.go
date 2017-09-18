package middleware

import (
	"errors"
	"net/http"
)

type MiddlewareHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request) bool
}

type Middleware struct {
	handlers []MiddlewareHandler
}

func StopPropagation() {
	panic(MiddlewareStopPropagationError)
}

func CreateNewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, h := range m.handlers {
		if !h.ServeHTTP(w, r) {
			return
		}
	}
}

func (m *Middleware) AddHandler(h MiddlewareHandler) {
	if m.handlers == nil {
		m.handlers = []MiddlewareHandler{h}
	} else {
		m.handlers = append(m.handlers, h)
	}
}
