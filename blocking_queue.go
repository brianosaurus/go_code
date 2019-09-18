package main

import (
	"fmt"
	"time"
)

var queue chan interface{}

func routine() {
	for msg := range queue {
		fmt.Println(msg)
		fmt.Println("blocking in function")
	}
}

func main() {
	// a blocking queue
	queue = make(chan interface{})

	go routine()

	queue <- "foo"
	queue <- "bar"
	fmt.Println("Sleeping for a sec")
	time.Sleep(2 * time.Second)
	fmt.Println("waking up")
	queue <- "far"
	queue <- 123456
	fmt.Println("Sleeping again")
	time.Sleep(2 * time.Second)
}
