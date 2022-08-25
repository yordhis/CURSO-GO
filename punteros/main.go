package main

// VIDEO 03-01- PUNTEROS
// LOS PUNTEROS SON LAS DIRECCIONES EN MEMERIA DE UNA VARIABLE
import "fmt"

func main() {
	fruit := "apple"
	// para acceder a la direccion en memoria utilizamos:
	// %v y el anspersan antes de la variable -> &fruit
	// ahora si queremos hacer un puntero declaramos una variable p con un asterisco
	// antes del tipo de dato
	//puntero
	var p *string
	p = &fruit
	// para editar el puntero utilizamos el mismo puntero *
	// y asignamos un valor y al asignar un nuevo valor la variable original cambia
	*p = "pino"
	fmt.Printf("Tipo: %T, Valor: %s, direccion: %v\n", fruit, fruit, &fruit)

	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s\n", p, p, *p)

}
