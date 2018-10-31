package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main() {
	var w io.Writer

	w = os.Stdout

	f := w.(*os.File)
	fmt.Println(f)

	c, ok := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	fmt.Println(c)
	fmt.Println(ok)

	var v interface{}
	v = 1234
	switch v.(type) {
	case string:
		fmt.Printf("The string is %s", v.(string))
	case int, uint, int8:
		fmt.Printf("The integer is %d", v)
	default:
		fmt.Printf("Unknown value: type=%T", v)
	}
}
