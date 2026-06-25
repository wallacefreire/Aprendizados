package main

import "fmt"

func main() {
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 9.0)
	slice3 = append(slice3, 9.0)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacity

	slice5 := make([]float32, 5)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
