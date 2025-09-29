package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	name := "Go developer"

	fmt.Println("Git/Github for", name)
	fmt.Println(math.Floor(2.71))

	upper := strings.ToUpper(name)
	fmt.Println("Upper:", upper)
}
