package tsv

import (
	"io"
)

const ErrRowSizeOverflow TsvError = "The row size exceed the size of fields names"

// DictReader provides a dictionary reader for TSV formatted data
type DictReader struct {
	r          *Reader
	fieldnames []string
}

// NewReader creates a new reader
func NewDictReader(r io.Reader, fieldnames []string) *DictReader {
	return &DictReader{
		r:          NewReader(r),
		fieldnames: fieldnames,
	}
}

// Read reads one row and returns it as a map of string
func (r *DictReader) Read() (map[string]string, error) {
	row, err := r.r.Read()
	if err != nil {
		return nil, err
	}

	if len(row) > len(r.fieldnames) {
		return nil, ErrRowSizeOverflow
	}

	record := make(map[string]string)
	for cellIndex, value := range row {
		name := r.fieldnames[cellIndex]
		record[name] = value
	}
	return record, nil
}

func (r *DictReader) ReadAll() ([]map[string]string, error) {
	var records []map[string]string
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
