package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[%v] %s", e.When.Format("2006-01-02 15:04:05.123"), e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(err) // at 2022-07-02 21:09:38.717181 +0900 KST m=+0.000231001, it didn't work
	}

}
