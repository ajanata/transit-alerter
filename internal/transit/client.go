package transit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

const base = "https://external.transitapp.com/v3"

type Client struct {
	apiKey string
	http   *http.Client
}

func New(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &Client{
		apiKey: apiKey,
		http:   httpClient,
	}
}

func (c *Client) get(ctx context.Context, uri string, params any, body any) error {
	reader, err := c.getRaw(ctx, uri, params)

	decode := json.NewDecoder(reader)
	err = decode.Decode(body)
	_ = reader.Close()
	if err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}

func (c *Client) getRaw(ctx context.Context, uri string, params any) (io.ReadCloser, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, fmt.Errorf("encoding params: %w", err)
	}

	url := fmt.Sprintf("%s%s?%s", base, uri, values.Encode())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("apiKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	return resp.Body, nil
}
