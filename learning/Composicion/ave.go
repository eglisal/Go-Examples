package main

import "fmt"

// Clase abstracta
type Ave struct {
	locomoción string
	nombre     string
	especie    string
}

// Implementacion del método abstracto
func (ave Ave) trasladate() {
	fmt.Println("Me traslado " + ave.locomoción)
}
