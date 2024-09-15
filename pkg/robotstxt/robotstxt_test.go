package robotstxt_test

import (
	"fmt"
	"testing"

	"github.com/affirm-bats-yodel/crawler/pkg/robotstxt"
	"github.com/stretchr/testify/assert"
)

// robotsTxtSampleDuckDuckGo a robots.txt from duckduckgo.com
//
// https://duckduckgo.com/robots.txt
const robotsTxtSampleDuckDuckGo = `
User-agent: *
Disallow: /lite
Disallow: /html

# No search result pages
Disallow: /*?

# chrome new tab page
Disallow: /chrome_newtab

# Email Protection
Disallow: /email/
Allow: /email/$
Allow: /email/privacy-guarantees
Allow: /email/privacy-terms

# Old Privacy page
Disallow: /2012-privacy-policy

User-agent: ia_archiver
Disallow: /

Sitemap: https://duckduckgo.com/sitemap.xml
`

// TestParser Test "robots.txt" Parser using
// duckduckgo.com robots.txt
func TestParserDuckDuckGo(t *testing.T) {
	p, err := robotstxt.NewParser(robotsTxtSampleDuckDuckGo)
	assert.NoError(t, err)

	t.Run("Allowed", checkAllowed(p))
	t.Run("ListSitemaps", listSiteMaps(p))
}

// TestParserDuckDuckGoIaArchiver Test "robots.txt" Parser using
// duckduckgo.com robots.txt and set agent as "ia_archiver"
func TestParserDuckDuckGoIaArchiver(t *testing.T) {
	p, err := robotstxt.NewParser(robotsTxtSampleDuckDuckGo, "ia_archiver")
	assert.NoError(t, err)

	t.Run("Allowed", checkAllowed(p, true))
}

// checkAllowed Test list of endpoints are allowed or not
func checkAllowed(p *robotstxt.Parser, allBlocks ...bool) func(t *testing.T) {
	return func(t *testing.T) {
		var b bool
		if allBlocks != nil && allBlocks[0] {
			b = true
		}
		endpointList := []struct {
			path    string
			allowed bool
		}{
			{
				path:    "/",
				allowed: true,
			},
			{
				path:    "/lite",
				allowed: false,
			},
			{
				path:    "/html",
				allowed: false,
			},
			{
				path:    "/?hps=1&q=hello&atb=v447-1&ia=web",
				allowed: false,
			},
			{
				path:    "/chrome_newtab",
				allowed: false,
			},
			{
				path:    "/email",
				allowed: true,
			},
			{
				path:    "/email/privacy-guarantees",
				allowed: true,
			},
			{
				path:    "/email/privacy-terms",
				allowed: true,
			},
		}

		for _, elem := range endpointList {
			if b {
				elem.allowed = false
			}
			t.Run(fmt.Sprintf("%q", elem.path), func(t *testing.T) {
				assert.Equal(t, elem.allowed, p.Allowed(elem.path))
			})
		}
	}
}

// listSiteMaps List Sitemaps via GetSitemaps and check
// it return a correct list of sitemap
func listSiteMaps(p *robotstxt.Parser) func(t *testing.T) {
	return func(t *testing.T) {
		assert.Equal(t, p.GetSitemaps(), []string{"https://duckduckgo.com/sitemap.xml"})
	}
}
