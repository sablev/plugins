// Code generated by goa v2.2.6, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/calc

package client

import (
	"fmt"
	"strconv"

	calc "goa.design/plugins/goakit/examples/calc/gen/calc"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddA string, calcAddB string) (*calc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v, nil
}
