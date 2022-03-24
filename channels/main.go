package main

import "fmt"

func test(messages chan string) {
	messages <- "ping"
}

func main() {
	messages := make(chan string)

	go test(messages)

	msg := <-messages
	fmt.Println(msg)
}
