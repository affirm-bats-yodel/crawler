package chromedp

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/affirm-bats-yodel/crawler/pkg/request"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

// NewHandler Create a new Playwright Handler
func NewHandler(headless bool) (*Handler, error) {
	return &Handler{
		Headless: headless,
	}, nil
}

// NewRemoteHandler Create a new Remote Allocator based Handler
func NewRemoteHandler(remoteAllocatorAddr string) (*Handler, error) {
	s := strings.TrimSpace(remoteAllocatorAddr)
	if len(s) == 0 {
		return nil, errors.New("error: empty remoteAllocatorAddr")
	}
	return &Handler{
		UseRemoteAllocator:  true,
		RemoteAllocatorAddr: s,
	}, nil
}

// Handler Chromedp Handler to support RIA
// (Rich Internet Application)
//
// Chromedp handler require to install chromium
// or google chrome
//
// https://github.com/chromedp/chromedp
type Handler struct {
	// Headless Turn Headless On and Off
	Headless bool
	// UseRemoteAllocator Use Remote allocator
	// rather than execute chromium
	UseRemoteAllocator bool
	// RemoteAllocatorAddr Remote Allocator Address
	//
	// Required if UseRemoteAllocator is set to true
	RemoteAllocatorAddr string
}

// Get implements request.Request.
//
// chromedp handler does not return any additonal data
// except ContentLength and Body
func (h *Handler) Get(ctx context.Context, url string) (*request.Response, error) {
	var body *strings.Reader

	ctx, cancel := h.GetAllocator(ctx)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			body = strings.NewReader(res)
			return nil
		}),
	)

	if err != nil {
		return nil, err
	}

	return &request.Response{
		ContentLength: body.Size(),
		Body:          io.NopCloser(body),
	}, nil
}

// Shutdown implements request.Request.
func (h *Handler) Shutdown(_ context.Context) error {
	return nil
}

// GetAllocator Generate Allocator by UseRemoteAllocator is true or not
//
// NOTE: It does not check RemoteAllocatorAddr is Empty or not, use with care
func (h *Handler) GetAllocator(ctx context.Context) (context.Context, context.CancelFunc) {
	if h.UseRemoteAllocator {
		return chromedp.NewRemoteAllocator(ctx, h.RemoteAllocatorAddr)
	}
	return chromedp.NewExecAllocator(ctx, append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", h.Headless),
	)...)
}

var _ request.Request = (*Handler)(nil)
