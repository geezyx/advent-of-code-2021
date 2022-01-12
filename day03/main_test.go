package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonUncommon(t *testing.T) {
	for name, tt := range map[string]struct {
		in               bit
		expectedCommon   string
		expectedUncommon string
	}{
		"equal": {
			in: bit{
				ones:   5,
				zeroes: 5,
			},
			expectedCommon:   "1",
			expectedUncommon: "0",
		},
		"one": {
			in: bit{
				ones:   5,
				zeroes: 4,
			},
			expectedCommon:   "1",
			expectedUncommon: "0",
		},
		"zero": {
			in: bit{
				ones:   4,
				zeroes: 5,
			},
			expectedCommon:   "0",
			expectedUncommon: "1",
		},
	} {
		t.Run(name, func(t *testing.T) {
			c := tt.in.common()
			assert.Equal(t, tt.expectedCommon, c)
			u := tt.in.uncommon()
			assert.Equal(t, tt.expectedUncommon, u)
		})
	}
}

func TestFilterInputs(t *testing.T) {
	for name, tt := range map[string]struct {
		in       []string
		index    int
		value    string
		expected []string
	}{
		"exact": {
			in:       []string{"001"},
			index:    0,
			value:    "0",
			expected: []string{"001"},
		},
		"exact second index": {
			in:       []string{"001"},
			index:    1,
			value:    "0",
			expected: []string{"001"},
		},
		"exact third index": {
			in:       []string{"001"},
			index:    2,
			value:    "1",
			expected: []string{"001"},
		},
		"exclude": {
			in:       []string{"001"},
			index:    0,
			value:    "1",
			expected: []string{},
		},
		"filter zeroes": {
			in:       []string{"001", "010", "100", "110"},
			index:    0,
			value:    "0",
			expected: []string{"001", "010"},
		},
		"filter ones": {
			in:       []string{"001", "010", "100", "110"},
			index:    0,
			value:    "1",
			expected: []string{"100", "110"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			res := filterInputs(tt.in, tt.index, tt.value)
			assert.Equal(t, tt.expected, res)
		})
	}
}
