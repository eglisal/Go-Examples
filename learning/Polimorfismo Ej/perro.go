package main

type Perro struct {
	cazador bool
}

// Implementacion del método abstracto
func (perro Perro) juega() string {
	return "Busco el palito"
}
