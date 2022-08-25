package main

import "fmt"

// VIDEO 06-03 ,04 PANIC Y RECOVER

// este metodo se usa para generar una cadena de llamadas
// para seguir el error hasta su punto de origen
// y no deja ejecutar las siuientes instrucciones

func main() {
	// ejemplo:
	// Crearemos una funcion donde ejecute una instruccion
	// y si detecta un error ejecute un PANIC

	division(10, 2)
	division(20, 4)
	division(100, 0) // AQUI SE GENERA EL -> PANIC <-
	// El panic se LEE DE ABAJO HACIA ARRIBA
	/*
		5
		5
		panic: Exploto

		goroutine 1 [running]:
		main.validarDivisor(...)
				C:/Users/osuna-dev/curso-go/panic/main.go:30
		main.division(0x683f60, 0xc000038000)
				C:/Users/osuna-dev/curso-go/panic/main.go:24 +0x8b
		main.main()
				C:/Users/osuna-dev/curso-go/panic/main.go:18 +0x45
		exit status 2
	*/
	division(80, 10)
}

// funcion division
func division(dividendo, divisor int) {

	// PARA CONTROLAR EL PANIC Y RECUPERARSE DE EL USAMOS EL RECOVER
	// PARA USAR EL RECOVER DEBEMOS USAR DEFER Y UNA FUNC ANONIMA
	// COMO LO SIGUIENTE
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperandome del PANIC", r)
		}
	}()

	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Exploto")
	}
}
