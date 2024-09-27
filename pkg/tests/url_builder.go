package tests

import (
	"net/url"
	"strconv"
	"strings"
)

type Url struct {
	scheme      string
	host        string
	port        int
	path        string
	queryParams string
}

type IUrlBuilder interface {
	Https() IUrlBuilder
	Host(string) IUrlBuilder
	Port(int) IUrlBuilder
	Path(string) IUrlBuilder
	QueryParams(map[string]string) IUrlBuilder
	GetUrl() *Url
}

type UrlBuilder struct {
	url *Url
}

func NewUrlBuilder() IUrlBuilder {
	return &UrlBuilder{
		url: &Url{
			scheme: "http", // default scheme is http
			port:   0,      // default port is 0 (unspecified)
		},
	}
}

func (b *UrlBuilder) Https() IUrlBuilder {
	b.url.scheme = "https"
	return b
}

func (b *UrlBuilder) Host(host string) IUrlBuilder {
	if host != "" {
		host = strings.TrimSpace(host)
		b.url.host = host
	}
	return b
}

func (b *UrlBuilder) Port(port int) IUrlBuilder {
	if port != 0 {
		b.url.port = port
	}
	return b
}

func (b *UrlBuilder) Path(path string) IUrlBuilder {
	if path != "" {
		path = strings.TrimSpace(path)
		b.url.path = path
	}
	return b
}

func (b *UrlBuilder) QueryParams(params map[string]string) IUrlBuilder {
	if len(params) > 0 {
		// Create a URL query string from the map
		values := url.Values{}
		for key, value := range params {
			if key == "" {
				continue
			}
			values.Set(key, value)
		}
		b.url.queryParams = values.Encode()
	}

	return b
}

func (b *UrlBuilder) GetUrl() *Url {
	return b.url
}

func (u *Url) Build() string {
	// Construct the base URL with scheme and host
	var builder strings.Builder
	builder.WriteString(u.scheme + "://")
	builder.WriteString(u.host)

	// Append port if it's set (not zero)
	if u.port != 0 {
		builder.WriteString(":" + strconv.Itoa(u.port))
	}

	// Append the path
	if u.path != "" {
		builder.WriteString("/" + strings.TrimPrefix(u.path, "/"))
	}

	// Append query parameters if they exist
	if u.queryParams != "" {
		builder.WriteString("?" + u.queryParams)
	}

	return builder.String()
}
