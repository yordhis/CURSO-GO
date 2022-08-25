package main

import "fmt"

func main() {

	emoji := "banana"

	// Modo 1
	/*switch emoji {
	case "cat":
		fmt.Println("es un gato")

	case "dog":
		fmt.Println("es un perro")

	case "banana":
		fmt.Println("es un banana")

	case "manzana":
		fmt.Println("es un manzana")

	default:
		fmt.Println("no es un animal ni una fruta")
	}*/

	// Modo 2
	// optimizamos los casos similares
	/*switch emoji {
	case "cat", "dog":
		fmt.Println("es un animal")

	case "banana", "manzana":
		fmt.Println("es un fruta")

	default:
		fmt.Println("no es un animal ni una fruta")
	}*/

	// Modo 3
	// Tambien se puede usar los operadores de comparacion y logicos
	// dentro de un switch
	// en este caso se debe quitar la variable emoji
	switch {
	case emoji == "cat" || emoji == "dog":
		fmt.Println("es un animal")

	case emoji == "banana" || emoji == "manzana":
		fmt.Println("es un fruta")

	default:
		fmt.Println("no es un animal ni una fruta")
	}
}
