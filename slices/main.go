package main

// VIDEO 03-03-SLICE
import "fmt"

func main() {
	// Los slice son apuntadores de arrays
	// Slice no poseen datos, son apuntadores a un array.
	set := [7]string{"leon", "caballo", "vaca", "mariposa", "loro", "avioneta", "avion"}
	//slice
	// el ultimo elemento del slice es excluyente
	animals := set[:5] // animales
	fly := set[3:]     // voladores
	// animals := set[0:5] // animales
	// fly := set[3:7]     // voladores

	fly[0] = "aguila"

	fmt.Println("array:", set)
	fmt.Println("animales:", animals)
	fmt.Println("voladores:", fly[0])
}
