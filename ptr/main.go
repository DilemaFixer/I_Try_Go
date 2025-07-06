package main

import (
	"fmt"
)

type foo struct {
	str string 
	num int
}

func newFoo() *foo{
	//complire auto. move allocation this struct to heap
	return &foo {str:"some txt", num:100}
	
	//or do it by hand
	//foo := new(foo)
	//foo.str = "some txt"
	//foo.num = 100
	//return foo
}

func main(){
	foo := newFoo()
	fmt.Println(*foo)
}
