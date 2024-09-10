package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	// Create a new http client, where there is a timeout limit of 10sec
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Try to get a response from url
	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	// Close the connection at the end of getting everything we need
	defer resp.Body.Close()

	// Read all data that came in
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	// Create a struct instance
	rssFeed := RSSFeed{}

	// Get everything into the struct instance and return once finished
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
