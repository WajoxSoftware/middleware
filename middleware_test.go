package middleware

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "bytes"
)

var handlerStack []string

type middlewareExample struct {
	Name string
	ResultCode bool
}

func (m middlewareExample) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
    handlerStack = append(handlerStack, m.Name)

    return m.ResultCode
}

func TestCreateNewMiddleware(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	mware := CreateNewMiddleware()

	if mware == nil {
		t.Error("Constructor returned nil")
	}
}

func TestAddHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()
	
	handlerStack = []string{}

	mware := CreateNewMiddleware()

	mware.AddHandler(&middlewareExample{"First", true})
	mware.AddHandler(&middlewareExample{"Second", true})
	mware.AddHandler(&middlewareExample{"Third", false})
	mware.AddHandler(&middlewareExample{"Last", true})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8000/apply.events.new_user/events", bytes.NewBufferString(""))
	mware.ServeHTTP(w, r)

	if len(handlerStack) != 3 {
    	t.Error("Incorrect number of used handlers")
	}

	if handlerStack[len(handlerStack) - 1] != "Third" {
		t.Error("Incorrect order")
	}

	if handlerStack[0] != "First" {
		t.Error("Incorrect order")
	}

	if handlerStack[1] != "Second" {
		t.Error("Incorrect order")
	}

	if handlerStack[2] != "Third" {
		t.Error("Incorrect order")
	}
}
