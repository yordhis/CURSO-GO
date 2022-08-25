package main

import "fmt"

// VIDEO 03-02- ARRAYS
//

func main() {
	// los arrays despues de haber declarado 3 posiciones
	// no se puede añadir mas posiciones ya que
	// los arrays en GO son fijos

	// var food [3]string
	// food[0] = "dona"
	// food[1] = "perro"
	// food[2] = "pizza"
	// fmt.Println(food)

	// Arrays literales
	food := [3]string{"dona", "perro", "pizza"}
	fmt.Println(food)
}
