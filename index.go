package main

	import (
		"fmt"
	)
func main()  {
	// slice 
	a := []int{2,5,10,40,50,60}
	b := a
	b[1] = 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(cap(a)) // Sức chứa
	
}




