package main 

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	fmt.Println("Create chan int");
	//ch <- 42 this code block main gorutine , will trow error "deadlock"
	/*
	go func() {
		ch <- 42
	}()
	// this gorutin will be blocked forever , deadlock . 
	Chan have no buf and recvrs
	*/

	go func() {
		fmt.Println("in gorutine")
		value := <-ch
		fmt.Println("chan give" , value)
	}()

	time.Sleep(2 * time.Second)
	ch <- 42
	time.Sleep(time.Second)
}
