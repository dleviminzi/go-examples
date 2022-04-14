package main

import (
	"fmt"
	"reflect"
)

type A interface {
	CallService
	LogService
}

type instanceOfA struct {
	Name string
	C    CallService
	L    LogService
}

func NewInstanceOfA() instanceOfA {
	return instanceOfA{Name: "bad"}
}

func NewGoodInstanceOfA() instanceOfA {
	f1 := func() { fmt.Print("yo") }
	f2 := func() { fmt.Print("yoyoyo") }

	return instanceOfA{Name: "good", C: CallService{f1}, L: LogService{f2}}
}

type CallService struct {
	method func()
}

type LogService struct {
	method func()
}

func print() {
	fmt.Print("hi")
}

func main() {
	a := NewInstanceOfA()
	fmt.Println("Checking nils in: ", a.Name)
	NilFields(a)

	b := NewGoodInstanceOfA()
	fmt.Println("\nChecking nils in: ", b.Name)
	NilFields(b)
}

func NilFields[T any](a T) {
	s := reflect.ValueOf(&a).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		v := reflect.ValueOf(f.Interface())
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			fmt.Println("Encountered nil")
		}
	}
}
