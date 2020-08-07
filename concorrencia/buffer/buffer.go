package main

import (
	"fmt"
	"time"
)

func buffer(ch chan int){
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!!!")
	ch <- 6
}

func main() {

	ch := make(chan int, 3)

	go buffer(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println("Fim")
}
