package main

import "fmt"

type Languages struct {
	name string
	isCompiled bool
}

func main(){
	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("The i is: ", i)
	}
	// while loop
	var j int = 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	slice := []Languages {
		{"Python",false},
		{"Go",true},
		{"C",true},
	}
	for n, c := range slice {
		fmt.Println("Here we go: ", n,c.name,c.isCompiled)
	}

}
