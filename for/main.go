package main

import "fmt"

func main() {

	// FOR clasico
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	// FOR Continuo
	/*i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}*/

	// FOR para siempre
	// este puede ser usado para soket o funciones que necesitan ser
	// escuchadas siempre
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}
