package main

import  "fmt"

func main(){
	// goish
	var my_goish_slice []int
	// initialize (type, initial_cap, max_cap)
	my_goish_slice = make([]int, 3, 5)
	my_goish_slice[0] = 1
	my_goish_slice[1] = 2
	my_goish_slice[2] = 3
	// THis is not how you append
	//my_goish_slice[3] = 4
	//my_goish_slice[4] = 5
	//my_goish_slice[5] = 6
	fmt.Println(my_goish_slice)
	//This is how you append
	my_goish_slice = append(my_goish_slice,4)
	fmt.Println(my_goish_slice)

	my_short_slice := []int{1,2,3,4,5,6}
	fmt.Println(my_short_slice)

	// deleting from slice
	my_short_slice = append(my_short_slice[:2],my_short_slice[3:]...)
	fmt.Println(my_short_slice)


}
