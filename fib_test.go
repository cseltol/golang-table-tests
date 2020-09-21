package fib_test

import (
	"testing"

	"fib"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "1_test",
			num:  1,
			want: 1,
		},
		{
			name: "2_test",
			num:  2,
			want: 1,
		},
		{
			name: "3_test",
			num:  3,
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib.Fib(tc.num))
		})
	}
}
