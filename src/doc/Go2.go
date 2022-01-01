package main

import (
	"fmt"
	"io"
	"os"
)

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '0':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0

}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '=', '#', '+', '%':
		return true
	}
	return false
}

/**
jump to another
*/
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

//func (file *File) Write(b []byte) (n int, err error)
func nextInt(b []byte, i int) (int, int) {
	if len(b) < 10 {
	}
	x := 0
	return x, i
}

//func ReadFull(r Reader, buf []byte) (n int, err error) {
//	for len(buf) > 0 && err == nil {
//		var nr int
//		nr, err = r.Read(buf)
//		n += nr
//		buf = buf[nr:]
//	}
//}

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run after all finish

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append function
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // if we return , f will be close
		}
	}
	return string(result), nil // if we return , f will be close

}
func main() {

	b := unhex('6')
	println(b)

	escape := shouldEscape('?')
	println(escape)

	str1 := []byte{'a'}
	str2 := []byte{'a', 'b'}
	compare := Compare(str1, str2)
	println(compare)

	var t interface{}
	t = 1
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case bool:
		fmt.Printf("boolean %t \n", t)
	case int:
		fmt.Printf("integer %t \n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t \n", *t) // t is *bool type
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t is * int type
	}

	by := []byte{'a'}
	i1, i2 := nextInt(by, 1)
	println(i1, i2)

	contents, err := Contents("")
	println(contents ," ", err)

	//后进先出（LIFO）的顺序执行
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}



}
