package openbd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DecodeBooks(b []byte) ([]Book, error) {
	var books []Book
	decoder := json.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&books)
	if err != nil {
		return nil, fmt.Errorf("decode books: %w", err)
	}
	return books, nil
}
