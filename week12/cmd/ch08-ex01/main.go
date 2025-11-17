package main

import "fmt"

type subscriber struct {
	name  string
	price int
}

func applyPrice(s *subscriber) {
	s.price = 10000
	s.name = "Kim inha"
}

func main() {
	var s1 subscriber
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
}
