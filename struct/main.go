package main

import "fmt"

func main(){
	p := newPerson("Mike")
	fmt.Println(*p)
}

type person struct {
	name string 
	age int
}

func newPerson(name string) *person {
	p := person { name : name } 
	p.age = 43
	return &p;
}
