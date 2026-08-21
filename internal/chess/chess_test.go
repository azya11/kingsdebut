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
