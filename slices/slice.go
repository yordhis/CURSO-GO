package main

import "fmt"

// SLICE PARTE 2
// VIDEO 03-04

func main() {
	// Len(): # de elementos en el slice
	// Cap(): # de elementos del array donde apunta el slice,
	// a partir del indice de donde se creo el slice

	/*food := [5]string{"pan", "fresa", "limon", "pizza", "pepito"}
	fruits := food[1:3]*/

	// para añadir elementos de manera dinamica a un array
	// por medio de un slice

	/*fruits = append(fruits, "pina", "melon", "mango", "banana")*/

	// al añadir estos dos elementos en este slice
	// automaticamente me modifica el array original ya que
	// un slice no puede guardar datos
	// y el APPEND lo que hace es añadir un elemento conforme a
	// la capacidad desde el indice inicial del slice

	// ¿Que pasa si añadimos un nuevo elemento sobrepasando la
	// capasidad del array?

	// lo que pasa es GO al ver que sobre paso la capacidad
	// pierde la referrencia de array original
	// Array[4]{"pan", "fresa", "limon", "pizza", "pepito"}
	// y crea una nueva referencia apartir del SLICE de fruits
	// y duplica la capacidad. es decir un nuevo array
	// new Array[8]{"fresa", "limon", "pina", "melon", "mango", "banana"}

	/*fmt.Println("food: ", food)
	fmt.Println("fruits: ", fruits)
	fmt.Println("Tamaño: ", len(fruits))
	fmt.Println("Capacidad: ", cap(fruits))*/

	// Slice literales
	// son aquellos que no apuntan a un array creado por nosotros
	// sino a uno generado por GO
	// Ejemplo:
	// Modo 1 de crostruir un slice literal
	/*fruits := []string{"fresa", "limon"}*/
	// Modo 2
	/*fruits := make([]string, 0, 3) // CERO es la cantidad de elemento y TRES es la capacidad
	fruits = append(fruits, "pina", "melon", "mango", "banana")
	fmt.Println("fruits: ", fruits)
	fmt.Println("Tamaño: ", len(fruits))
	fmt.Println("Capacidad: ", cap(fruits))*/

	// Valor de Slice en CERO
	/* var fruits []string */ // no inicializado
	fruits := []string{}      // inicializado pero vacio
	fmt.Println("fruits: ", fruits)
	fmt.Println("Tamaño: ", len(fruits))
	fmt.Println("Capacidad: ", cap(fruits))
	fmt.Println(fruits == nil)

}
