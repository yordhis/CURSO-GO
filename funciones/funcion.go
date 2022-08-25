package main

import (
	"errors"
	"fmt"
)

func main() {
	// esta instruccion lo que hace es leer un archivo
	// lo que se quiere es provocar un error aproposito
	// para controlarlo
	/*content, err := ioutil.ReadFile("./flith.txt")*/

	// controlamos el retorno
	/*if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		// si capturamos un error no continua
		return
	}
	fmt.Printf(string(content))*/

	result, err := division(10, 2)

	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		// si capturamos un error no continua
		return
	}
	fmt.Println(result)

}

// Creando nuestra funcion y capturamos errores

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {

		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

// tambien se puede usar los PARAMETROS NOMBRADOS
// para inicializar variables en 0 o NIL
// ejemplo (result(0) int, err(nil) error)

// no es muy recomendado usar este tipo de sintaxis
// cuando son funciones muy grandes
// ya que no son tan legibles
/*func division(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir por cero")
		return
	}
	result = dividendo / divisor
	return
}*/
