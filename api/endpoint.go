package api

import "fmt"

type Endpoint string

const (
	BooksEndpoint Endpoint = "get"
)

func (e Endpoint) Complement(host string) string {
	return fmt.Sprintf("%s/%s", host, string(e))
}
