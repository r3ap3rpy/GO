package main

import "fmt"

type Languages struct {
	name string
	version string
}

func (language *Languages) Reversion(newVersion string){
	language.version = newVersion
	//fmt.Println(language)
}


func main(){
	languages := Languages{name:"Python",version:"3.7"}
	fmt.Println(languages)
	languages.Reversion("2.7")
	fmt.Println(languages)
}
