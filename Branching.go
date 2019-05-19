package main

import "fmt"

func IfDemo(x int, y int)(result string){
	if x > y {
		result = "X is greater than Y"
	} else if x < y {
		result = "Y is greater than X"
	} else {
		result = "X is equal to Y"
	}
	return
}

func SwitchDemo(name string)(greeting string){
	switch name {
		case "Daniel":
			greeting = "Hello buddy, " + name
			fallthrough
		case "Eve":
			greeting = "Hello Mrs. " + name
		case "Janet":
			greeting = "Hello Ms. " + name
		case "Eveling", "Ildiko", "Naomi":
			greeting = "Hello Colleagues, " + name
		default:
			greeting = "Welcome stranger, " + name
	}
	return
}

func main(){
	fmt.Println(IfDemo(5, 5))
	fmt.Println(SwitchDemo("Daniel"))
	fmt.Println(SwitchDemo("Janet"))
	fmt.Println("Unknown")
}
