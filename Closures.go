package main

import "fmt"

func decreaser() func() int {
	decreaseFrom := 10
	return func() int {
		decreaseFrom--
		return decreaseFrom
	}
}

func main(){
	dec := decreaser()
	fmt.Println(dec())
	fmt.Println(dec())

	decc := decreaser()
	fmt.Println(decc())
	fmt.Println(decc())
}
