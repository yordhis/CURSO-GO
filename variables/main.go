package main

import "fmt"

func main() {
	//declaracion version 1
	// var dog string
	// dog = "dog"

	//declaracion version 2
	//simplificamos todo en una linea
	// var dog string = "dog"

	//declaracion version 3
	//simplificamos todo en una linea y multi variables
	// var dog, gat string = "dog", "gat"

	//declaracion version 4
	//simplificamos todo en una linea y multi variable y quitamos
	//el tipo de datos
	// var dog, gat = "dog", "gat"

	//declaracion version 5
	//simplificamos todo en una linea y multi variable y quitamos
	//el tipo de datos y quitamos VAR e incluimos el chorjan :=
	//GO detecta automaticamente que es una variable
	// dog, gat := "dog", "gat"

	//declaracion version 6
	//el shorjan funciona siempre que halla una nueva variable
	//declarada de lo contrario arroja:
	//\main.go:33:6: no new variables on left side of :=
	dog, gat := "dog", "gat"
	gat, face := "gato", "carita"

	//siempre hay que usar la variable de lo contrario
	// el error es: \main.go:5:6: dog declared but not used

	//para usar e imprimir la variable usamos la libretia fmt
	fmt.Println(dog, gat, face)
}
