package tsv

type TsvError string

func (e TsvError) Error() string {
	return string(e)
}
