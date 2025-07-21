package main

import (
	"fmt"
	"time"
)

func callfunc(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func PrintNum(ch chan int) {
	for i := 0; i < 5; i++ {
		go func() {
			callfunc(ch)
		}()
	}
	time.Sleep(time.Second * 2)

}

var counter = 0

func getval() {
	for i := 0; i < 5; i++ {
		counter++
	}
}

func main() {

	go getval()
	go getval()
	go getval()
	time.Sleep(time.Second * 2)
	fmt.Println(counter)

	ch := make(chan int, 5)

	PrintNum(ch)

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Second * 5)

	//dummy docker file
	//from gplang.1.21.1-alpine
	//workdir ./app
	//copy . .
	//run go -o build main.go
	//cmd ["main.go"]

	//service
	//image
	//container name

}
