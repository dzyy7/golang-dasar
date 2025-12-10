package main

import "fmt"

func main() {
	hari := "Minggu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Budak korporat")
	case "Sabtu","Minggu":
		fmt.Println("Prei")
	default:
		fmt.Println("nggur")
	}

}