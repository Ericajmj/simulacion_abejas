package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *Simulacion) Layout(outsideWidth, outsideHeight int) (int, int) { // laytour establece el tamaño de la ventana simulación
	return s.screenW, s.screenH
}

// La función principal del programa
func main() {
	// Se intenta crear una nueva instancia de la simulación.
	simulacion, err := NuevaSimulación()
	if err != nil {
		log.Fatal(err) //// En caso de error al iniciar la simulación, se imprime el error y el programa se cierra.
	}
	fmt.Println("Simulación iniciada") // Mensaje indicando que la simulación se ha iniciado con éxito.

	if err := ebiten.RunGame(simulacion); err != nil { // Se ejecuta la simulación utilizando la biblioteca Ebiten
		log.Fatal(err) // imprime el error y el programa se cierra.
	}
}
