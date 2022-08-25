package main

import "fmt"

// VIDEO 03-05
// LOS MAPS

func main() {
	// map[tipo de dato del map]tipo de dao que va a almacenar
	animals := make(map[string]string)
	animals["cat"] = "cat"
	animals["dog"] = "dog"
	fmt.Println(animals)

	// segunda manera de declarar los MAP
	// esta manera utiliza las llaves despues de la declaracion
	// y dentro de las llaves ingresamos los valores
	// como un array asociativo
	fruits := map[string]string{
		"apple":  "apple",
		"banana": "banana", // las comas son obligatorias
	}
	fmt.Println(fruits)
	// para eliminar un elemento del map
	// se usa el metodo preconstruido delete()
	/*delete(fruits, "banana")*/

	// validamossi la llave existe con la siguiente
	// expresion
	/*emoji, ok := animals["gorila"]
	fmt.Println(emoji, ok)*/

	// utilizamos una estructura de control
	if _, ok := animals["gorila"]; !ok {
		animals["gorila"] = "gorila"
	}
	fmt.Println(animals)
}
