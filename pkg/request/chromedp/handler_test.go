package playwright_test

import (
	"context"
	"testing"
	"time"

	"github.com/affirm-bats-yodel/crawler/pkg/request/playwright"
)

func TestNewHandler(t *testing.T) {
	t.Logf("install browser")
	if err := playwright.Install(playwright.Firefox); err != nil {
		t.Error(err)
		return
	}

	h, err := playwright.NewHandler(playwright.Firefox)
	if err != nil && h == nil {
		t.Error(err)
		return
	}
	defer h.Shutdown(context.Background())

	time.Sleep(time.Second * 2)
}
