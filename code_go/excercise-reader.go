package main

import "golang.org/x/tour/reader"

//type red int

type myReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (a myReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	//var t red
	reader.Validate(myReader{})
}
