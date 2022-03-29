/*
The goal here is to create a way to avoid doing writing the unmarshalling and
marshalling immediately before and after logic in every step of a pipeline.
This way, we have a generic unmarshal and marshal wrapper that takes the pipeline
step and uses the input/output type to infer the type to unmarshal and marshal
back into.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("starting...")

	callA := UnmarsalWrapper(runA)
	callB := UnmarsalWrapper(runB)

	a, _ := json.Marshal(EventTypeA{A: 1000})
	b, _ := json.Marshal(EventTypeB{B: "Daniel"})

	ra, _ := callA(a)
	rb, _ := callB(b)

	fmt.Println(string(ra))
	fmt.Println(string(rb))
}

type EventTypeA struct {
	A int `json:"A"`
}

func runA(event EventTypeA) (EventTypeA, error) {
	// isolated logic happening to unmarshalled type
	fmt.Println("Here in Event Type A: ", event.A)
	return event, nil
}

type EventTypeB struct {
	B string `json:"B"`
}

func runB(event EventTypeB) (EventTypeB, error) {
	// isolated logic happening to unmarshalled type
	fmt.Println("Here in Event Type B: ", event.B)
	return event, nil
}

type runner[R any] func(event R) (R, error)

func UnmarsalWrapper[R any](run runner[R]) func(a []byte) ([]byte, error) {
	return func(a []byte) ([]byte, error) {
		var s R
		err := json.Unmarshal(a, &s)
		if err != nil {
			fmt.Println(err)
		}

		run(s)

		b, _ := json.Marshal(s)
		return b, nil
	}
}
