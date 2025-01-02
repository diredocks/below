package page

import (
	"below/internal/service"

	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []struct {
		Loc string `xml:"loc"`
	} `xml:"url"`
}

func ParseSitemap(data []byte) ([]service.Page, error) {
	var sitemap Sitemap
	err := xml.Unmarshal(data, &sitemap)
	if err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}

	var pages []service.Page
	for _, urlEntry := range sitemap.URLs {
		parsedURL, err := url.Parse(urlEntry.Loc)
		if err != nil {
			return nil, fmt.Errorf("error parsing URL: %w", err)
		}
		pages = append(pages, service.Page{
			Site: parsedURL.Host,
			Path: parsedURL.Path,
		})
	}

	return pages, nil
}

func FetchSitemap(sitemapURL string) ([]byte, error) {
	resp, err := http.Get(sitemapURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching sitemap: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return data, nil
}
