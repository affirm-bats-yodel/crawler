package playwright

import (
	"context"
	"errors"
	"fmt"

	"github.com/affirm-bats-yodel/crawler/pkg/request"
	"github.com/playwright-community/playwright-go"
)

// NewHandler Create a new Playwright Handler
func NewHandler(bt BrowserType) (*Handler, error) {
	var br playwright.Browser
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	switch bt {
	case Chromium:
		br, err = pw.Chromium.Launch()
	case Firefox:
		br, err = pw.Firefox.Launch()
	case Webkit:
		br, err = pw.WebKit.Launch()
	default:
		err = fmt.Errorf("error: undefined browser type: %d", bt)
	}
	return &Handler{
		Playwright: pw,
		Browser:    br,
	}, err
}

// Handler Playwright Handler to support RIA
// (Rich Internet Application)
type Handler struct {
	Playwright *playwright.Playwright
	Browser    playwright.Browser
}

// Get implements request.Request.
func (h *Handler) Get(ctx context.Context, url string) (*request.Response, error) {
	panic("unimplemented")
}

// Shutdown implements request.Request.
func (h *Handler) Shutdown(ctx context.Context) error {
	var allErrs []error
	// Stop the Browser if defined
	if h.Browser != nil {
		if err := h.Browser.Close(); err != nil {
			allErrs = append(allErrs, err)
		}
	}
	// then stop playwright session
	if err := h.Playwright.Stop(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) > 0 {
		return errors.Join(allErrs...)
	}
	return nil
}

var _ request.Request = (*Handler)(nil)

type BrowserType int

const (
	Chromium BrowserType = iota
	Firefox
	Webkit
)
