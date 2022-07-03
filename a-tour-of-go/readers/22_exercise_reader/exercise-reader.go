package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	var onlyChar byte = 'A'
	bSize := len(b)
	for i := 0; i < bSize; i++ {
		b[i] = onlyChar
	}
	return bSize, nil
}

func main() {
	reader.Validate(MyReader{})
}
