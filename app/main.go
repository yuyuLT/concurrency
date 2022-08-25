package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, concurrency!")

	//
	message1 := make(chan string)
	message2 := make(chan string)

	go printMessage(message1)
	go printMessage2(message2)

	pmsg1 := <- message1
	pmsg2 := <- message2

	fmt.Println(pmsg1)
	fmt.Println(pmsg2)

	fmt.Println("finish program!")
}

func printMessage(message chan string) {
	time.Sleep(time.Second * 5)
	msg1 := "this is message 1"
	message <- msg1
}

func printMessage2(message chan string) {
	time.Sleep(time.Second * 5)
	msg2 := "this is message 2"
	message <- msg2
}