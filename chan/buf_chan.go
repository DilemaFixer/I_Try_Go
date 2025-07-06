package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int , 2)
	fmt.Println("Crate chan with buf")
	ch <- 42
	ch <- 32
	//ch <- 21 buf is full , deadlock of main gorutine

	go func(){
		for {
			value := <- ch
			fmt.Println("Value from gorutine :" , value)
		}
	}()

	time.Sleep(time.Second)
	ch <- 21
	time.Sleep(time.Second)
	ch <- 87
	time.Sleep(time.Second)
	ch <- 90
}
