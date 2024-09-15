package request

import (
	"context"
	"io"
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

// GetContentType Get "Content-Type" from Header
func (r *Response) GetContentType() string {
	return r.Header.Get("Content-Type")
}
