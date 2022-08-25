package main

import "fmt"

// VIDEO 05-06 FECHA 23/11/2021
// RECIBIR Y RETORNAR FUNCIONES

func main() {
	//SLICE
	/*nums := []int{1, 4, 70, 5, 67, 90, 2}

	result := filter(nums, func(num int) bool {
		if num >= 10 {
			return true
		}
		return false
	})

	fmt.Println(result)*/

	// FUNCION QUE RETORNE UNA FUNCION
	// sintaxi para ejecutar el retorno de doble func
	// Modo 1
	x := hello("Yordhis")("Osuna")

	fmt.Println(x)
	//modo 2
	/*
		x := hello("Yordhis")
		fmt.Println(x("Osuna"))
	*/

}

// implementar una funcion que reciva una func
// y filtre y retorne una func

// nums: es el slicesolicitado
// callback: es la funcion que va amanejar la condicion parafiltrar
// int[]: esto es lo que retornara la funcion un SLICE
/*func filter(nums []int, callback func(int) bool) []int {
	// implementamos
	// declaramos la variable SLICE que va a
	result := []int{}

	// Recorremos el slice (Array)
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}*/

// FUNCION QUE RETORNE UNA FUNCION

func hello(name string) func(string) string {
	return func(lastName string) string {
		return "hello " + name + " " + lastName
	}
}
