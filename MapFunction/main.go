package main

import "fmt"

type D struct {
	name   string
	number float64
}

func (a D) Run() {
	fmt.Println(a.name, ": Creating d")
}

type De struct {
	name   string
	number int
}

func (b De) Run() {
	fmt.Println(b.name, ": Creating delivery journal entry")
}

func FetchD(jobType string) Runner {
	return D{name: jobType, number: 1.2}
}

func FetchDe(jobType string) Runner {
	return De{name: jobType, number: 1}
}

type Runner interface {
	Run()
}

type RunnerReturner func(string) Runner

func CallRunOnMap(m map[string]RunnerReturner) {
	for k, v := range m {
		v(k).Run()
	}
}

func main() {
	ma := map[string]RunnerReturner{"d": FetchD, "de": FetchDe}

	CallRunOnMap(ma)
}
