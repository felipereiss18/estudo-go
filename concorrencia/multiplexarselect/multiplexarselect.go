package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <- chan string  {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s falando: %d",pessoa, i)
		}
	}()
	return ch
}

func multiplexar(entrada1, entrada2 <- chan string) <- chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case sel := <- entrada1:
				ch <- sel
			case sel := <- entrada2:
				ch <- sel
			}
		}
	}()
	return ch
}

func main() {

	ch := multiplexar(falar("Ana"), falar("Maria"))
	println(<- ch, <- ch)
	println(<- ch, <- ch)
	println(<- ch, <- ch)
}
