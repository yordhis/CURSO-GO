package main

import "fmt"

func main() {
	//bool, string, numeric
	//DATO BOOL
	// var a bool = true

	//DATO STRING
	//no es valido usar comillas simples
	// var a string = "hola"

	//DATO NUMERIC
	var a uint8 = 255
	var b uint16 = 255
	// no se pueden realizar operaciones con tipos de datos diferentes
	// para hacerlo posible debemos realizar un casting convirtiendo
	// a la variable a en uint16 para poder sumarlo
	// y como las dos variables son del mismo tipo se le asigna
	// ala variable C el tipo de dato UINT16
	c := uint16(a) + b

	// Si deseamos declarar una variable pero que aun
	// no vamos a usar utilizamos el guion bajo -> _ <-
	// identificador blank
	var _ string = "test"
	_ = 22

	// Por ultimo si declaramos una variable y no le asignamos un dato
	// GO se encarga de asignarle CERO
	// var z bool -> false
	// var z string -> vacio ""
	// var z uint16 -> 0

	fmt.Printf("Tipo: %T, Valor: %v", c, c)
}

// VALORES DEL INT & UINT

//UINT
//uint8 unsigned 8-bit integers (0 to 255)
//uint8 unsigned 16-bit integers (0 to 65535)
//uint8 unsigned 32-bit integers (0 to 4294967295)
//uint8 unsigned 64-bit integers (0 to 18446744073709551615)

//INT
//int8 signed 8-bit integers (-128 to 127)
//int8 signed 16-bit integers (-32768 to 32767)
//int8 signed 32-bit integers (-2147483648 to 2147483647)
//int8 signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

//FLOAT
//float32
//float64

//ALIAS
// BYTE -> uint8
// RUNE -> int32
// Representa un punto de codigo UNICODE
// es decir que se puede almacenar caracteres usando
// las comillas simples y solo puede tener un solo caracter
