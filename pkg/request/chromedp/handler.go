package chromedp

import (
	"context"
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

// Handler Playwright Handler to support RIA
// (Rich Internet Application)
type Handler struct {
	// Headless Turn Headless On and Off
	Headless bool
}

// Get implements request.Request.
//
// chromedp handler does not return any additonal data
// except ContentLength and Body
func (h *Handler) Get(ctx context.Context, url string) (*request.Response, error) {
	var body *strings.Reader

	ctx, cancel := chromedp.NewExecAllocator(ctx, append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", h.Headless),
	)...)
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

var _ request.Request = (*Handler)(nil)
