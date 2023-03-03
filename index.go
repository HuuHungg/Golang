package main

	import (
		"fmt"
	)

func main () {
	number := 2
	switch  {
		case number <= 3:
			fmt.Println("Nho hon hoac bang 3")
			fallthrough
		case number <=5:
			fmt.Println("Nho hon hoac bang 5")
			fmt.Println("Nho hon hoac bang 5")
			break
			fmt.Println("Nho hon hoac bang 5")
			fallthrough // kết tiếp
		default:
			fmt.Println("Khong thuoc truong hop nao")
	}

		var i interface{} = 34.5
		switch i.(type) {
			case int:
					fmt.Println("Kieu du lieu int")
			case float64:
					fmt.Println("Kieu du lieu float64")
			default:
					fmt.Println("Khong biet kieu du lieu")
		}
	
	
}








