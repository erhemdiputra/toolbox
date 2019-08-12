package httprequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrStatusCodeNotOk = errors.New("status code not ok")
)

const (
	ContentTypeApplicationJSON = "application/json"
	ContentTypeFormURLEncoded  = "application/x-www-form-urlencoded"
)

type HTTPRequest struct {
	Client *http.Client
}

func NewHTTPRequest(timeout time.Duration) *HTTPRequest {
	return &HTTPRequest{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func parseToURLValues(param map[string]string) url.Values {
	queryString := make(url.Values)

	for key, value := range param {
		queryString.Add(key, value)
	}

	return queryString
}

func (h *HTTPRequest) Get(URL string, header http.Header, param map[string]string, out interface{}) error {
	urlValues := parseToURLValues(param)

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return errors.New("failed to create get request")
	}

	request.Header = header
	request.URL.RawQuery = urlValues.Encode()

	return h.DoRequest(request, out)
}

func (h *HTTPRequest) Post(URL string, header http.Header, param map[string]string, payload interface{}, out interface{}) error {
	var paramBuffer *bytes.Buffer
	var err error

	if payload != nil {
		header.Set("Content-Type", ContentTypeApplicationJSON)
		marshalled, err := json.Marshal(payload)
		if err != nil {
			return errors.New("failed to marshal payload")
		}
		paramBuffer = bytes.NewBuffer(marshalled)

	} else {
		header.Set("Content-Type", ContentTypeFormURLEncoded)
		queryStringEncoded := parseToURLValues(param).Encode()
		paramBuffer = bytes.NewBufferString(queryStringEncoded)
	}

	request, err := http.NewRequest("POST", URL, paramBuffer)
	if err != nil {
		return errors.New("failed to create post request")
	}

	request.Header = header

	return h.DoRequest(request, out)
}

func (h *HTTPRequest) DoRequest(request *http.Request, out interface{}) error {
	response, err := h.Client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated &&
		response.StatusCode != http.StatusNoContent {
		return ErrStatusCodeNotOk
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	return nil
}
