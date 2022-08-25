package main

import "fmt"

// VIDEO 03-06

// ESTRUCTURAS

// las estructuras en GO son similares a CLASES  en
// otros lenguajes de prograamacion

func main() {
	// declaracion de las variable de la estructura
	type course struct {
		// declaramos las variables
		Name     string
		Profesor string
		Country  string
	}

	// asignacion de valores a la estructura modo 1
	db := course{
		Name:     "base de datos",
		Profesor: "Alexis",
		Country:  "Colombia",
	}

	// asignacion de valores a la estructura modo 2
	/*git := course{"git", "Beto", "Bolivia"}*/

	// Para asignar a un solo campo debemos usar el modo 1
	/*css := course{Profesor: "Alvaro"}*/

	// podemos tambien crear punteros
	p := &db // optenemos la direccion
	// para cambiar los valores lo podemos
	// hacer desde el metodo de desreferenciacion (*Variable).NombreDelCampo
	// tambien se puede escribir mas sencillo: Variable.NombreDelCampo
	// porque en el casos de las estructras GO se encarga de hacer
	// desreferenciacion.

	/*(*p).Profesor = "Yordhis"*/
	p.Profesor = "Yordhis"

	// Tambien se puede acceder a los valores
	// usando el puntoNombreDeCampo Ejemplo: db.Name
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
	/*fmt.Printf("%+v\n", git.Country)
	fmt.Printf("%+v\n", css.Profesor)*/
}
