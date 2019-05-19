package main

import "fmt"

func main(){
	queue := make(chan string,2)
	queue <- "Go"
	queue <- "Python"
	close(queue)

	for element := range queue {
		fmt.Println(element)
	}


}
