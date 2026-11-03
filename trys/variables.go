package main

import "fmt"

func GetVariables() {
	a := "initial"
	fmt.Println(a)

	var b, c int = 1, 3
	fmt.Println(b, " ", c)

	var d bool = true
	fmt.Println(!d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
