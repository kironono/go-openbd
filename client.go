package openbd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kironono/go-openbd/api"
)

type Client struct {
	Client *http.Client
	Host   string
}

func NewClient(client *http.Client, host string) *Client {
	return &Client{
		client,
		host,
	}
}

func DefaultClient() *Client {
	return &Client{
		Client: http.DefaultClient,
		Host:   "https://api.openbd.jp/v1",
	}
}

func (c *Client) Books(ids []string) ([]Book, error) {
	endpoint := api.BooksEndpoint.Complement(c.Host)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	q := req.URL.Query()
	q.Add("isbn", strings.Join(ids, ","))
	req.URL.RawQuery = q.Encode()

	body, err := get(c, req)
	if err != nil {
		return nil, err
	}

	// decode books
	books, err := DecodeBooks(body)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func get(c *Client, req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	return body, nil
}
