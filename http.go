package gotoolbox

import (
	"errors"
	"net/http"
)

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}
