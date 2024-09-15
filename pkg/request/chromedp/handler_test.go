package chromedp_test

import (
	"context"
	"testing"
	"time"

	"github.com/affirm-bats-yodel/crawler/pkg/request/chromedp"
)

func TestNewHandler(t *testing.T) {
	h, err := chromedp.NewHandler()
	if err != nil && h == nil {
		t.Error(err)
		return
	}
	defer h.Shutdown(context.Background())

	time.Sleep(time.Second * 2)
}
