# TSV

TSV is a library (and a set of tool) to manipulate Tabular separated values.

It is inspired by [solidsnack/tsv](https://github.com/solidsnack/tsv) and thoughts from Donald Merand: [TSV - The Best Spreadsheet Format](https://donaldmerand.com/code/2011/09/20/tsv-the-best-spreadsheet-format.html)

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

### tsv-jsonline

Converts tsv records into [jsonline](http://jsonlines.org/)

#### Installation

    $ go get github.com/squioc/tsv/cmd/tsv-jsonline

#### Usage

    $ tsv-jsonline --columns first,second
    cell1     cell2
    ^D
    {"first":"cell1","second":"cell2"}

    $ cat cells.tsv | tsv-jsonline --columns first,second
    {"first":"cell1","second":"cell2"}

### tsv-fmt

Format the tsv input in order to vertically align columns

#### Installation

    $ go get github.com/squioc/tsv/cmd/tsv-fmt

#### Usage

    $ tsv-fmt cells.tsv 
    cell1     cell2
    columns1  columns2

    $ cat cells.tsv | tsv-fmt
    cell1     cell2
    columns1  columns2
