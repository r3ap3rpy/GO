package main

import "fmt"

func main(){
	var my_string string = "Welcome to GO programming!"
	var my_string_pointer *string

	my_string_pointer = &my_string

	fmt.Println(my_string, *my_string_pointer, my_string_pointer)

	*my_string_pointer = "Hello!"

	fmt.Println(my_string, *my_string_pointer, my_string_pointer)
}
