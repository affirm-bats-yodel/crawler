package httphandler

import (
	"context"
	"net/http"

	"github.com/affirm-bats-yodel/crawler/pkg/request"
)

// NewHandler Create new HTTP Request handler
//
// agents: you can add name of the agent that appends
// "User-Agent" header
func NewHandler(allowRedirect bool, agents ...string) (*Handler, error) {
	var agent string
	if agents != nil && agents[0] != "" {
		agent = agents[0]
	}
	return &Handler{
		AllowRedirect: allowRedirect,
		Agent:         agent,
	}, nil
}

// Handler HTTP Handler that implementment Request
type Handler struct {
	// AllowRedirect Allow Redirect
	AllowRedirect bool
	// Agent A Name of the agent
	Agent string
}

const (
	headerContentType = "Content-Type"
)

// Get implements request.Request.
func (h *Handler) Get(ctx context.Context, url string) (*request.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if h.Agent != "" {
		req.Header.Set("User-Agent", h.Agent)
	}
	resp, err := h.getHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	return &request.Response{
		StatusCode:    resp.StatusCode,
		ContentLength: resp.ContentLength,
		Header:        &resp.Header,
		Cookies:       resp.Cookies(),
		Body:          resp.Body,
	}, nil
}

// Shutdown implements request.Request.
func (*Handler) Shutdown(_ context.Context) error {
	return nil
}

// getHTTPClient Generate HTTP Client
func (h *Handler) getHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			if h.AllowRedirect {
				return nil
			}
			return http.ErrUseLastResponse
		},
	}
}

var _ request.Request = (*Handler)(nil)
