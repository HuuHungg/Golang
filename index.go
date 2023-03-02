package main

	import (
		"fmt"
	)
func main()  {
	// points := [...]int {7,10,5,6}
	// fmt.Printf("%v %T", points, points)
	
	var course [5]string
	course[0] = "alex"
	course[1] = "Tom"
	course[2] = "Yuh"
	course[3] = "Space"
	
	fmt.Printf("%v %T\n", course, course)
	fmt.Printf("%v %T\n", course[2], course)
	fmt.Println(len(course))
	
}

