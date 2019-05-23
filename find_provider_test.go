package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostname(t *testing.T) {
	assert := assert.New(t)
	for k, v := range map[string]string{
		"https://mais.uol.com.br/": "mais.uol.com.br",
		"http://www.wootled.com/":  "wootled.com",
		"http://yfrog.com":         "yfrog.com",
		"https://www.youtube.com":  "youtube.com",
		"https://www.znipe.tv":     "znipe.tv",
		"http://":                  "",
		"":                         "",
	} {
		t.Run(k, func(t *testing.T) { assert.Equal(v, getHostname(k)) })
	}
}

func TestMakeCandidate(t *testing.T) {
	assert.NotNil(t, makeCandidate(Provider{
		Name: "YouTube",
		URL:  "https://www.youtube.com/",
		Endpoints: []Endpoint{
			Endpoint{
				Schemes: []string{
					"https://*.youtube.com/watch*",
					"https://*.youtube.com/v/*\"",
					"https://youtu.be/*",
				},
				URL:       "https://www.youtube.com/oembed",
				Discovery: true,
			},
		},
	}))
}

func TestFindProvider(t *testing.T) {
	assert.NotNil(t, findProvider("https://www.beautiful.ai/"))
}
