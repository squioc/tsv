package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/squioc/tsv/pkg/tsv"
)

var (
	usageLine string = "Usage: %v [options]\n"
)

func formatRecords(records [][]string, widths []int) []string {

	nbColumns := len(widths)
	formatters := make([]string, nbColumns)
	for index, width := range widths {
		formatters[index] = fmt.Sprintf("%%-%ds", width)
	}

	lines := make([]string, len(records))
	for index, record := range records {
		var buffer bytes.Buffer
		for cellIndex, cell := range record {
			buffer.WriteString(fmt.Sprintf(formatters[cellIndex], cell))
			buffer.WriteString("  ")
		}
		lines[index] = buffer.String()
	}
	return lines
}

func widthsFromRecords(records [][]string) []int {

	if len(records) == 0 {
		return []int{}
	}

	nbColumns := len(records[0])
	widths := make([]int, nbColumns)

	for _, record := range records {
		for columnIndex := 0; columnIndex < nbColumns; columnIndex++ {
			valueLen := len(record[columnIndex])
			if len(widths) < nbColumns {
				widths = append(widths, valueLen)
			}
			if widths[columnIndex] < valueLen {
				widths[columnIndex] = valueLen
			}
		}
	}

	return widths
}

func usage() {
	fmt.Fprintf(os.Stderr, usageLine, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var lines int
	flag.Usage = usage
	flag.IntVar(&lines, "n", 0, "The number of lines to read")

	// parses arguments
	flag.Parse()
	args := flag.Args()

	var reader io.Reader
	if len(args) == 0 {
		reader = os.Stdin
	} else {
		file, err := os.Open(args[0])

		if err != nil {
			log.Fatal(err)
		}
		reader = file
		defer file.Close()
	}

	tsvReader := tsv.NewReader(reader)
	records, err := tsvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	widths := widthsFromRecords(records)
	rows := formatRecords(records, widths)
	for index := 0; index < len(rows) && (lines == 0 || index < lines); index++ {
		fmt.Println(rows[index])
	}
}
