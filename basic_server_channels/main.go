package main

import (
	"fmt"
)

type CommandType int

const (
	GetCommand = iota
	SetCommand
	IncCommand
)

type Command struct {
	ty			CommandType
	name 		string
	val 		int
	replyChan 	chan int
}



func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
		close(ch1)
		println("Hello World")

	}()
	v := 2
	exit := false
	for {
		select {
		case v2, ok := <-ch1:
			if ok {
				fmt.Println(v, v2)
				println("read from ch1", v2, "value.")
			} else {
				exit = true
			}
		case v3, ok := <-ch2:
			if ok {
				fmt.Println(v, v3)
				println("read from ch1", v3, "value.")
			} else {
				exit = true
			}
		case ch2 <- v:
			println("wrote", v, "into ch2")
		//default:
		//	println("Default")
		}
		if exit {
			break
		}
	}
	fmt.Println("Hello World 2")
}