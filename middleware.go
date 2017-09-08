package middleware

import (
	"errors"
	"net/http"
)

var MiddlewareStopPropagationError = errors.New("MiddlewareStopPropagationError")

type MiddlewareHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
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
	defer func() {
		if r := recover(); r != nil {
			if r != MiddlewareStopPropagationError {
				panic(r)
			}

			return
		}
	}()

	for _, h := range m.handlers {
		h.ServeHTTP(w, r)
	}
}

func (m *Middleware) AddHandler(h MiddlewareHandler) {
	if m.handlers == nil {
		m.handlers = []MiddlewareHandler{h}
	} else {
		m.handlers = append(m.handlers, h)
	}
}
