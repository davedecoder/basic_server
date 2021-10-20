package main

import "fmt"

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

type ByteSize float64

const (
	_ = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

const (
	read = 1 << iota
	write
	remove

	admin = read | write | remove
)

func main(){
	fmt.Printf("The value of KB    is %v\n", KB)
	fmt.Printf("The value of MB    is %v\n", MB)
	fmt.Printf("The value of GB    is %v\n", GB)
	fmt.Printf("The value of TB    is %v\n", TB)
	fmt.Printf("The value of PB    is %v\n", PB)
	fmt.Printf("read = %v\n", read)
	fmt.Printf("write = %v\n", write)
	fmt.Printf("remove = %v\n", remove)
	fmt.Printf("admin = %v\n", admin)
}
