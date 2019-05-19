package main

import "fmt"

type UserInteger int

type Languages struct {
	Name string
	IsCompiled bool
}

func main(){
	var a UserInteger = 99
	var language = Languages{Name: "Go", IsCompiled: true}
	fmt.Println(a)
	fmt.Println(language)
	fmt.Printf("The name of the language is: %s, and it's compiled: %s",language.Name,language.IsCompiled)

}
