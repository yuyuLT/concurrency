package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, concurrency!")
	go delay()

	time.Sleep(time.Second * 1)
}

func delay() {
	fmt.Println("Hello, goroutine!")
}
