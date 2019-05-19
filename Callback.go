package main

import "fmt"

type Language struct {
	name string
	version string
}

func Reversion(language *Language, callback chan *Language){
	language.version = "1.11"
	callback <- language
}

func main(){
	lang := new(Language)
	lang.name = "go"
	lang.version = "1.10"
	fmt.Println(lang)
	ch := make(chan *Language)
	go Reversion(lang, ch)
	newVersion := <- ch
	fmt.Println(newVersion)
}
