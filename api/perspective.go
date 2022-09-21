package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
)

func isToxic(markdown string) bool {

	perspectiveKey := os.Getenv("PERSPECTIVE_KEY")

	if perspectiveKey == "" {
		return false
	}

	pl := os.Getenv("PERSPECTIVE_LIMIT")
	perspectiveLimit, err := strconv.ParseFloat(pl, 32)

	if err != nil {
		logger.Errorf("error: Perspective limit not ok: %v", err)
		return true
	}

	perspectiveLang := os.Getenv("PERSPECTIVE_LANGUAGE")

	url := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=" + perspectiveKey
	method := "POST"

	payload := strings.NewReader(`{comment: {text: "` + markdown + `"},
       languages: ["` + perspectiveLang + `"],
       requestedAttributes: {TOXICITY:{}} }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logger.Errorf("error: cannot validate comment using Perspective: %v", err)
		return true
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logger.Errorf("error: cannot validate comment using Perspective: %v", err)
		return true
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Errorf("error: cannot validate comment using Perspective: %v", err)
		return true
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		logger.Errorf("error: cannot parsePerspective result", err)
		return true
	}

	var value float64
	var ok bool
	var resp bool

	value, ok = jsonParsed.Path("attributeScores.TOXICITY.summaryScore.value").Data().(float64)

	if !ok {
		logger.Errorf("error: no valid Perspective value")
		return true
	}

	if value > perspectiveLimit {
		resp = true
	} else {
		resp = false
	}

	return resp
}
