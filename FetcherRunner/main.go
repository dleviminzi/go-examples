package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// FetcherRunner allows us to write a function job that will take any structure
// that fetches data and creates entries or events or anything with that data.
// The job can then trigger the fetch and trigger the create function.
type FetcherCreater interface {
	Fetcher
	Creater
}

type Fetcher interface {
	Fetch()
}

type Creater interface {
	Create()
}

type common struct {
	Type string `json:"type" copier:"must"`
	A    int    `json:"a" copier:"must"`
	B    int    `json:"b" copier:"must"`
}

type A struct {
	common
	Name string `json:"name,omitempty"`
}

type B struct {
	common
	Fl float32 `json:"fl,omitempty"`
}

// Note that the fetch is not returning a type specific to what it is fetching.
// To do that would break our interface. Until generics we must specify the
// return type in the interface.
func (a *A) Fetch() {
	jsonFileA, _ := os.Open("./specA.json")
	byteValue, _ := ioutil.ReadAll(jsonFileA)
	json.Unmarshal(byteValue, a)
}

func (a *A) Create() {
	fmt.Println("Creating entry for: ")
	fmt.Printf("%#v \n", a)
}

func (b *B) Fetch() {
	jsonFileB, _ := os.Open("./specB.json")
	byteValue, _ := ioutil.ReadAll(jsonFileB)
	json.Unmarshal(byteValue, b)
}

func (b *B) Create() {
	fmt.Println("Creating entry for: ")
	fmt.Printf("%#v \n", b)
}

func job(j FetcherCreater) {
	j.Fetch()
	j.Create()
}

func main() {
	// Any new structure can be added by simply using a new job() call without
	// so long as it is fetcher creater
	job(&A{})
	job(&B{})
}
