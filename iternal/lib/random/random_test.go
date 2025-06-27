package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	testArr := []struct {
		length string
		size   int
	}{
		{
			length: "size = 2",
			size:   2,
		},
		{
			length: "size = 10",
			size:   10,
		},
		{
			length: "size = 30",
			size:   30,
		},
	}

	for _, tc := range testArr {
		t.Run(tc.length, func(t *testing.T) {
			str1 := NewRandomString(tc.size)
			str2 := NewRandomString(tc.size)

			assert.Len(t, str1, tc.size)
			assert.Len(t, str2, tc.size)

			assert.NotEqual(t, str1, str2)
		})
	}
}
