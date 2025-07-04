package main

import "fmt"

func main(){
	fmt.Println("Go loops test \n");
	fmt.Println("i := 5 \nfor i < 5");

	i := 0
	for i < 5 {
		fmt.Println("i : " , i);
		i++
	}
	fmt.Println("");
	fmt.Println("for j := 0; j < 5; j++");
	for j := 0; j < 5;j++{
		fmt.Println("j : " , j);
	}

	fmt.Println("");
	fmt.Println("for i := range 5");
	for i := range 5 {
		fmt.Println("i : " , i);
	}
	fmt.Println("");
	fmt.Println("for { \n// some code \n break \n}\n");
	for {
		fmt.Println("loop with break");
		break
	}
}
