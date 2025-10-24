package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "1000b"
	number, err := strconv.Atoi(text)
	if err != nil{
		fmt.Println("Error konversi:", err.Error())
	}else{
		fmt.Println("number", number)
	}

}