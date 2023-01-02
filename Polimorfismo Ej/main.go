package main

import (
	"fmt"
)

func main() {
	perro := Perro{true}
	gato := Gato{7}

	fmt.Println("Tipo Perro: ", perro)

	fmt.Println("Tipo Gato: ", gato)

	misAnimales := []AnimalDomestico{perro, gato}

	for _, animal := range misAnimales {
		fmt.Println(animal.juega())
	}
}
