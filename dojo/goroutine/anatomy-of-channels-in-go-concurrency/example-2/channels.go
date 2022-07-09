package main

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[squares] push: %d's square\n", i)
		c <- i * i
		// time.Sleep(time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan int)

	go squares(c)

	time.Sleep(time.Second)

	for {
		v, ok := <-c
		if ok == false {
			break
		}
		fmt.Println(v, ok)
	}

	// 이 루프는 실행이 안 된다.
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("---END---")
}
