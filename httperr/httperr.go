// Copyright 2015 Tamás Gulácsi. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

// Package httperr provides a small helper for returning statuscode-accompanied errors.
package httperr

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var _ = error(httpError{})

type httpError struct {
	error
	code int
}

// Code returns the error's accompanied StatusCode.
func (he httpError) StatusCode() int {
	return he.code
}
func (he httpError) Cause() error {
	return he.error
}

// New returns a new error with the specified error and StatusCode.
// If code is zero, http.StatusInternalServerError is used.
// If err is nil, a new error is created from the status code.
func New(err error, code int) *httpError {
	if code == 0 {
		if he, ok := err.(*httpError); ok {
			return he
		}
		if he, ok := err.(interface {
			StatusCode() int
		}); ok {
			code = he.StatusCode()
		} else {
			code = http.StatusInternalServerError
		}
	}
	if err == nil {
		err = errors.New(http.StatusText(code))
	}
	return &httpError{error: err, code: code}
}

// Newf returns a new httpError with the given code,
// and the message is created from msg and the args.
func Newf(code int, msg string, args ...interface{}) *httpError {
	return New(errors.New(fmt.Sprintf(msg, args)), code)
}

// vim: fileencoding=utf-8:
