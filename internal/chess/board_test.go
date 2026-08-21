// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	type testCase struct {
		//input
		name string
		a    int
		b    int
		//output
		expected int
		err      bool
	}

	tests := []testCase{
		{name: "adds 2 pos", a: 1, b: 2, expected: 3},
		{name: "adds 1 pos 1 neg", a: 6, b: -2, expected: 4},
		{name: "adds 2 neg", a: -1, b: -2, expected: -3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Sum(test.a, test.b)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStr2Sqr(t *testing.T) {
	type testCase struct {
		//input
		name string
		str  string
		//output
		expected int
		err      bool
	}

	tests := []testCase{
		{name: "e4 = 45", str: "e4", expected: 28},
		{name: "a1 = 0", str: "a1", expected: 0},
		{name: "h8 = 63", str: "h8", expected: 63},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Str2Sqr(test.str)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
