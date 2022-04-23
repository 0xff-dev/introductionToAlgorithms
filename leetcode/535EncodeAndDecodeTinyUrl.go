package leetcode

import (
	"fmt"
	"net/url"
)

type Codec struct {
	forShortMap map[string]string
	domain      map[string]int
}

func FConstructor() Codec {
	return Codec{
		forShortMap: make(map[string]string),
		domain:      make(map[string]int),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	ul, err := url.Parse(longUrl)
	if err != nil {
		return ""
	}

	hostname := ul.Hostname()
	schema := ul.Scheme

	if idx, ok := this.forShortMap[longUrl]; ok {
		return fmt.Sprintf("%s://%s%s", schema, hostname, idx)
	}
	this.domain[hostname]++
	this.forShortMap[longUrl] = fmt.Sprintf("/%d", this.domain[hostname])
	return fmt.Sprintf("%s://%s/%d", schema, hostname, this.domain[hostname])
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	ul, err := url.Parse(shortUrl)
	if err != nil {
		return ""
	}

	hostname := ul.Hostname()
	path := ul.EscapedPath()
	if _, ok := this.domain[hostname]; !ok {
		return ""
	}

	for u, p := range this.forShortMap {
		if p == path {
			return u
		}
	}
	return ""
}
