package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected 
	StateError 
	StateRetrying
)

var stateName = map[ServerState] string {
	StateIdle : "idle",
	StateConnected : "connected",
	StateError : "error",
	StateRetrying : "retrying",
}

func (ss ServerState) ToString() string {
	return stateName[ss]
}

func main(){
	fmt.Println(StateConnected.ToString())
}

