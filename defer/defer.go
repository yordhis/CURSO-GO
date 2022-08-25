package main

import (
	"fmt"
	"os"
)

// VIDEO 06-02 DEFER PARTE 2

// Cuales son los casos de uso??
// Son para cerrar archivos, controladores de base de datos,
// limpiar recurso, cerrar conecciones de red.

func main() {
	// crear archivo
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
	}
	// con defer ya no es necesario usar en otras secciones del codigo
	// el file.close porque el defer se encarga de enlistar despues de un
	// RETURN y ejecutarse
	defer file.Close()

	_, err = file.Write([]byte("Hello Soldado"))

	if err != nil {
		// Sin DEFER tendriamos que colocar
		// en cada bloque de codigo donde se use file
		// el metodo -> close <- para cerrar el archivo
		/*file.Close()*/
		fmt.Printf("Ocurrio un error al Escribir el archivo: %v", err)
	}

}
