package shuntingYardAlg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testFailInput = `2 / 0 + 1 2`

func TestFail(t *testing.T) {
	tokens := Tokenize(testFailInput)
	_, err := ShuntingYardAlg(tokens)
	if err == nil {
		t.Errorf("Test FAIL failed: expected error")
	}
}

func TestVariousCases(t *testing.T) {
	tests := map[string]struct {
		input  string
		result float64
		err    error
	}{
		"Test 1": {
			input:  "1 + 2",
			result: 3,
			err:    nil,
		},
		"Test 2": {
			input:  "2 * 3",
			result: 6,
			err:    nil,
		},
		"Test 3": {
			input:  "1 + 2 * 3",
			result: 7,
			err:    nil,
		},
		"Test 4": {
			input:  "(1 + 2) * 3",
			result: 9,
			err:    nil,
		},
		"Test 5": {
			input:  "2 * -3",
			result: -6,
			err:    nil,
		},
		"Test 6": {
			input:  "1.5 * 2",
			result: 3,
			err:    nil,
		},
		"Test 7": {
			input:  "(-1)(-1)",
			result: 1,
			err:    nil,
		},
		"Test 8": {
			input:  "(-1 + 2) * 3 - 4 / 2",
			result: 1,
			err:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(test.input)
			actualResult, err := ShuntingYardAlg(tokens)
			if test.err != nil {
				require.Error(t, err)
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, test.result, actualResult)
			}
		})
	}
}
