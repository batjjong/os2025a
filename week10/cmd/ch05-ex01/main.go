package main

import (
	"fmt"
)

func main() {
	//arrayBool := [2]bool{true, false} //배열 리터럴
	arrayInt := [3]int{-1, 11, 7}
	for i := 0; i < len(arrayInt); i++ {
		//fmt.Println(i, arrayBool[i]) //runtime error 인덱스 값이 맞지 않음
		fmt.Println(i, arrayInt[i])
		fmt.Println()
	}

	//fmt.Println(reflect.TypeOf(arrayBool))
	//fmt.Printf("%#v\n", arrayBool)
	//fmt.Printf("%#v\n", arrayInt)
}
