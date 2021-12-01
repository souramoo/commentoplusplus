package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func updateUrlLastModTime(pageUrl string) {
	pathArray := strings.SplitN(pageUrl, "/", 2)
	domain, path := pathArray[0], pathArray[1]

	url := StagingUrl
	if domain == CommentoProdUrl {
		url = ProdUrl
	}

	jsonStr := fmt.Sprintf("{\"url_entries\": {\"url\": %s, \"last_mod\": %s)}}", path, time.Now().Format(time.RFC3339))

	_, _ = http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))

	return
}
