package main

import "fmt"

func main() {

	var point = 8

	if point == 10 {
		fmt.Println("Passed the test with perfect")
	} else if point > 5 {
		fmt.Println("Passed the test")
	} else if point == 4 {
		fmt.Println("Almost passed the test")
	} else {
		fmt.Println("failed, your point is %d\n", point)
	}
}
