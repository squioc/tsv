package tsv

import (
	"bufio"
	"io"
	"strings"
)

// Reader provides a reader for TSV formatted data
type Reader struct {
	r *bufio.Reader
}

// NewReader creates a new reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

// readLine reads one line at a time
func (r *Reader) readLine() (string, error) {
	line, err := r.r.ReadSlice('\n')
	if err != nil {
		return "", err
	}

	return normalizeLine(string(line)), nil
}

// Read reads and transform one row
func (r *Reader) Read() ([]string, error) {
	line, err := r.readLine()
	if err != nil {
		return []string{}, err
	}

	cells := strings.Split(line, "\t")
	row := make([]string, 0, len(cells))
	for _, token := range cells {
		row = append(row, Unescape(token))
	}

	return row, nil
}

// ReadAll reads all records
func (r *Reader) ReadAll() ([][]string, error) {
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}
