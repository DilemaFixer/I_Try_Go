package main

import "fmt"

type base struct {
	num int
}

type emb struct {
	base
	str string
}

func main(){
	em := emb { base : base { num : 20 } , str : "some string"} 

	fmt.Println("Emb str :" , em.str)
	fmt.Println("Base struct prip in emb :" , em.num)
}
