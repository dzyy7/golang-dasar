package main
import "fmt"

func main() {
	number := [5]int{1, 2, 3,4,5}

	fmt.Println("Jumlah elemen :", len(number))
	fmt.Println("Index ke 2 :", number[2])
	number[2] = 100
	fmt.Println("Index ke 2 setelah diubah :", number[2])

	fmt.Println("ini adalah array : ")
	for index, value := range number {
	{
		fmt.Println("Index ke ", index, "=", value)
	}

	arr1:= [3]int {10,20,30}
	arr2:= [3]int {10,20,30}

	fmt.Println("Apakah arr1 dan arr2 sama ? ", arr1 == arr2)
	fmt.Println("Apakah arr1 dan arr2 berbeda ? ", arr1 != arr2)
}}