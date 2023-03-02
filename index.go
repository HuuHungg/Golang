package main

	import (
		"fmt"
	)
func main()  {
	// slice 
	a := []int{2,5,10,40,50,60,80,90}
	b := a[:] // hai dấu chấm  đứng một mình là lấy hết các phần tử của a
	c := a[3:] // con trỏ bắt đầu từ index vị trí số 3
	d := a[:5] 
	e := a[3:5]
	fmt.Printf("a %v, %v, %v\n", a, len(a),cap(a))	
	fmt.Printf("b %v, %v, %v\n", b, len(b),cap(b))	
	fmt.Printf("c %v, %v, %v\n", c, len(c),cap(c))	
	fmt.Printf("d %v, %v, %v\n", d, len(d),cap(d))	
	fmt.Printf("e %v, %v, %v\n", e, len(e),cap(e))	

}





