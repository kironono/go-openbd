package openbd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		value []byte
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			name: "Hyphen-separated date date only",
			args: args{
				value: []byte(`"2001-01-02"`),
			},
			want: Date{time.Date(2001, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "Hyphen-separated date date time",
			args: args{
				value: []byte(`"2001-01-02 12:13:14"`),
			},
			want: Date{time.Date(2001, 1, 2, 12, 13, 14, 0, time.UTC)},
		},
		{
			name: "Hyphen-separated date zero date time",
			args: args{
				value: []byte(`"0000-00-00 00:00:00"`),
			},
			want: Date{time.Time{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{time.Time{}}
			err := d.UnmarshalJSON(tt.args.value)
			require.NoError(t, err)
			assert.Equal(t, tt.want, d)
		})
	}
}
