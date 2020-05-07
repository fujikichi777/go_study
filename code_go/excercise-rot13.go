package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	a io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) { //io.Reader
	if r.a == nil {
		print("再帰だ！")
	}
	n, err := r.a.Read(b) //io.Reader
	for i := range b {
		if 65 <= b[i] && b[i] <= 90 {
			b[i] = (b[i]-65+13)%26 + 65
		} else if 97 <= b[i] && b[i] <= 122 {
			b[i] = (b[i]-97+13)%26 + 97
		}
		if err == io.EOF {
			break
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
