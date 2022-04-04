package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func versionPrint() error {
	logger.Infof("starting Commento %s", version)
	return nil
}