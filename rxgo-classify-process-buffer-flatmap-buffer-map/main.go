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
				MediaType: mediaTypes[i%(len(mediaTypes)-1)],
				Name:      name,
				Duration:  i%len(names)*10 + len(names),
			}

			raw <- rxgo.Of(ex)
		}
		close(raw)
	}()

	// split into different streams
	obs := rxgo.FromChannel(raw).GroupBy(4, classify)

	items := []rxgo.Item{}
	for item := range obs.Observe() {
		items = append(items, item)
	}

	// merge all streams
	final := []rxgo.Observable{
		items[0].V.(rxgo.Observable).Map(process).BufferWithCount(3).FlatMap(enrich[*A]).BufferWithCount(2).Map(transmit[A]),
		items[1].V.(rxgo.Observable).Map(process).BufferWithCount(3).FlatMap(enrich[*B]).BufferWithCount(2).Map(transmit[B]),
		items[2].V.(rxgo.Observable).Map(process).BufferWithCount(3).FlatMap(enrich[*C]).BufferWithCount(2).Map(transmit[C]),
		items[3].V.(rxgo.Observable).Map(trashF),
	}
	finalOb := rxgo.Merge(final)

	// observe (run) all merged streams
	for item := range finalOb.Observe() {
		if item.E != nil {
			fmt.Println(item.E.Error())
		}
		// fmt.Println(fmt.Sprintf("%+v\n", item.V))
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

type HasExternal interface {
	SetExternal(string)
}

func transmit[T any](ctx context.Context, i any) (any, error) {
	toTransmit := []T{}
	arr := i.([]any)

	for _, a := range arr {
		transmitable := a.(*T)
		toTransmit = append(toTransmit, *transmitable)
	}

	fmt.Println(toTransmit)
	return toTransmit, nil
}

func enrich[T HasExternal](i rxgo.Item) rxgo.Observable {
	enriched := []T{}

	exArr := i.V.([]any)
	for _, x := range exArr {
		z := x.(T)
		z.SetExternal("populated")

		enriched = append(enriched, z)
	}

	return rxgo.Just(enriched)()
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
		return &A{e: ex}, nil
	case "chat":
		ex.Duration = 15
		return &B{e: ex}, nil
	case "email":
		ex.Duration = 25
		return &C{e: ex}, nil
	}

	return ex, nil
}

type A struct {
	e Ex
}

func (a *A) SetExternal(s string) {
	a.e.External = s
}

type B struct {
	e Ex
}

func (b *B) SetExternal(s string) {
	b.e.External = s
}

type C struct {
	e Ex
}

func (c *C) SetExternal(s string) {
	c.e.External = s
}
