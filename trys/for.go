package main

import "fmt"

func ForLoop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1

	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
}

func ThirdLoop() {
	for i := range 3 {
		fmt.Println("range", i)
	}
}

func InfiniteLoop() {
	for {
		fmt.Println("loop")
		break
	}
}

func Declarations() {
	for i := range 5 {
		fmt.Println("This loop worked ", i, " times.")
	}
}
