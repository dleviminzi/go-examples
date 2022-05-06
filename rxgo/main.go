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
}

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		mediaTypes := []string{"call", "chat", "email"}
		names := []string{"daniel", "laura", "jonathan", "harel", "jude", "landon", "brandon", "azeez", "tyler", "simon", "max", "moni", "jacob"}

		for i, name := range names {
			ex := Ex{
				MediaType: mediaTypes[i%3],
				Name:      name,
				Duration:  i * 10,
			}

			ch <- rxgo.Of(ex)
		}

		close(ch)
	}()
	// Create a regular Observable
	observable := rxgo.FromChannel(ch).
		Filter(verify).
		Map(process).
		BufferWithCount(3)

	// observe
	for item := range observable.Observe() {
		if item.Error() {
			fmt.Println("do something. error :: ", item.E)
		}
		fmt.Println("items are :: ", item.V)
	}
}

func verify(i any) bool {
	if i.(Ex).MediaType == "" {
		return false
	}
	return true
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
