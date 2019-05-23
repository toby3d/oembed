package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid", func(t *testing.T) {
		resp, err := Extract("https://www.youtube.com/watch?v=8jPQjjsBbIc", &Params{
			MaxWidth:  250,
			MaxHeight: 250,
		})
		assert.NoError(err)
		assert.NotNil(resp)
	})
	t.Run("invalid", func(t *testing.T) {
		for _, url := range []string{
			"",
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
				_, err := Extract(url, nil)
				assert.Error(err)
			})
		}
	})
}

func TestHasProvider(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert.True(t, HasProvider("https://www.youtube.com/watch?v=8jPQjjsBbIc"))
	})
	t.Run("false", func(t *testing.T) {
		assert.False(t, HasProvider("https://blog.toby3d.me/"))
	})
}
