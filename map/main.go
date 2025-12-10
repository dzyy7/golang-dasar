package main

import "fmt"

func main() {
	hewan := map[string]string{
		"nama":  "Gajah",
		"rumah": "hutan",
	}
	fmt.Println("namaku",hewan["nama"])
	fmt.Println("rumahnya di", hewan["rumah"])
}