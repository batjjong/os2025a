package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"reflect"
)

func main() {
	/*
		var length float64 = 3.2
		var width int = 2

		//fmt.Println("Area is", width*length)
		//타입이 달라서 실행이 안됨
		fmt.Println("면적", float64(width)*length)
		fmt.Println("길이 > 너비?", length > float64(width))

		fmt.Println("원본", reflect.TypeOf(length))
		fmt.Println("형변환", reflect.TypeOf(float64(length)))

		var now time.Time = time.Now()
		//var month time.Month = now.Month()
		month := now.Month()
		var day int = now.Day()
		fmt.Println(month)
		fmt.Println(day)

		univ := "Go$ rocks$"
		changer := strings.NewReplacer("$", "!")
		changed := changer.Replace(univ)
		fmt.Println(changed)
	*/
	/*
		//쉐도잉 shadowing
		//var fmt string = "inha"

		var int int = 7 //int int를 선언(쉐도잉)
		var kim int = 0 //int를 쓰지 못함
		fmt.Println(int)
	*/

	//a, b := 10,9 //동시에 선언 할당 가능

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') //에러 무시
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) //프로그램 중지하고 시간+에러 메세지 출력
	}

	i = strings.TrimSpace(i)                //trim
	score, err := strconv.ParseFloat(i, 64) //형변환
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
	fmt.Println(i)

}
