package main

import "fmt"

func main() {

	const (
		square        = "kotak"
		isToday  bool = true
		numeric  uint = 1
		floatNum      = 2.2
	)

	fmt.Println("======")
	fmt.Println(square)
	fmt.Println(isToday)
	fmt.Println(floatNum)

	const (
		a = "constant"
		b
	)

	fmt.Println(a)
	fmt.Println(b)
	
	fmt.Println("=======")

	const (
		today string = "monday"
		sekarang
		isToday2 = true
	)

	fmt.Println(today)
	fmt.Println(sekarang)
	fmt.Println(isToday)

	fmt.Println("=======")

	const one, two = 1, 2
	const three, four string = "three", "four"

	fmt.Println(one, two)
	fmt.Println(three, four)

}
