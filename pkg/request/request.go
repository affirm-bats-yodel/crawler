package request

import (
	"context"
	"errors"
	"io"
	"mime"
	"net/http"
)

// Request an interface for Handle Request
type Request interface {
	// Get request to get Page
	Get(ctx context.Context, url string) (*Response, error)
	// Shutdown Gracefully Shutdown Externally managed resources
	Shutdown(ctx context.Context) error
}

type Response struct {
	// StatusCode HTTP Status Code
	StatusCode int
	// ContentLength Size of the Content
	ContentLength int64
	// Header Response Header
	Header *http.Header
	// Cookies List of the Cookies
	Cookies []*http.Cookie
	// Body a Body
	Body io.ReadCloser
}

// ContentType a Wrapper from mime.ParseMediaType
// result
type ContentType struct {
	MediaType string
	Params    map[string]string
}

// ErrEmptyContentType Error when "Content-Type" does not exist
// on response header.
var ErrEmptyContentType = errors.New("error: empty Content-Type on Header")

// GetContentType Parse "Content-Type" from Header
func (r *Response) GetContentType() (*ContentType, error) {
	ct := r.Header.Get("Content-Type")
	if ct == "" {
		return nil, ErrEmptyContentType
	}
	mediaType, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return nil, err
	}
	return &ContentType{
		MediaType: mediaType,
		Params:    params,
	}, nil
}
