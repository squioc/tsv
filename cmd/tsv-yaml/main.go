package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/squioc/tsv/pkg/tsv"
)

var (
	usageLine string = "Usage: %v [options]\n"
)

func readAndPrintRecords(reader io.Reader, fieldnames []string, encoder func(interface{}) ([]byte, error)) {
	tsvReader := tsv.NewDictReader(reader, fieldnames)

	records, err := tsvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	} else {
		var encoded []byte
		encoded, err = encoder(records)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Print(string(encoded))
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageLine, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var columns string
	flag.Usage = usage
	flag.StringVar(&columns, "columns", "", "The list of name of each column (separated by comma)")

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
	readAndPrintRecords(reader, fieldnames, yaml.Marshal)
}
