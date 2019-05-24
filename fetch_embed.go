package oembed

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
	template "github.com/valyala/fasttemplate"
)

// Params represent a optional parameters for Extract method.
type Params struct {
	MaxWidth  int
	MaxHeight int
}

func fetchEmbed(url string, provider *Provider, params *Params) (*OEmbed, error) {
	resourceURL := provider.Endpoints[0].URL
	resourceURL = template.ExecuteString(resourceURL, "{", "}", map[string]interface{}{"format": "json"})

	link := http.AcquireURI()
	defer http.ReleaseURI(link)
	link.Update(resourceURL)
	qa := link.QueryArgs()
	qa.Add("format", "json")
	qa.Add("url", url)

	if params != nil && params.MaxWidth != 0 {
		qa.Add("maxwidth", strconv.Itoa(params.MaxWidth))
	}
	if params != nil && params.MaxHeight != 0 {
		qa.Add("maxheight", strconv.Itoa(params.MaxHeight))
	}
	link.SetQueryStringBytes(qa.QueryString())

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetRequestURIBytes(link.FullURI())

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := http.Do(req, resp); err != nil {
		return nil, Error{
			Message: err.Error(),
			URL:     url,
		}
	}

	var oEmbed OEmbed
	if err := json.UnmarshalFast(resp.Body(), &oEmbed); err != nil {
		return nil, Error{
			Message: err.Error(),
			URL:     url,
		}
	}
	oEmbed.ProviderName = provider.Name
	oEmbed.ProviderURL = provider.URL
	return &oEmbed, nil
}
