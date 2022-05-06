package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

type Ex struct {
	MediaType string
	Name      string
	Duration  int
	External  string
}

var enrichedToProcess = make(chan rxgo.Item)

func main() {
	rawToEnrich := make(chan rxgo.Item)
	go func() {
		mediaTypes := []string{"call", "chat", "email"}
		names := []string{"daniel", "laura", "jonathan", "harel", "jude", "landon", "brandon", "azeez", "tyler", "simon", "max", "moni", "jacob"}

		for i, name := range names {
			ex := Ex{
				MediaType: mediaTypes[i%3],
				Name:      name,
				Duration:  i * 10,
			}

			rawToEnrich <- rxgo.Of(ex)
		}

		close(rawToEnrich)
	}()

	// Create a regular Observable
	ob1 := rxgo.FromChannel(rawToEnrich).
		BufferWithCount(3).
		Map(enrich)

	go func() {
		for range ob1.Observe() {
		}
		close(enrichedToProcess)
	}()

	observable := rxgo.FromChannel(enrichedToProcess).
		Map(process)

	// observe
	for item := range observable.Observe() {
		if item.Error() {
			fmt.Println("do something. error :: ", item.E)
		}
		fmt.Println("item :: ", item.V)
	}
}

func verify(i any) bool {
	if i.(Ex).MediaType == "" {
		return false
	}
	return true
}

func enrich(ctx context.Context, i any) (any, error) {
	exArr := i.([]any)
	for _, x := range exArr {
		z := x.(Ex)
		z.External = "populated"

		enrichedToProcess <- rxgo.Of(z)
	}

	return i, nil
}

func process(ctx context.Context, i any) (any, error) {
	var ex Ex
	var ok bool

	if ex, ok = i.(Ex); !ok {
		return ex, errors.New("could not assert type Ex")
	}

	switch ex.MediaType {
	case "call":
		ex.Duration = 10
	case "chat":
		ex.Duration = 15
	case "email":
		ex.Duration = 25
	}

	return ex, nil
}
