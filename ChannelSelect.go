package main

import "fmt"

type Compiled struct {
	name string
	version string
}

type Interpreted struct {
	name string
	version string
}

func main(){
	cChannel := make(chan Compiled, 1)
	iChannel := make(chan Interpreted, 1)

	python := Interpreted{name:"Python3",version:"3.7"}
	goo :=Compiled{name:"Go", version: "1.11"}

	fmt.Println(python, goo)
	//iChannel <- python
	cChannel <- goo
	select {
		case language := <- iChannel:
			fmt.Println("We have received an interpreted language: ",language)
		case language := <- cChannel:
			fmt.Println("We have received a compiled language: ",language)
		default:
			fmt.Println("Nothing to do here!")
	}

}
