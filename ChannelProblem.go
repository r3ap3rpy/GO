package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(4)
	mutex := make(chan bool, 1)
	for i:=1; i<100;i++{
		for j:=1;j<100;j++{
			mutex <- true
			go func(){
			fmt.Printf("%d + %d = %d\n",i,j,i+j)
			<- mutex
			}()
		}
	}
}
