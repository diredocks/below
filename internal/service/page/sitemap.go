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

func ParseSitemap(data []byte) (*service.Site, []service.Page, error) {
	var sitemap Sitemap
	err := xml.Unmarshal(data, &sitemap)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing XML: %w", err)
	}

	if len(sitemap.URLs) == 0 {
		return nil, nil, fmt.Errorf("sitemap contains no URLs")
	}

	// Extract the host from the first URL entry
	parsedURL, err := url.Parse(sitemap.URLs[0].Loc)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing URL: %w", err)
	}
	site := &service.Site{Host: parsedURL.Host}

	var pages []service.Page
	for _, urlEntry := range sitemap.URLs {
		parsedURL, err := url.Parse(urlEntry.Loc)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing URL: %w", err)
		}
		pages = append(pages, service.Page{
			Path: parsedURL.Path,
		})
	}

	return site, pages, nil
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

func ProcessSitemap(sitemapURL string) (int64, error) {
	// Fetch the sitemap
	data, err := FetchSitemap(sitemapURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch sitemap: %w", err)
	}

	// Parse the sitemap
	site, pages, err := ParseSitemap(data)
	if err != nil {
		return 0, fmt.Errorf("failed to parse sitemap: %w", err)
	}
	site.SiteMap = sitemapURL // Fill Sitemap URL

	// Insert parsed data into the database
	affected, err := InsertPagesDB(site, pages)
	if err != nil {
		return 0, fmt.Errorf("failed to insert sitemap into database: %w", err)
	}

	return affected, nil
}
