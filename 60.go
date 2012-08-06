package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error){
	d, err := r.r.Read(b)
	if d > 0 {
		for i := 0; i < d; i++ {
			// A = 65, Z = 90
			// a = 97, z = 122
			if b[i] >= 65 && b[i] <= 91 {
				b[i] -= 65
				b[i] += 13
				b[i] %= 26
				b[i] += 65
			} else if b[i] >= 97 && b[i] <= 122 {
				b[i] -= 97
				b[i] += 13
				b[i] %= 26
				b[i] += 97
			}
		}
	}
	return d, err
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}