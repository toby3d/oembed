package oembed

import "net/url"

func isValidURL(src string) bool {
	_, err := url.ParseRequestURI(src)
	return err == nil
}
