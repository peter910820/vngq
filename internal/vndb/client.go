package vndb

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const endpoint string = "https://api.vndb.org/kana/"

func sendPostRequest(apiRoute string, jsonBytes []byte) ([]byte, error) {

	resp, err := http.Post(endpoint+apiRoute, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error for response: %d", resp.StatusCode)
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return r, nil
}
