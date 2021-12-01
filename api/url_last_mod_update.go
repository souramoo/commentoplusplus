package main

import (
	"bytes"
	"fmt"
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

	jsonStr := fmt.Sprintf("{\"url_entries\": {\"url\": %s, \"last_mod\": %s)}}", path, time.Now().Format(time.RFC3339))

	_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))

	if err != nil {
		return err
	}
	return nil
}
