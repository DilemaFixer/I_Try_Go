package main

import "fmt"

func main(){
	a := [5]int{100, 1000, 1000, 13412, 5432}
	fmt.Println("a : " , a);

	a[0] = 34
	a[1] = 85
	a[2] = 395
	fmt.Println("a : " , a);
	
	fmt.Println("a len : " , len(a));
}
