package main

import "fmt"

func main() {
	var point = 9

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("Not Bad")
		} else if point == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You can do it")
			if point == 0 {
				fmt.Println("Try Harder!")
			}
		}
	}
}
