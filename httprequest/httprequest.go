package httprequest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ErrStatusCodeNotOk = errors.New("status code not ok")
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

func (h *HTTPRequest) DoRequest(method, url string, out interface{}) error {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	response, err := h.Client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return ErrStatusCodeNotOk
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	return nil
}
