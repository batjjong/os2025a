package main

import (
	"fmt"
)

func main() {

	//number := [3]int{-1, 11, 7}
	//for i := 0; i < len(number); i++ {
	//	fmt.Println(i, number[i])
	//}

	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}
