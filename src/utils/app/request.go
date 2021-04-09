package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validator.FieldError) {
	// for _, err := range errors {
	// logger.Logger.Info(err.Key, err.Value())
	// }

	return
}

// Request do request/
func Request(method string, url string, headers map[string]string, bodyRequest interface{}) (res *http.Response, err error) {
	reqBody, err := json.Marshal(bodyRequest)

	if err != nil {
		log.Fatal(err)
	}

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 30 * time.Second,
	}

	var client = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		panic(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}
