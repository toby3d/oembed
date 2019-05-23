package oembed

import (
	"golang.org/x/xerrors"
)

var (
	ErrInvalidInputURL = xerrors.New("invalid input url")
	ErrNoProviderFound = xerrors.New("no provider found with given url")
)

func Extract(url string, params *Params) (*Response, error) {
	if !isValidURL(url) {
		return nil, ErrInvalidInputURL
	}
	if c := findProvider(url); c != nil {
		return fetchEmbed(url, *c, params)
	}
	return nil, ErrNoProviderFound
}

func HasProvider(url string) bool {
	return findProvider(url) != nil
}
