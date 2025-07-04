package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	fmt.Print("Input some text =) : ");
	reader := bufio.NewReader(os.Stdin);
	text , _ := reader.ReadString('\n');
	fmt.Print(text);
}
