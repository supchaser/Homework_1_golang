package uniqFunc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueFunction(t *testing.T) {
	tests := map[string]struct {
		input []string
		options UniqOptions
		result []string
	} {
		"Без параметров":  {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{},
			result: []string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		
		"С параметром input_file": {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{},
			result: []string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		
		"С параметрами input_file и output_file": {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{},
			result: []string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		
		"С параметром -c": {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{Count: true},
			result: []string {
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik.",
			},
		},
		
		"С параметром -d": {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{Repeated: true},
			result: []string {
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		
		"С параметром -u": {
			input: []string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options: UniqOptions{Unique: true},
			result: []string {
				"",
				"Thanks.",
			},
		},
		
		"С параметром -i": {
			input: []string {
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
			options: UniqOptions{IgnoreCase: true},
			result: []string {
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
				"I love music of kartik.",
			},
		},
		
		"С параметром -f num(1)": {
			input: []string {
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			options: UniqOptions{SkipFields: 1},
			result: []string {
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		
		"С параметром -s num(1)": {
			input: []string {
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			options: UniqOptions{SkipChars: 1},
			result: []string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},

		"С параметром -f num(1) и -s num(1)": {
			input: []string {
				"We love music.",
				"I love music.",
				"They lave music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",

			},
			options: UniqOptions{SkipFields: 1, SkipChars: 1},
			result: []string {
				"We love music.",
				"They lave music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult := UniqFunc(test.input, test.options)
			require.Equal(t, test.result, actualResult, "Тест провален для случая: %s", name)
		})
	}
}
