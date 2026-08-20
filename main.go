package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HTTPLoader handles fetching resources over HTTP.
type HTTPLoader struct {
	Client *http.Client
}

// Load fetches a resource using the provided context to ensure cancellation propagation.
func (h *HTTPLoader) Load(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return h.Client.Do(req)
}

func main() {
	loader := &HTTPLoader{Client: &http.Client{}}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Example usage
	_, err := loader.Load(ctx, "https://httpbin.org/delay/2")
	if err != nil {
		fmt.Printf("Request failed as expected: %v\n", err)
	}
}