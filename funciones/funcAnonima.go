package main

import "fmt"

func main() {

	// FUNCION ANONIMA EN VARIABLE
	x := func() {
		fmt.Println("hola FUNCION ANONIMA EN VARIABLE")
	}
	// ejecutando la Func Anonima
	x()

	// FUNCION ANONIMA AUTO EJECUTADA
	func() {
		fmt.Println("hola FUNCION ANONIMA AUTO EJECUTADA")
	}()

	// FUNCION ANONIMA AUTO EJECUTADA CON PARAMETROS
	func(name string) {
		fmt.Println("hola FUNCION ANONIMA AUTO EJECUTADA", name)
	}("CON PARAMETROS")
}
