package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func versionPrint() error {
	logger.Infof("starting Commento %s", version)
	return nil
}

func versionCheckStart() error {
	go func() {
		printedError := false
		errorCount := 0
		latestSeen := ""

		for {
			resp, err := http.Get("https://api.github.com/repos/souramoo/commentoplusplus/releases/latest")
			if err != nil {
				errorCount++
				// print the error only once; we don't want to spam the logs with this
				// every five minutes
				if !printedError && errorCount > 5 {
					logger.Errorf("error checking version: %v", err)
					printedError = true
				}
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errorCount++
				if !printedError && errorCount > 5 {
					logger.Errorf("error reading body: %s", err)
					printedError = true
				}
				continue
			}

			type response struct {
				Latest string `json:"tag_name"`
			}

			r := response{}
			json.Unmarshal(body, &r)

			if !strings.HasPrefix(version, r.Latest) && r.Latest != latestSeen {
				logger.Infof("New update available! Latest version: %s", r.Latest)
				latestSeen = r.Latest
			}

			errorCount = 0
			printedError = false

			time.Sleep(5 * time.Minute)
		}
	}()

	return nil
}
