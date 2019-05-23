package oembed

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const source string = "https://oembed.com/providers.json"

var (
	providersList []Provider

	target = filepath.Join(
		os.Getenv("GOPATH"), "src", "gitlab.com", "toby3d", "oembed", "assets", "providers.json",
	)
)

func init() {
	if err := fetch(source, target); err != nil {
		panic(err)
	}
}

func fetch(url, target string) error {
	status, src, err := http.Get(nil, url)
	if err != nil || status != http.StatusOK {
		if src, err = ioutil.ReadFile(target); err != nil {
			return err
		}
	}

	if err := json.Unmarshal(src, &providersList); err != nil {
		return err
	}

	if status == http.StatusOK {
		if err = ioutil.WriteFile(target, src, os.ModePerm); err != nil {
			return err
		}
	}

	log.Println("providers.json has been updated")
	return nil
}
