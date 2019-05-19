// named collections of methods
package main

import "fmt"

type Languages []Language

type Language struct {

	name string
	version string
}

type Reversionable interface {
	Reversion(newVersion string)
}

func (language *Language) Reversion(newVersion string){
	language.version = newVersion
}

func ReversionTo(v Reversionable){
	v.Reversion("2.7")
}

func main(){
	languages := Languages{
		{"Python", "3.7"},
		{"Go","1.11"},
	}
	fmt.Println(languages)
	ReversionTo(&languages[0])
	fmt.Println(languages)
}
