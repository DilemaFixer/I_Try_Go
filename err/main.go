package main

import (
	"fmt"
	"errors"
	"log"
)

func divide(a, b float64) (float64 , error) {
	if b == 0 {
		return 0 , errors.New("Can't divide on zero");
	}
	return a / b , nil
}

func validAge(age int) error {
	if age < 0 {
		return fmt.Errorf("Age can't be lass than zero");
	}

	return nil
}

func main(){
	result, err := divide(10, 0)
	if err != nil {
	    log.Printf("Error: %v", err)
	    return
	}
	fmt.Println(result)
}
