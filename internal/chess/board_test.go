// MIT License

// Copyright (c) 2026 Aziz Shamuratov

package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for Str2Sqr method
func TestStr2Sqr(t *testing.T) {

	//Base struct of the method you will test
	type testCase struct {
		//input
		name string
		str  string
		//output
		expected int
		err      bool
	}

	//Actual tests
	tests := []testCase{
		{name: "a1 = 0", str: "a1", expected: 0},
		{name: "e4 = 45", str: "e4", expected: 28},
		{name: "h8 = 63", str: "h8", expected: 63},
	}

	//Loop to run tests one by one
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Str2Sqr(test.str)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
