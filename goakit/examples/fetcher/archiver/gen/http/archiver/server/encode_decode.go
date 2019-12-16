// Code generated by goa v3.0.9, DO NOT EDIT.
//
// archiver HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/archiver/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/archiver

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	archiverviews "goa.design/plugins/v3/goakit/examples/fetcher/archiver/gen/archiver/views"
)

// EncodeArchiveResponse returns an encoder for responses returned by the
// archiver archive endpoint.
func EncodeArchiveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archiverviews.ArchiveMedia)
		enc := encoder(ctx, w)
		body := NewArchiveResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeArchiveRequest returns a decoder for requests sent to the archiver
// archive endpoint.
func DecodeArchiveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ArchiveRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateArchiveRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewArchivePayload(&body)

		return payload, nil
	}
}

// EncodeReadResponse returns an encoder for responses returned by the archiver
// read endpoint.
func EncodeReadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archiverviews.ArchiveMedia)
		enc := encoder(ctx, w)
		body := NewReadResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeReadRequest returns a decoder for requests sent to the archiver read
// endpoint.
func DecodeReadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if id < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("id", id, 0, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewReadPayload(id)

		return payload, nil
	}
}

// EncodeReadError returns an encoder for errors returned by the read archiver
// endpoint.
func EncodeReadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewReadNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewReadBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
