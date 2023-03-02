package main

import (
	"fmt"
)

func main() {
	// Text
	// Byte -> UTF-8
	// Rune -> UTF-32

	var a string = "こんにちは世界"
	var b []byte = []byte(a)
	fmt.Printf("%v %T\n", b, b)
	
	// Chuyển kiểu
	var c byte = 'H'
	fmt.Printf("%v, %T\n", string(c), c)

	var d rune = '世'
	fmt.Printf("%v, %T\n", string(d), d)
}

