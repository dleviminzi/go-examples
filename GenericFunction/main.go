package main

// This playground uses a development build of Go:
// devel go1.19-c9b60632eb Fri Mar 4 14:10:38 2022 +0000

import "fmt"

func getEventPointers[T any](events []T) []*T {
	var eventPointers []*T

	for _, i := range events {
		temp := i
		eventPointers = append(eventPointers, &temp)
	}

	return eventPointers
}

type eventA struct {
	details string
}

type eventB struct {
	details string
	id      int
}

func main() {
	bunchOfEventsA := []eventA{eventA{details: "a"}, eventA{details: "b"}}
	pA := getEventPointers(bunchOfEventsA)
	fmt.Println(pA)
	
	bunchOfEventsB := []eventB{eventB{details: "a", id: 1}, eventB{details: "b", id:3}}
	pB := getEventPointers(bunchOfEventsB)
	fmt.Println(pB)
}