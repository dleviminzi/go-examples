package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/copier"
)

// structs ---------------------------------------------------------------------
type everything struct {
	common
	specificToA
	specificToB
}

type common struct {
	Type string `json:"type" copier:"must"`
	A    int    `json:"a" copier:"must"`
	B    int    `json:"b" copier:"must"`
}

type specificToA struct {
	Name string `json:"name,omitempty"`
}

type fullA struct {
	common
	specificToA
}

type specificToB struct {
	Fl float32 `json:"fl,omitempty"`
}

type fullB struct {
	common
	specificToB
}

// methods ---------------------------------------------------------------------
func RunA(a fullA) {
	fmt.Println("code specific to specificA")
	fmt.Println(a.Name)
}

func RunB(b fullB) {
	fmt.Println("code specific to specificB")
	fmt.Println(b.Fl)
}

// RunE determines the specific type of the generic everything struct. It then 
// copies the struct into a smaller specific struct (potentially one with methods)*
// and then passes it to the appropriate run function. 
func RunE(e *everything) {
	switch e.Type {
	case "A":
		var fA fullA
		copier.Copy(&fA.common, e.common)
		copier.Copy(&fA.specificToA, e.specificToA)
		fmt.Printf("\n %#v \n", fA)
		// could run more specific methods on fA here 
		RunA(fA)
	case "B":
		var fB fullB
		copier.Copy(&fB.common, e.common)
		copier.Copy(&fB.specificToB, e.specificToB)
		fmt.Printf("\n %#v \n", fB)
		RunB(fB)
	}
}


// -----------------------------------------------------------------------------
func main() {
	// essentially generic fetch
	jsonFileA, _ := os.Open("./specA.json")
	byteValue, _ := ioutil.ReadAll(jsonFileA)
	var eA everything
	json.Unmarshal(byteValue, &eA)

	jsonFileB, _ := os.Open("./specB.json")
	byteValue, _ = ioutil.ReadAll(jsonFileB)
	var eB everything
	json.Unmarshal(byteValue, &eB)

	// generic run
	RunE(&eB)
	RunE(&eA)
}
