package main

import "fmt"

const (
	Language = "English"
	Type = "Spoken"
	Popularity = "High"
)

//iota

/**/

/*
const (
	A = 1
	B = 2
	C = 3
	D = iota
	E = iota
	F
	G = iota
)*/

const (
	A = iota
	B
	C
	D
	E
	F
	G
)

func main(){
	// var a, b, c = 1, 2 ,3
	//Language = "Hungarian"
	fmt.Println(Language, Type, Popularity)
	fmt.Printf("A: %d, B: %d, C: %d, D: %d, E: %d, F: %d, G: %d", A,B,C,D,E,F,G)
}
