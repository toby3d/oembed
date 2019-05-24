package oembed

import (
	"regexp"
	"strings"

	http "github.com/valyala/fasthttp"
)

type providerCandidate struct {
	Domain       string
	ProviderName string
	ProviderURL  string
	Schemes      []string
	URL          string
}

func getHostname(url string) string {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.Update(url)
	if u.Host() != nil {
		return strings.TrimPrefix(string(u.Host()), "www.")
	}
	return ""
}

func makeCandidate(p Provider) providerCandidate {
	endpoint := p.Endpoints[0]
	domain := getHostname(endpoint.URL)
	if domain != "" {
		domain = strings.TrimPrefix(domain, "www.")
	} else {
		domain = ""
	}

	return providerCandidate{
		ProviderName: p.Name,
		ProviderURL:  p.URL,
		Schemes:      endpoint.Schemes,
		URL:          endpoint.URL,
		Domain:       domain,
	}

}

func findProvider(url string) *Provider {
	var candidates []Provider
	for _, provider := range Providers {
		provider := provider

		endpoint := provider.Endpoints[0]
		domain := getHostname(endpoint.URL)
		if domain != "" {
			domain = strings.TrimPrefix(domain, "www.")
		} else {
			domain = ""
		}

		if len(endpoint.Schemes) == 0 {
			if !strings.Contains(url, domain) {
				continue
			}
			candidates = append(candidates, provider)
			continue
		}
		for _, scheme := range endpoint.Schemes {
			scheme := scheme
			reg := regexp.MustCompile(strings.Replace(scheme, "*", "(.*)", -1))
			if !reg.MatchString(url) {
				continue
			}

			candidates = append(candidates, provider)
			break
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	return &candidates[0]
}
