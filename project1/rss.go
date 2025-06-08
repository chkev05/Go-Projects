package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RSSFeed represents the structure of an RSS feed.
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// RSSItem represents an individual item in an RSS feed.
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// urlToFeed fetches an RSS feed from the given URL and returns it as an RSSFeed struct.
func urlToFeed(url string) (RSSFeed, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to fetch RSS feed: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to read response body: %v", err)
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to unmarshal XML: %v", err)
	}
	return rssFeed, nil

}
