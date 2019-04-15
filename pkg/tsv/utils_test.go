package tsv_test

import (
	"github.com/squioc/tsv/pkg/tsv"
	"testing"
)

func TestEscape(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "no escape",
			input:          "Nothing to escape",
			expectedOutput: "Nothing to escape",
		},
		{
			name:           "should escape",
			input:          "\t\\\r\n",
			expectedOutput: "\\t\\\\\\r\\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act
			result := tsv.Escape(test.input)

			// Assert
			if result != test.expectedOutput {
				t.Errorf("Mismatching result. Expected: '%s'. Got: '%s'", test.expectedOutput, result)
			}
		})
	}
}

func TestUnescape(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "no escape",
			input:          "Nothing to escape",
			expectedOutput: "Nothing to escape",
		},
		{
			name:           "should escape",
			input:          "\\t\\\\\\r\\n",
			expectedOutput: "\t\\\r\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act
			result := tsv.Unescape(test.input)

			// Assert
			if result != test.expectedOutput {
				t.Errorf("Mismatching result. Expected: '%s'. Got: '%s'", test.expectedOutput, result)
			}
		})
	}
}
