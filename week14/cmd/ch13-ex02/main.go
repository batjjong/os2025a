package main

import "fmt"

func abc(channel chan string) {
	channel <- "KBS1\n"
	channel <- "KBS2\n"
	channel <- "KBS3\n"
}

func def(channel chan string) {
	channel <- "MBC1\n"
	channel <- "MBC2\n"
	channel <- "MBC3\n"
}

func ghi(channel chan string) {
	channel <- "SBS1\n"
	channel <- "SBS2\n"
	channel <- "SBS3\n"
}

func main() {
	KBS := make(chan string)
	MBC := make(chan string)
	SBS := make(chan string)
	go abc(KBS)
	go def(MBC)
	go ghi(SBS)
	fmt.Print(<-KBS)
	fmt.Print(<-MBC)
	fmt.Print(<-SBS)

	fmt.Print(<-KBS)
	fmt.Print(<-MBC)
	fmt.Print(<-SBS)

	fmt.Print(<-KBS)
	fmt.Print(<-MBC)
	fmt.Print(<-SBS)

	fmt.Println()
}
