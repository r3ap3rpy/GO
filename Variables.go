package main

import "fmt"

func main(){
	var a int
	var b float64
	var c string 
	var d bool

	fmt.Println(a,b,c,d)

	var e, f int

	fmt.Println(e,f)

	g := 99

	fmt.Println(g)

	var h, j, k = 3.99, "string", true

	fmt.Println(h, j, k)

	var l, m, n int = 100, 200, 300

	fmt.Println(l, m, n)
}
