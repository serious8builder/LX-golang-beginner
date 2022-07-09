package main

import (
	"fmt"
)

/*
채널에 write 만한 상태인 경우
fatal error: all goroutines are asleep - deadlock!

*/

func greet(c chan string) {
	//time.Sleep(10000)
	fmt.Println("Hello " + <-c)
}

func main() {
	c := make(chan string)

	go greet(c)
	c <- "Ryan"
	fmt.Println("---END---")

}
