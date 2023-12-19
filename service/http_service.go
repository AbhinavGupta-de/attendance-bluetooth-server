package service

import (
	"bytes"
	"io"
	"net/http"
)

func PostData(url string, data []byte) (string, error) {
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(responseBody), nil
}
