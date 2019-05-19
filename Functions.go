package main

import "fmt"

func Simplest(){
	fmt.Println("Simplest")
}


func SimpleReturn (message string)(string){
	return "Got message: " + message
}

func MultipleReturn (message string)(ok string, err string){
	ok = "Got the message: " + message 
	err = "No error this time"
	return
}

func VariableNumberArgs (message string, rest ...string)(ok int, err string){
	fmt.Println("The message was: ", message)
	fmt.Println("The rest was: ", rest)
	ok = len(rest)
	err = "Not sure if there was any error"
	return
}

func main(){
	Simplest()
	returned := SimpleReturn("This was the message!")
	fmt.Println("Returned from SimpleReturn: ",returned)
	ok, err := MultipleReturn("Not sure if it works!")
	fmt.Println(ok, err)
	okay, errror := VariableNumberArgs("Welcome to my Go course", "The","rest","stays", "the", "rest")
	fmt.Println(okay, errror)
}
