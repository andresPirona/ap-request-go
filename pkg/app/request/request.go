package request

import (
	"errors"
	"fmt"
	"github.com/andresPirona/ap-request-go/pkg/app/auth"
	"github.com/andresPirona/ap-request-go/pkg/app/domain"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HandleRequest Entry point , central function to handle request
func HandleRequest(opts domain.Options) ([]byte, *http.Response, error) {

	// Http method check
	errMethod := domain.MethodCheck(opts.Method)
	if errMethod != nil {
		return nil, nil, errMethod
	}

	// Check auth
	if opts.Auth != nil {
		errAuth := opts.Auth.Check()
		if errAuth != nil {
			return nil, nil, errAuth
		}
	}

	// Make request
	response, err := makeRequestCentral(opts)
	if err != nil {
		msg := fmt.Sprintf("err_call_request  %s ", opts.Url)
		return nil, nil, errors.New(msg)
	}

	// Read Body
	result, errRB := ioutil.ReadAll(response.Body)
	if errRB != nil {
		msg := fmt.Sprintf("err_read_body  %s ", opts.Url)
		return nil, nil, errors.New(msg)
	}

	return result, response, nil

}

func makeRequestCentral(opts domain.Options) (*http.Response, error) {
	client := createClient()
	var req *http.Request

	var body string
	if opts.Body != nil {
		body = *opts.Body
	}

	// Request initialize
	req, _ = http.NewRequest(opts.Method, opts.Url, strings.NewReader(body))
	if opts.Headers != nil {
		for hName, hValue := range opts.Headers {
			req.Header.Add(hName, hValue)
		}
	} else {
		// default headers
		req.Header.Add("Content-Length", strconv.Itoa(len(body)))
		req.Header.Add("Content-Type", "application/json")
	}

	// Set auth methods
	auth.SetAuthMethod(req, opts.Auth)

	return client.Do(req)
}

// Initializer for the client
func createClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    45 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{Transport: tr}
}
