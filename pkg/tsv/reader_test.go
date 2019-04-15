package tsv_test

import (
	"github.com/squioc/tsv/pkg/tsv"
	"io"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  [][]string
	}{
		{
			name:  "no row",
			input: strings.NewReader(""),
			want:  [][]string{},
		},
		{
			name:  "empty row",
			input: strings.NewReader("\n"),
			want: [][]string{
				{""},
			},
		},
		{
			name:  "one row",
			input: strings.NewReader("cell1\tcell2\tcell3\n"),
			want: [][]string{
				{"cell1", "cell2", "cell3"},
			},
		},
		{
			name:  "multiple rows",
			input: strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\ncell31\tcell32\tcell33\n"),
			want: [][]string{
				{"cell11", "cell12", "cell13"},
				{"cell21", "cell22", "cell23"},
				{"cell31", "cell32", "cell33"},
			},
		},
		{
			name:  "row with escaped characters",
			input: strings.NewReader("NL\\n\ttab\\t\tbackslash\\\\\tCR\\r\n"),
			want: [][]string{
				{"NL\n", "tab\t", "backslash\\", "CR\r"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			reader := tsv.NewReader(test.input)

			for rowNumber, expectedRow := range test.want {
				// Act
				actualRow, err := reader.Read()

				// Assert
				if err != nil && err != io.EOF {
					t.Errorf("Unexpected error: %s", err)
				}

				expectedLen := len(expectedRow)
				actualLen := len(actualRow)
				if actualLen != expectedLen {
					t.Errorf("Mismatching length. Expected: %d. Got: %d", expectedLen, actualLen)
				} else {
					for columnIndex, expectedCell := range expectedRow {
						actualCell := actualRow[columnIndex]
						if actualCell != expectedCell {
							t.Errorf("Mismatching value at (%d, %d). Expected: %s. Got: %s", rowNumber, columnIndex, expectedCell, actualCell)
						}
					}
				}
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	tests := []struct {
		name            string
		input           io.Reader
		expectedRecords [][]string
	}{
		{
			name:            "no row",
			input:           strings.NewReader(""),
			expectedRecords: [][]string{},
		},
		{
			name:  "empty row",
			input: strings.NewReader("\n"),
			expectedRecords: [][]string{
				{""},
			},
		},
		{
			name:  "one row",
			input: strings.NewReader("cell1\tcell2\tcell3\n"),
			expectedRecords: [][]string{
				{"cell1", "cell2", "cell3"},
			},
		},
		{
			name:  "multiple rows",
			input: strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\ncell31\tcell32\tcell33\n"),
			expectedRecords: [][]string{
				{"cell11", "cell12", "cell13"},
				{"cell21", "cell22", "cell23"},
				{"cell31", "cell32", "cell33"},
			},
		},
		{
			name:  "row with escaped characters",
			input: strings.NewReader("NL\\n\ttab\\t\tbackslash\\\\\tCR\\r\n"),
			expectedRecords: [][]string{
				{"NL\n", "tab\t", "backslash\\", "CR\r"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			reader := tsv.NewReader(test.input)

			// Act
			actualRecords, err := reader.ReadAll()

			// Assert
			if err != nil && err != io.EOF {
				t.Errorf("Unexpected error: %s", err)
			}

			expectedLen := len(test.expectedRecords)
			actualLen := len(actualRecords)
			if actualLen != expectedLen {
				t.Errorf("Mismatching records length. Expected: %d. Got: %d", expectedLen, actualLen)
			} else {
				for rowIndex, expectedRow := range test.expectedRecords {
					actualRow := actualRecords[rowIndex]
					expectedRowLen := len(expectedRow)
					actualRowLen := len(actualRow)

					if actualRowLen != expectedRowLen {
						t.Errorf("Mismatching row length. Expected: %d. Got: %d", expectedRowLen, actualRowLen)
					} else {
						for columnIndex, expectedCell := range expectedRow {
							actualCell := actualRecords[rowIndex][columnIndex]
							if actualCell != expectedCell {
								t.Errorf("Mismatching value at (%d, %d). Expected: %s. Got: %s", rowIndex, columnIndex, expectedCell, actualCell)
							}
						}
					}
				}
			}
		})
	}
}
