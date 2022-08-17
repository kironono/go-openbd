package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint_Complement(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name     string
		endpoint Endpoint
		args     args
		want     string
	}{
		{
			name:     "Complement simple",
			endpoint: "get",
			args: args{
				host: "https://example.com",
			},
			want: "https://example.com/get",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.endpoint.Complement(tt.args.host)
			assert.Equal(t, tt.want, actual)
		})
	}
}
