package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var str1, str2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
var m map[byte]byte

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	for {
		read, err := r.r.Read(p)
		if err == io.EOF {
			return 0, err
		} else if read == 0 {
			return 0, nil
		}
		for i := 0; i < read; i++ {
			p[i] = m[p[i]]
		}
		p = p[:read]
		fmt.Printf("%c", p)
	}
}

func initMap () {
	str1, str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	m = make(map[byte]byte)
	length := len(str1)
	for i := 0; i < length; i++ {
		m[(str1[i])] = str2[i]
	}
}


func main() {
	initMap()
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
}
