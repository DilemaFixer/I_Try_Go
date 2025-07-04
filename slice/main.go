package main

import (
	"fmt"
)

func main(){
	s := make([]string, 10);
	for i := range 10 {
		s[i] = "google,"
	}

	fmt.Println(s);
}
