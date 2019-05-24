package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchEmbed(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid", func(t *testing.T) {
		resp, err := fetchEmbed(
			"https://www.youtube.com/watch?v=8jPQjjsBbIc",
			&Provider{
				Name: "YouTube",
				URL:  "https://www.youtube.com/",
				Endpoints: []Endpoint{{
					Schemes: []string{
						"https://*.youtube.com/watch*",
						"https://*.youtube.com/v/*\"",
						"https://youtu.be/*",
					},
					URL:       "https://www.youtube.com/oembed",
					Discovery: true,
				}},
			},
			&Params{
				MaxWidth:  250,
				MaxHeight: 250,
			},
		)
		assert.NoError(err)
		assert.NotNil(resp)
	})
	t.Run("invalid", func(t *testing.T) {
		for _, url := range []string{
			"htt:/abc.com/failed-none-sense",
			"https://abc.com/failed-none-sense",
			"http://badcom/146753785",
			"https://674458092126388225",
			"http://www.ted.com/talks/something-does-not-exist",
			"https://soundcloud^(*%%$%^$$%$$*&(&)())",
			"https://www.flickr.com/services/oembed/?url=http%3A//www.flickr.com/photos/bees/23416sa/",
		} {
			url := url
			t.Run(url, func(t *testing.T) {
				provider := findProvider(url)
				if provider == nil {
					provider = &Provider{Endpoints: []Endpoint{Endpoint{}}}
				}

				_, err := fetchEmbed(url, provider, nil)
				assert.Error(err)
			})
		}
	})
}
