package main

import "fmt"

//  VIDEO 05-07
// FUNCION VARIATICA
func main() {

	fmt.Println(sum(1, 2, 3))

}

func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}

	return result
}
