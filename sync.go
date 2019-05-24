package oembed

import (
	"log"
	"time"

	"github.com/gobuffalo/packr"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SourceURL is a official url of supported providers list
const SourceURL string = "https://oembed.com/providers.json"

// Providers contains all default (or new synced) providers
var Providers []Provider //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	if err := Sync(SourceURL); err != nil {
		panic(err)
	}
}

// Sync try update Providers variable via request and parsing body of sourceURL
func Sync(sourceURL string) error {
	status, src, err := http.GetTimeout(nil, sourceURL, 2*time.Second)
	if err != nil || status != http.StatusOK {
		if src, err = packr.NewBox("./assets").Find("providers.json"); err != nil {
			return Error{
				Message: err.Error(),
				URL:     sourceURL,
			}
		}
	}

	if err = json.Unmarshal(src, &Providers); err != nil {
		return Error{
			Message: err.Error(),
			URL:     sourceURL,
		}
	}

	log.Println("providers.json has been updated")
	return nil
}
