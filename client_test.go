package openbd

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_NewClient(t *testing.T) {
	type args struct {
		client *http.Client
		host   string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "New client",
			args: args{
				client: http.DefaultClient,
				host:   "https://example.com",
			},
			want: &Client{
				Client: http.DefaultClient,
				Host:   "https://example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewClient(tt.args.client, tt.args.host)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestClient_DefaultClient(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "New default client",
			want: &Client{
				Client: http.DefaultClient,
				Host:   "https://api.openbd.jp/v1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DefaultClient()
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestClient_Books(t *testing.T) {
	type fields struct {
		Client *http.Client
		Host   string
	}
	type args struct {
		ids []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Book
	}{
		{
			name: "Get a books with single ISBN",
			fields: fields{
				Client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `[{
							"summary": {
								"author": "AutherName",
								"cover": "https:\/\/example.com\/1234567890.jpg",
								"isbn": "1234567890",
								"pubdate": "20010101",
								"publisher": "PublisherName",
								"series": "SeriesName",
								"title": "Title",
								"Volume": "vol01"
							}
						}]`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
						Header:     http.Header{},
					}

				}),
				Host: "https://example.com/",
			},
			args: args{
				ids: []string{"1234567890"},
			},
			want: []Book{
				{
					Summary: summary{
						Author:    "AutherName",
						Cover:     "https://example.com/1234567890.jpg",
						Isbn:      "1234567890",
						Pubdate:   "20010101",
						Publisher: "PublisherName",
						Series:    "SeriesName",
						Title:     "Title",
						Volume:    "vol01",
					},
				},
			},
		},
		{
			name: "Get books with multiple ISBNs",
			fields: fields{
				Client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `[{
							"summary": {
								"author": "AutherName",
								"cover": "https:\/\/example.com\/1111111111111.jpg",
								"isbn": "1111111111111",
								"pubdate": "20010101",
								"publisher": "PublisherName",
								"series": "SeriesName",
								"title": "Title",
								"Volume": "vol01"
							}
						},{
							"summary": {
								"author": "AutherName",
								"cover": "https:\/\/example.com\/2222222222222.jpg",
								"isbn": "2222222222222",
								"pubdate": "20010101",
								"publisher": "PublisherName",
								"series": "SeriesName",
								"title": "Title",
								"Volume": "vol01"
							}
						}]`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
						Header:     http.Header{},
					}

				}),
				Host: "https://example.com/",
			},
			args: args{
				ids: []string{"1111111111111", "2222222222222"},
			},
			want: []Book{
				{
					Summary: summary{
						Author:    "AutherName",
						Cover:     "https://example.com/1111111111111.jpg",
						Isbn:      "1111111111111",
						Pubdate:   "20010101",
						Publisher: "PublisherName",
						Series:    "SeriesName",
						Title:     "Title",
						Volume:    "vol01",
					},
				},
				{
					Summary: summary{
						Author:    "AutherName",
						Cover:     "https://example.com/2222222222222.jpg",
						Isbn:      "2222222222222",
						Pubdate:   "20010101",
						Publisher: "PublisherName",
						Series:    "SeriesName",
						Title:     "Title",
						Volume:    "vol01",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			c := NewClient(tt.fields.Client, tt.fields.Host)
			actual, err := c.Books(ctx, tt.args.ids)
			require.NoError(t, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}

type roundTripFunc func(request *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func mockHTTPClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}
