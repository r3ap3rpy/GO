package main

import (
	"time"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(2)
	duration, _ := time.ParseDuration("10ms")
	go func() {
		for i:=0;i<100;i++{
			println("Hello from thread 1")
			time.Sleep(duration)
		}
	}()

	go func(){
		for j:=1;j<100;j++{
			println("Hello from thread 2")
			time.Sleep(duration)
		}
	}()
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
