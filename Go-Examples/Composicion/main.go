package main

import (
	"fmt"
)

func main() {
	gorrion := Gorrion{
		Ave{"volando", "Tari", "Passer domesticus"}, true}
	pingu := Pinguino{Ave{"nadando", "Pingu", "Aptenodytes patagonicus"}, true}

	fmt.Println("Tipo Gorrion: ", gorrion)
	gorrion.trasladate()

	fmt.Println("Tipo Piguino: ", pingu)
	pingu.trasladate()

	misAves := []Ave{pingu.Ave, gorrion.Ave}

	fmt.Println(misAves)
}
