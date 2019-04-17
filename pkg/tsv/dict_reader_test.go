package tsv_test

import (
	"github.com/squioc/tsv/pkg/tsv"
	"io"
	"strings"
	"testing"
)

func TestDictRead(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		input   io.Reader
		want    []map[string]string
		wantErr error
	}{
		{
			name:   "read records",
			fields: []string{"key1", "key2", "key3"},
			input:  strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\ncell31\tcell32\tcell33\n"),
			want: []map[string]string{
				map[string]string{
					"key1": "cell11",
					"key2": "cell12",
					"key3": "cell13",
				},
				map[string]string{
					"key1": "cell21",
					"key2": "cell22",
					"key3": "cell23",
				},
				map[string]string{
					"key1": "cell31",
					"key2": "cell32",
					"key3": "cell33",
				},
			},
			wantErr: nil,
		},
		{
			name:    "row size overflow",
			fields:  []string{"key1", "key2", "key3"},
			input:   strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\tcell24\ncell31\tcell32\tcell33\n"),
			want:    nil,
			wantErr: tsv.ErrRowSizeOverflow,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			reader := tsv.NewDictReader(test.input, test.fields)

			for recordIndex, expectedRecord := range test.want {
				// Act
				actualRecord, err := reader.Read()

				// Assert
				if err != nil && err != io.EOF && err != test.wantErr {
					t.Errorf("Unexpected error. Expected: %s. Got: %s", test.wantErr, err)
				}

				expectedLen := len(expectedRecord)
				actualLen := len(actualRecord)
				if actualLen != expectedLen {
					t.Errorf("Mismatching length. Expected: %d. Got: %d", expectedLen, actualLen)
				} else {
					for key, expectedValue := range expectedRecord {
						actualValue, ok := actualRecord[key]
						if !ok {
							t.Errorf("Missing field. name: '%s', at %d", key, recordIndex)
						}
						if actualValue != expectedValue {
							t.Errorf("Mismatching value for field '%s' et %d. Expected: %s. Got: %s", key, recordIndex, expectedValue, actualValue)
						}
					}
				}
			}
		})
	}
}

func TestDictReadAll(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		input   io.Reader
		want    []map[string]string
		wantErr error
	}{
		{
			name:   "read records",
			fields: []string{"key1", "key2", "key3"},
			input:  strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\ncell31\tcell32\tcell33\n"),
			want: []map[string]string{
				map[string]string{
					"key1": "cell11",
					"key2": "cell12",
					"key3": "cell13",
				},
				map[string]string{
					"key1": "cell21",
					"key2": "cell22",
					"key3": "cell23",
				},
				map[string]string{
					"key1": "cell31",
					"key2": "cell32",
					"key3": "cell33",
				},
			},
			wantErr: nil,
		},
		{
			name:    "row size overflow",
			fields:  []string{"key1", "key2", "key3"},
			input:   strings.NewReader("cell11\tcell12\tcell13\ncell21\tcell22\tcell23\tcell24\ncell31\tcell32\tcell33\n"),
			want:    nil,
			wantErr: tsv.ErrRowSizeOverflow,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			reader := tsv.NewDictReader(test.input, test.fields)

			// Act
			actualRecords, err := reader.ReadAll()

			// Assert
			if err != nil && err != io.EOF && err != test.wantErr {
				t.Errorf("Unexpected error. Expected: %s. Got: %s", test.wantErr, err)
			}

			expectedLen := len(test.want)
			actualLen := len(actualRecords)
			if actualLen != expectedLen {
				t.Errorf("Mismatching records length. Expected: %d. Got: %d", expectedLen, actualLen)
			} else {
				for rowIndex, expectedRecord := range test.want {
					actualRecord := actualRecords[rowIndex]
					expectedRecordLen := len(expectedRecord)
					actualRecordLen := len(actualRecord)

					if actualRecordLen != expectedRecordLen {
						t.Errorf("Mismatching record length. Expected: %d. Got: %d", expectedRecordLen, actualRecordLen)
					} else {
						for key, expectedValue := range expectedRecord {
							actualValue, ok := actualRecord[key]
							if !ok {
								t.Errorf("Missing key '%s' at %d", key, rowIndex)
							}
							if actualValue != expectedValue {
								t.Errorf("Mismatching value for key '%s' at %d. Expected: %s. Got: %s", key, rowIndex, expectedValue, actualValue)
							}
						}
					}
				}
			}
		})
	}
}
