package main

type Gato struct {
	cant_vidas int
}

// Implementacion del método abstracto
func (gato Gato) juega() string {
	return "Me siento en el regazo"
}
