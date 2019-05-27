// Package oembed add utils for supporting oEmbed fetching data,
package oembed

import "golang.org/x/xerrors"

// Extract try fetch oEmbed object for input url with params (if represent).
// Return OEmbed if success.
func Extract(url string, params *Params) (*OEmbed, error) {
	if !isValidURL(url) {
		return nil, Error{
			Message: "invalid input url",
			URL:     url,
		}
	}
	if provider := findProvider(url); provider != nil {
		resp, err := fetchEmbed(url, provider, params)
		if err != nil {
			return nil, Error{
				Message: err.Error(),
				URL:     url,
				Details: xerrors.Caller(1),
			}
		}
		return resp, nil
	}

	return nil, Error{
		Message: "no provider found with given url",
		URL:     url,
	}
}

// HasProvider checks what input url has oEmbed provider
func HasProvider(url string) bool {
	return findProvider(url) != nil
}
