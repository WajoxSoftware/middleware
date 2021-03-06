Lightweight middleware package
========================

[![Build Status](https://travis-ci.org/WajoxSoftware/middleware.svg?branch=master)](https://travis-ci.org/WajoxSoftware/middleware)
[![Maintainability](https://api.codeclimate.com/v1/badges/52f98da053fe04d96e6c/maintainability)](https://codeclimate.com/github/WajoxSoftware/middleware/maintainability)

This is a Golang package, that implements middleware pattern.
Just 37 lines of code.

## Installation

```
# run this command
go get github.com/WajoxSoftware/middleware

# import package
import (
    // ... some packages
    "github.com/WajoxSoftware/middleware"
    // ... some packages
)

```
## Usage example

```
/**
 * router and auth implements MiddlewareHandler interface
 * 
 */

package main

import (
    // ... some packages
    "github.com/WajoxSoftware/middleware"
)

// ... some code

mware := middleware.CreateNewMiddleware()
mware.AddHandler(auth)
mware.AddHandler(router)

// ... some code

```

## Definition of MiddlewareHandler interface
```
import (
    "net/http"
)

type MiddlewareHandler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

## MiddlewareHandler example
```
import (
    "net/http"
)
// ...
type Auth struct {
    AuthToken string
}

func CreateNewAuth(authToken string) *Auth {
    return &Auth{authToken}
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
    isValidToken := a.AuthToken == r.URL.Query().Get("token")

    if (!isValidToken) {
        w.WriteHeader(403)

        return false
    }

    return true
}
// ...

```
