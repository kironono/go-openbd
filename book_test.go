package openbd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBook_DecodeBooks(t *testing.T) {
	type args struct {
		value []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []Book
		wantErr bool
	}{
		{
			name: "Decode books",
			args: args{
				value: []byte(`[{
					"summary": {
						"isbn": "1234567890"
					}
				}]`),
			},
			want: []Book{
				{
					Summary: summary{
						Isbn: "1234567890",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Decode failed",
			args: args{
				value: []byte(`{}`),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := DecodeBooks(tt.args.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
