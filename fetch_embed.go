package oembed

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
	template "github.com/valyala/fasttemplate"
)

type Params struct {
	MaxWidth  int
	MaxHeight int
}

func fetchEmbed(url string, provider providerCandidate, params *Params) (*Response, error) {
	resourceUrl := provider.URL
	resourceUrl = template.ExecuteString(resourceUrl, "{", "}", map[string]interface{}{"format": "json"})

	link := http.AcquireURI()
	defer http.ReleaseURI(link)
	link.Update(resourceUrl)
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
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}
	response.ProviderName = provider.ProviderName
	response.ProviderURL = provider.ProviderURL
	return &response, nil
}
