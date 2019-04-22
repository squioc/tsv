package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/squioc/tsv/pkg/tsv"
)

var (
	usageLine string = "Usage: %v [options]\n"
)

func readAndPrintLines(reader io.Reader, fieldnames []string, lines int, encoder func(interface{}) ([]byte, error)) {
	tsvReader := tsv.NewDictReader(reader, fieldnames)

	for i := 0; lines == 0 || i < lines; i++ {
		records, err := tsvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		} else {
			var encoded []byte
			encoded, err = encoder(records)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(string(encoded))
			}
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageLine, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var columns string
	var lines int
	flag.Usage = usage
	flag.StringVar(&columns, "columns", "", "The list of name of each column (separated by comma)")
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

	fieldnames := strings.Split(columns, ",")
	readAndPrintLines(reader, fieldnames, lines, json.Marshal)
}
