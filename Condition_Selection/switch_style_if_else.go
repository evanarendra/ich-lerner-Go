package main

import "fmt"

func main() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("You need to learn more")
		}
	}
}
