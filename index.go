package main

	import (
		"fmt"
	)

func main () {
	number := 90
	guess := 85
	
	if !(number < 10 || !false || guess > 100) {
		fmt.Println("Con số bạn đoán nhỏ hơn 10 hoặc là nhà hơn 100")
	} else {
		if guess < number {
			fmt.Println("Con so muc tieu lon hon con so ban dau")
		} else if guess > number {
			fmt.Println("Con so muc tieu nho hon con so ban dau")
		} else guess == number {
			fmt.Println("Ban da doan dung")
		}
	}


	
}








