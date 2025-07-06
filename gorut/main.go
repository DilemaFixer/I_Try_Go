package main

import (
	"fmt" 
	"time"
)

func f(size int ,text string){
	for i := range size {
		fmt.Println(text,i)
		time.Sleep(time.Second)
	}
}

func main(){
	go f(5 , "Range 1 :")
	go f(7 , "Range 2 :")
	time.Sleep(time.Minute)
}
