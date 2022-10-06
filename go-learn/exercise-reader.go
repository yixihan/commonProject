package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{
}

func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

//type MyReader struct{
//	r strings.Reader
//	str string
//}
//
// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
//func (mr MyReader) Read (b []byte) (int, error) {
//
//	length := len(b)
//	var sb strings.Builder
//	for sb.Len() < length {
//		sb.WriteString(mr.str)
//	}
//	fmt.Println(sb)
//
//	for {
//		mr.r.Reset(sb.String())
//		_, err := mr.r.Read(b)
//		fmt.Printf("%c", b)
//		if err == io.EOF {
//			return 0, err
//		}
//	}
//}
//
//func main() {
//	reader.Validate(MyReader{
//		r: strings.Reader{},
//		str: "A",
//	})
//}

