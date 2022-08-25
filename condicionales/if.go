package main

import "fmt"

func main() {
	/*emoji := "cambur"
	if emoji == "cactup" {
		fmt.Println("es un cactus")
	} else if emoji == "manzana" {
		fmt.Println("es un manzana")
	} else {
		fmt.Printf("el emoji es: %s\n", emoji)
	}*/

	// tambien GO tiene una particularidad
	// de manejar variables internas dentro de las condiciones
	// y no se pueden usar afuera del bloque condicional

	if emoji := "carita"; emoji == "cactus" {
		fmt.Println("es un cactus")
	} else if emoji == "carita" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s\n", emoji)
	}
}
