// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/calc

package server

import (
	calc "goa.design/plugins/v3/goakit/examples/calc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	return &calc.AddPayload{
		A: a,
		B: b,
	}
}
