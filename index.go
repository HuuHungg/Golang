package main

import (
	"fmt"
)

// Kiểu dữ liệu khai báo Enum tập hợp nhiều hằng số lại với nhau

const (
	// iota chỉ sử dụng trong một block 
	red = iota
	yellow 
	blue 
	black 
)
const (
	// iota chỉ sử dụng trong một block 
	red1 = iota + 5
	yellow1 

)


func main() {
	const i = "Hello"
	const a = 12
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", a, a)

	fmt.Printf("%v %T\n", red, red)
	fmt.Printf("%v %T\n", yellow, yellow)
	fmt.Printf("%v %T\n", blue, blue)
	fmt.Printf("%v %T\n", black, black)
	fmt.Printf("%v %T\n", red1, red1)
	fmt.Printf("%v %T\n", yellow1, yellow1)
}

