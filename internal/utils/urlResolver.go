package utils

import "net/url"

func ResolveURL(src, base string) string {
	u, err := url.Parse(src)
	if err != nil {
		return ""
	}

	if u.IsAbs() {
		return u.String()
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}

	return baseURL.ResolveReference(u).String()
}
