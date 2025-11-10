package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile" //go get github.com/headfirstgo/datafile
)

func main() {
	weight, err := datafile.GetFloats("meatweight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for i := 0; i < len(weight); i++ {
		sum = sum + weight[i]
	}
	weeks := float64(len(weight))
	fmt.Println("평균 : ", sum/weeks)
}
