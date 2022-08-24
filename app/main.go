package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	fmt.Println("Hello, concurrency!")

	var wg sync.WaitGroup
	
	wg.Add(2)

	go printMessage(&wg)
	go printMessage2(&wg)

	wg.Wait()

	fmt.Println("finish program!")
}

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("this is message 1")
}

func printMessage2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("this is message 2")
}