package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Println(j, " ")
		}
		fmt.Println()
	}
}
