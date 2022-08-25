package main

import (
	"fmt"
)

func main() {
	// *** FOR RANGE SLICE ***
	nums := []uint8{2, 4, 6, 8}
	// ejemplo 1
	// se imprimira Edteam 4 veces por los elementos del slice
	// serian 4 iteraciones
	for range nums {
		fmt.Println("EDteam")
	}

	// Ejemplo 2
	// como imprimi los valores
	// el for nos provee de dos valores
	// i: que posee los indices
	// v: que posee los valores (esta variable solo es una copia del
	// y no de le puede asignar un valor )
	for i, v := range nums {
		fmt.Println("Indice:", i, "Valor", v)
	}
	// este codigo no asignarianada al slice nums
	/*for _, v := range nums {
		V *= 2
	}
	fmt.Println(nums)*/
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)
	//*************************************

	// *** FOR RANGE MAP ***

	sports := map[string]string{"basketball": "pelota1", "soccer": "pelota2"}
	// En este caso en ves de i usamos:
	// KEY que representa las llaves basketball y soccer
	// V sigue siendo los valores de las llaves pelota1 y pelota2
	for key, v := range sports {
		fmt.Println("Key", key, "Valor", v)
	}
	//************************************

	// *** FOR RANGE STRING ***

	hello := "hello"
	// Para aplicar un for a un string y obtener letra por letra
	// hay que aplicar el casting en la variable V -> string(v)
	// de lo contrario nos devuelve los valores en BYTE de cada letra
	for _, v := range hello {
		fmt.Println(string(v))
	}

}
