//go:generate ffjson $GOFILE
package oembed

type (
	// Provider represent a single provider info
	Provider struct {
		Name      string     `json:"provider_name"`
		URL       string     `json:"provider_url"`
		Endpoints []Endpoint `json:"endpoints"`
	}

	// Endpoint represent a single endpoint of Provider
	Endpoint struct {
		Schemes   []string `json:"schemes,omitempty"`
		URL       string   `json:"url"`
		Discovery bool     `json:"discovery,omitempty"`
		Formats   []string `json:"formats,omitempty"`
	}

	// OEmbed can specify a resource type, such as photo or video.
	// Each type has specific parameters associated with it.
	OEmbed struct {
		// The resource type.
		Type string `json:"type"` // required

		// The oEmbed version number.
		Version string `json:"version"` // required

		// A text title, describing the resource.
		Title string `json:"title,omitempty"`

		// The name of the author/owner of the resource.
		AuthorName string `json:"author_name,omitempty"`

		// A URL for the author/owner of the resource.
		AuthorURL string `json:"author_url,omitempty"`

		// The name of the resource provider.
		ProviderName string `json:"provider_name,omitempty"`

		// The url of the resource provider.
		ProviderURL string `json:"provider_url,omitempty"`

		// The suggested cache lifetime for this resource, in seconds.
		// Consumers may choose to use this value or not.
		CacheAge int `json:"cache_age,omitempty"`

		// A URL to a thumbnail image representing the resource.
		// The thumbnail must respect any maxwidth and maxheight parameters.
		// If this parameter is present, thumbnail_width and thumbnail_height must also be present.
		ThumbnailURL string `json:"thumbnail_url,omitempty"`

		// The width of the optional thumbnail.
		// If this parameter is present, thumbnail_url and thumbnail_height must also be present.
		ThumbnailWidth int `json:"thumbnail_width,omitempty"`

		// The height of the optional thumbnail.
		// If this parameter is present, thumbnail_url and thumbnail_width must also be present.
		ThumbnailHeight int `json:"thumbnail_height,omitempty"`

		URL string `json:"url,omitempty"`
	}

	// Photo is used for representing static photos.
	Photo struct {
		// The source URL of the image. Consumers should be able to insert this URL into an <img> element.
		// Only HTTP and HTTPS URLs are valid.
		URL string `json:"url"` // required

		// The width in pixels of the image specified in the url parameter.
		Width int `json:"width"` // required

		// The height in pixels of the image specified in the url parameter.
		Height int `json:"height"` // required
	}

	// Video is used for representing playable videos.
	Video struct {
		// The HTML required to embed a video player. The HTML should have no padding or margins.
		// Consumers may wish to load the HTML in an off-domain iframe to avoid XSS vulnerabilities.
		HTML string `json:"html"` // required

		// The width in pixels required to display the HTML.
		Width int `json:"width"` // required

		// The height in pixels required to display the HTML.
		Height int `json:"height"` // required
	}

	// Link type allow a provider to return any generic embed data (such as title and author_name), without
	// providing either the url or html parameters. The consumer may then link to the resource, using the URL
	// specified in the original request.
	// Link string

	// Rich is used for rich HTML content that does not fall under one of the other categories.
	Rich struct {
		// The HTML required to display the resource. The HTML should have no padding or margins.
		// Consumers may wish to load the HTML in an off-domain iframe to avoid XSS vulnerabilities.
		// The markup should be valid XHTML 1.0 Basic.
		HTML string `json:"html"` // required

		// The width in pixels required to display the HTML.
		Width int `json:"width"` // required

		// The height in pixels required to display the HTML.
		Height int `json:"height"` // required
	}
)
