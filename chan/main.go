package main 

import (
	"fmt"
	"runtime"
)

func trySend(ch chan<- int , value int) bool {
	select {
	case ch <- value :
		return true
	default:
		return false
	}
}

func tryReceive(ch <-chan int) (int , bool){
	select{
	case value := <- ch :
		return value , true
	default:
		return -1 , false
	}
}

func main(){
	ch := make(chan int , 2)
	ch <- 32 
	// if uncomment line below ,buffer will haven't place for take one number
	//ch <- 74 
	
	isSuccessfully := trySend(ch , 55)

	fmt.Println("Is send successfully :", isSuccessfully)

	var foo int
	if runtime.selectnbrecv(&foo , ch){
		fmt.Println("All is ok , result :" , foo)
	}
}
