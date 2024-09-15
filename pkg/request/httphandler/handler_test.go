package httphandler_test

import (
	"bytes"
	"context"
	"flag"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/affirm-bats-yodel/crawler/pkg/request/httphandler"
	"github.com/stretchr/testify/assert"
)

var requestURL *string = flag.String("url", "", "a url to request")

func TestHandler_Get(t *testing.T) {
	if !assert.NotEmpty(t, requestURL, "url should've defined") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	h, err := httphandler.NewHandler(false)
	assert.NoError(t, err)

	t.Logf("request to: %s", *requestURL)
	res, err := h.Get(ctx, *requestURL)
	if !assert.NoError(t, err) {
		return
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	ct, err := res.GetContentType()
	if assert.NoError(t, err) {
		assert.Equal(t, "text/html", ct.MediaType)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, res.ContentLength, "content-length should greator than 0")
	_, err = io.Copy(&buf, res.Body)
	if assert.NoError(t, err) {
		t.Logf("body: %s", buf.Bytes())
		assert.NotEmpty(t, buf.Bytes())
	}
}
