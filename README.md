# TSV

TSV is a library (and a set of tool) to manipulate Tabular separated values.

It is inspired by [solidsnack/tsv](github.com/solidsnack/tsv) and thoughts from Donald Merand: [TSV - The Best Spreadsheet Format](https://donaldmerand.com/code/2011/09/20/tsv-the-best-spreadsheet-format.html)

## Installation

    $ go get github.com/squioc/tsv/pkg/tsv

## Tools companion

TSV holds a set of tools to manipulate Tabular data.

### tsv-yaml

Converts tsv records into yaml

#### Installation

    $ go get github.com/squioc/tsv/cmd/tsv-yaml

#### Usage

    $ tsv-yaml --columns first,second
    cell1     cell2
    ^D
    - first: cell1
      second: cell2

    $ cat cells.tsv | tsv-yaml --columns first,second
    - first: cell1
      second: cell2

### tsv-json

Converts tsv records into json

#### Installation

    $ go get github.com/squioc/tsv/cmd/tsv-json

#### Usage

    $ tsv-json --columns first,second
    cell1     cell2
    ^D
    [{"first":"cell1","second":"cell2"}]

    $ cat cells.tsv | tsv-json --columns first,second
    [{"first":"cell1","second":"cell2"}]
