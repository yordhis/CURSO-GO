package main

import (
	"fmt"
	"strings"
)

// *** FUNCIONES ***

func main() {
	// llamamos una funcion con
	// solo colocar su nombre y parenteci
	// ejemplo:
	hello()
	helloCustomer("Yordhis", "Osuna")

	// variable
	// este caso no funciona
	// porque al ingresar la variable como paramentro
	// en change() simplemente se esta ingresando una copia de la
	// variable original y por eso no se puede hacer el cambio de
	// valor en la variable
	/*emoji := "pelito"
	change(emoji)
	fmt.Println(emoji)*/
	emoji := "pelito"
	// en ves de ingresar la variable normal
	// ingresamos la direccion en memoria de a variable con &
	change(&emoji)
	fmt.Println(emoji)

	//*** retorno normal
	total := suma(2, 3)
	fmt.Println(total)

	//*** retorno multiple
	texto := "YorDhiS"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)

}

// *** funcion simble ***
func hello() {
	fmt.Println("hello EDteam")
}

//***************************

// *** funcion con parametros ***
func helloCustomer(firstName string, lastName string) {
	fmt.Printf("Hello %s %s \n", firstName, lastName)
}

//*********************************

// *** funcion por referencia ***
// usamos el puntero para hacer referencia
// a la direccion de la variable *string
func change(value *string) {

	// desreferenciamos la variable VALOR
	// para que apunte a la direccion enmemoria
	*value = "carita feliz"
}

//*******************************

// *** retornar funcion ***
// Go permite indicar en elultimo parametro el tipo de valor
// siempre que sean todos int, string etc.
func suma(num1, num2 int) int {
	return num1 + num2
}

//*************************

// *** Retorno de multiples valores en las funciones ***
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}
