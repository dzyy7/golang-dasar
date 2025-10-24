package main

import "fmt"

func main() {
	arr := [6] int {10,20,30,40,50,60}

	slice1 := arr[:]
	fmt.Println("Ini adalah slice 1 : ",slice1)

	slice2:= arr[:2]
	fmt.Println("Ini adalah slice 2 : ",slice2)

	slice3:= arr[3:]
	fmt.Println("Ini adalah slice 3 : ",slice3)

	slice4:= arr[1:2]
	fmt.Println("Ini adalah slice 4 : ",slice4)
}