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

func classify(item rxgo.Item) int {
	switch item.V.(Ex).MediaType {
	case "call":
		return 0
	case "chat":
		return 1
	case "email":
		return 2
	default:
		return 3
	}
}

func main() {
	raw := make(chan rxgo.Item)
	go func() {
		mediaTypes := []string{"call", "chat", "email", "sms"}
		names := []string{"daniel", "laura", "jonathan", "harel", "jude", "landon", "brandon", "azeez", "tyler", "simon", "max", "moni", "jacob", "daniel", "laura", "jonathan", "harel", "jude", "landon", "brandon", "azeez", "tyler", "simon", "max", "moni", "jacob"}
		fmt.Println(len(names))

		for i, name := range names {
			ex := Ex{
				MediaType: mediaTypes[i%len(mediaTypes)],
				Name:      name,
				Duration:  i%len(names)*10 + len(names),
			}

			raw <- rxgo.Of(ex)
		}
		close(raw)
	}()

	// split into different streams
	obs := rxgo.FromChannel(raw).Map(process, rxgo.WithBufferedChannel(10), rxgo.WithCPUPool()).GroupBy(4, classify, rxgo.WithBufferedChannel(10), rxgo.WithCPUPool())

	items := []rxgo.Item{}
	for item := range obs.Observe() {
		items = append(items, item)
	}

	// merge all streams
	final := []rxgo.Observable{
		items[0].V.(rxgo.Observable).BufferWithCount(3).Map(enrich),
		items[1].V.(rxgo.Observable).BufferWithCount(3).Map(enrich),
		items[2].V.(rxgo.Observable).BufferWithCount(3).Map(enrich),
		items[3].V.(rxgo.Observable).Map(trashF),
	}
	finalOb := rxgo.Merge(final)

	// observe (run) all merged streams
	for item := range finalOb.Observe() {
		if item.E != nil {
			fmt.Println(item.E.Error())
		}
		fmt.Println(fmt.Sprintf("%+v\n", item.V))
	}
}

func verify(i any) bool {
	if i.(Ex).MediaType == "" {
		return false
	}
	return true
}
func trashF(ctx context.Context, i any) (any, error) {
	ex := i.(Ex)
	fmt.Println(fmt.Sprintf("ex: %+v\n", ex))
	return "", nil
}

func enrich(ctx context.Context, i any) (any, error) {
	enriched := []Ex{}

	exArr := i.([]any)
	for _, x := range exArr {
		z := x.(Ex)
		z.External = "populated"

		enriched = append(enriched, z)
	}

	return enriched, nil
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
