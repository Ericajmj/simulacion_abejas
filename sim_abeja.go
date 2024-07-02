package main

import (
	"fmt"
	"math"
)

// Estructura que representa una abeja //
type abeja struct {
	id int // número único de la abeja comenzando en 1.

	x float64 // metros
	y float64 // metros
	// velocidad float64 // metros por segundo
	angulo float64 // grados (tetha). Contrario al reloj. Dirección hacia la cual se está moviendo la abeja.

	recorrido float64 // distancia recorrida total (metros).

	// indican la ubicación de la abeja
	enPanal bool // si está volando, es obrera.
	enFlor  bool // Está en flor

	regresando bool // regresando al panal

	// Contadores utilizados para gestionar ciertos comportamientos de la abeja //
	polenRecolectado       int
	cooldownGiroAleatorio  int
	cooldownGiroHaciaPanal int
	olerCooldown           int

	viajesRealizados int // contador que registra la cantidad de viajes realizados por la abeja
}

// Método que simula el movimiento de la abeja en el espacio //
func (a *abeja) Mover() {
	distancia := velocidadAbeja * tiempoPorTick
	a.recorrido += distancia
	// Convertir ángulo en radianes: angulo*math.Pi/180
	a.x += distancia * math.Cos(a.angulo*math.Pi/180)
	a.y += distancia * math.Sin(a.angulo*math.Pi/180)
}

// Rebotar como laser:
// a.angulo = math.Atan((a.y-y)/(a.x-x)) * 180 / math.Pi

// Girar abeja hacia una coordenada en metros.
func (a *abeja) DirigirHacia(x, y float64) {
	// oldAngulo := a.angulo
	a.angulo = math.Atan2(a.y-y, a.x-x) * 180 / math.Pi
	a.angulo += 180 // corregir el que y sea positiva hacia abajo.
	// fmt.Printf("Abeja[%d] irá a (%0.1f,%0.1f) desde (%0.1f,%0.1f) con %0.1fº (antes %0.1fº)\n", a.id, x, y, a.x, a.y, a.angulo, oldAngulo)
}

// Posición, ángulo o estados de la abeja //
func (a abeja) String() string {
	estatus := ""
	if a.enFlor {
		estatus += " enFlor"
	}
	if a.enPanal {
		estatus += " enPanal"
	}
	if a.regresando {
		estatus += " regresando"
	}
	return fmt.Sprintf("Abeja: %0.1fº recorrido %0.1fm (%0.1fm, %0.1fm) polen %d %v", a.angulo, a.recorrido, a.x, a.y, a.polenRecolectado, estatus)
}

// Flores en los jardines urbanos de interés de las abejas acorde al Síndrome de polinización
func EsFlorDeInteres(hexColor string) bool {
	switch hexColor {
	case "009DE0": //Flores azules
		return true
	case "F759E1": //Flores rosas
		return true
	case "FFCC33": //Flores amarillas
		return true
	case "75447A": //Flores moradas
		return true
	default:
		return false
	}
}

// Método olfato para que la abeja identifique si hay flor cercana a ella
func (s *Simulacion) OlerFloresCercanas(a *abeja) {
	if a.olerCooldown != 0 {
		a.olerCooldown--
		return
	}
	a.olerCooldown = 200

	x := ToPixelsInt(a.x)
	y := ToPixelsInt(a.y)

	porVisitar := [][2]int{
		// primero revisar direcciones cardinales //Coordenas 2 pixeles a la redonda donde visita la abeja la flor
		{-1, +0},
		{+0, -1},
		{+0, +1},
		{+1, +0},
		{-2, +0},
		{+0, -2},
		{+0, +2},
		{+2, +0},

		{-1, -1},
		{-1, +1},
		{+1, -1},
		{+1, +1},

		{-2, -2},
		{-2, -1},
		{-2, +1},
		{-2, +2},

		{-1, -2},
		{-1, +2},
		{+1, +2},
		{+1, -2},

		{+2, -2},
		{+2, -1},
		{+2, +1},
		{+2, +2},
	}

	for _, xy := range porVisitar {
		if EsFlorDeInteres(rgbaToHex(s.espacio.At(x+xy[0], y+xy[1]))) {
			a.DirigirHacia(ToMetros(float64(x+xy[0])), ToMetros(float64(y+xy[1])))
			return
		}
	}
}
