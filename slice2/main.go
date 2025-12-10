package main

import "fmt"

func main() {
	s:= make([]int, 3, 6)
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	s=append(s, 10,50)
	fmt.Println("Setelah di append : ",s)
}