package main

import "fmt"

// VIDEO06-01 DEFER parte 1

// DEFER: es diferir o aplazar
func main() {
	// defer lo que hace es aplazar
	// la instrccion y la en lista
	// para ejecutarla luego
	// ejemplo:
	/*
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
	*/
	// la primera que difiere es la ultima
	// en ejecutarse
	/*defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)*/

	// otro ejemplo es:
	a := 5
	// se espera que imprima la instruccion diferida el 10
	// pero no lo hace ya que cuando DEFER aplaza la instruccion
	// guarda los argumento ingresados en los parametros en ese
	// momento; por lo tanto imprime el 5 que se asigno en ese momento
	defer fmt.Println("Defer:", a)

	a = 10
	fmt.Println(a)

}
