package main

import "fmt"

// VIDEO 02-05-OPERADORES ARITMETICO Y DE ASIGNACION
func main() {
	// Operadores ariméticos -> (), *, /, %, +, -
	var a = 4 + 2*9/2
	// Operadores de Asignación -> =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	// Declaración post-incremento y post-decremento: ++, --
	// (No son una expresion sino una declaración)
	var c = 3
	// c = c++ + 2 esto es un error
	// c++ correcto
	c-- //correcto

	// VIDEO 02-06-OPERADORES DE COMPARACION

	// OPERADORES DE COMPARACION: >, <, ==, !=, >=, <=
	// el resultado de las compraciones me retornan un valor booleano (TRUE O FALSE; 0, 1)
	// fmt.Println(4 > 5)
	// fmt.Println(4 < 5)
	// fmt.Println(4 >= 5)
	// fmt.Println(4 <= 5)
	// fmt.Println(4 != 5)
	// fmt.Println(4 == 5)
	// Operadores lógicos -> &&, ||
	// Unario -> !
	fmt.Println(!((4 == 5) || (9 < 1)))
	fmt.Println(!((4 == 5) && (9 < 1)))

	//

	//mostrar el reultado
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(c--) es invalido porque este metodo espera es una expresion
	// fmt.Println(c++) y no una declaracion -> -- o ++

}
