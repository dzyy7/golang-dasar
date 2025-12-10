package main

import "fmt"

func main() {
	var nilai int
	fmt.Println("Masukkan Nilai: ")
	fmt.Scanln(&nilai)

	if nilai>=90 {
		fmt.Println("nilai bagus")
	}else if nilai>=75{
		fmt.Println("Lumayan")
	}else{
		fmt.Println("dongo")
	}
}