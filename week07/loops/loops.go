package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 3.2
	var width int = 2

	//fmt.Println("Area is", width*length)
	//타입이 달라서 실행이 안됨
	fmt.Println("면적", float64(width)*length)
	fmt.Println("길이 > 너비?", length > float64(width))

	fmt.Println("원본", reflect.TypeOf(length))
	fmt.Println("형변환", reflect.TypeOf(float64(length)))

}
