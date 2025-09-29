package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//float32 = float
	//float64 = double

	//var name string       선언
	//name = "Kim Inha"     할당

	//var name = "Kim Inha" 선언 + 할당

	name := "Kim Inha" //   선언 + 할당

	var f float32
	var d float64

	f = 3.21
	d = 3.21

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(name))

	fmt.Println(math.Ceil(2.31))
	fmt.Println(strings.Title(name))
}
