package request

import "context"

// Request an interface for Handle Request
type Request interface {
	// Get request to get Page
	//
	// return a HTML Content
	Get(ctx context.Context, url string) ([]byte, error)
	// Shutdown Gracefully Shutdown Externally managed resources
	Shutdown(ctx context.Context) error
}
