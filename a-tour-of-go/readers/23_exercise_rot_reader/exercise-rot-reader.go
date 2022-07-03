package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	for {
		n, error := rot.r.Read(b)

		if error != nil {
			return n, error
		}

		for i := 0; i < n; i++ {
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] = 'A' + (b[i]-'A'+13)%26
			} else if b[i] >= 'a' && b[i] <= 'z' {
				b[i] = 'a' + (b[i]-'a'+13)%26
			}
		}
		return n, error
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
