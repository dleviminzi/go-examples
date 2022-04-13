package main

import (
	"fmt"
	"reflect"
)

type A struct {
	s1 int
	s2 int
	c  *CallService
}

type CallService struct {
	method func()
}

func NewA() A {
	return A{s1: 1, s2: 2}
}

func main() {
	a := NewA()
	fields := reflect.VisibleFields(reflect.TypeOf(a))
	for _, field := range fields {
		fmt.Println(field.Name, field.Type)
	}

}
