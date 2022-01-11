package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func updateUrlLastModTime(pageUrl string) error {
	pathArray := strings.SplitN(pageUrl, "/", 2)
	domain, path := pathArray[0], pathArray[1]

	url := StagingEndPoint
	if domain == CommentoProdDomain {
		url = ProdEndPoint
	}

	body, err := getPayloadBody(path, time.Now().Format(time.RFC3339))

	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	return nil
}

type (
	UrlEntry struct {
		URL     string `json:"url"`
		LastMod string `json:"last_mod"`
	}
	Request struct {
		UrlEntries []UrlEntry `json:"url_entries"`
	}
)

func getPayloadBody(path string, time string) ([]byte, error) {
	var entries []UrlEntry

	entry := UrlEntry{
		URL:     path,
		LastMod: time,
	}

	entries = append(entries, entry)

	data := &Request{
		UrlEntries: entries,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
