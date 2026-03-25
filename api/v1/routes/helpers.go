package analytics_worker

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Body       string
}

func GetResponse(ctx context.Context, url string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err!= nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Do(req)
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := http.ReadBody(resp.Body)
	if err!= nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}, nil
}

func LogError(ctx context.Context, err error) {
	log.Printf("Error: %v, Context: %v", err, ctx)
}

func LogInfo(ctx context.Context, message string) {
	log.Printf("Info: %v, Context: %v", message, ctx)
}

func LogDebug(ctx context.Context, message string) {
	log.Printf("Debug: %v, Context: %v", message, ctx)
}

func IsError(err error) bool {
	if err!= nil {
		return true
	}
	return false
}

func GetStatusCode(err error) int {
	if err!= nil {
		return err.(http.Response).StatusCode
	}
	return 0
}

func GetStatusMessage(err error) string {
	if err!= nil {
		return err.Error()
	}
	return ""
}

func GetError(err error) error {
	if err!= nil {
		return err
	}
	return nil
}

func GetResponseString(ctx context.Context, url string) (string, error) {
	resp, err := GetResponse(ctx, url)
	if err!= nil {
		return "", err
	}

	return resp.Body, nil
}

func GetResponseStatusCode(ctx context.Context, url string) (int, error) {
	resp, err := GetResponse(ctx, url)
	if err!= nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func GetResponseHeaders(ctx context.Context, url string) (http.Header, error) {
	resp, err := GetResponse(ctx, url)
	if err!= nil {
		return nil, err
	}

	return resp.Header, nil
}