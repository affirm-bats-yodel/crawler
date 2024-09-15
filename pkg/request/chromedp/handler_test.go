package chromedp_test

import (
	"bytes"
	"context"
	"flag"
	"io"
	"testing"
	"time"

	"github.com/affirm-bats-yodel/crawler/pkg/request/chromedp"
	"github.com/stretchr/testify/assert"
)

var requestURL *string = flag.String("url", "", "a url to request")

func TestNewHandler(t *testing.T) {
	h, err := chromedp.NewHandler(true)
	if !assert.NoError(t, err) {
		return
	}
	defer h.Shutdown(context.Background())

	time.Sleep(time.Second * 2)
}

func TestHandler_Get(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if !assert.NotEmpty(t, requestURL, "--url should not empty") {
		return
	}

	h, err := chromedp.NewHandler(false)
	if !assert.NoError(t, err) {
		return
	}
	defer h.Shutdown(context.Background())

	res, err := h.Get(ctx, *requestURL)
	if assert.NoError(t, err) {
		var buf bytes.Buffer
		assert.NotEmpty(t, res.ContentLength)
		_, err := io.Copy(&buf, res.Body)
		if assert.NoError(t, err) {
			t.Logf("body: %s", buf.Bytes())
			assert.NotEmpty(t, buf.String())
		}
	}
}
