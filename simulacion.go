package main

import (
	"errors"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// Espacio de simulación
	ticksPorSegundo = 50   // Instantes en que se actualiza la realidad.
	factorDeTiempo  = 20.0 // Cámara rápida mayor que 1 o lenta menor que 1.
	metrosPorPixel  = 2.0  // Cuántos metros representa cada pixel. Default 2.0
	imgFondoPath    = "imagenes/mapa-stars.png"

	// Abejas
	numeroDeAbejas            = 150
	velocidadAbeja            = 2.0 // metros por segundo
	direccionAleatoriaEnVuelo = true
	limitePolenRecolectado    = 400
	radioDeExploracion        = 1200 // metros
	viajesPorAbeja            = 3
	toleranciaPanal           = 5 // metros de tolerancia para que detecte que la abeja llegó al panal
)

var tiempoPorTick float64 = 1.0 * factorDeTiempo / ticksPorSegundo // segundos que avanza la simulación cada tick

type panal struct {
	x, y float64

	imgPanal *ebiten.Image
	imgOpt   *ebiten.DrawImageOptions
}

// Memoria de la simulación
type Simulacion struct {
	screenW int // pixeles
	screenH int // pixeles

	imgAbeja  *ebiten.Image
	imgAbeja2 *ebiten.Image
	imgFondo  *ebiten.Image

	espacio image.Image

	abejas []abeja

	panal panal

	ticks  uint64 // transcurridos
	frames uint64 // transcurridos

	// ultimaAbejaSalio       uint64 // segundos transcurridos
	// siguienteAbejaPorSalir int    // index

	pausa bool
}

// Datos generales de la simulación en texto.
func (s Simulacion) String() string {
	return fmt.Sprintf("Simulación: tick %d  frame %v  segundos %v  vel_abejas %0.1fm/s ", s.ticks, s.frames, s.SegundosTranscurridos(), velocidadAbeja)
}

// Segundos transcurridos desde el inicio de la simulación.
func (s Simulacion) SegundosTranscurridos() uint64 {
	return s.ticks * factorDeTiempo / ticksPorSegundo
}

func NuevaSimulación() (*Simulacion, error) {

	s := Simulacion{abejas: []abeja{},
		panal: panal{x: 1250, y: 800},
	}
	var err error

	// Dimensiones del espacio
	s.screenW, s.screenH, s.espacio = cargarFondoSimulacion(imgFondoPath)
	if s.screenH == 0 || s.screenW == 0 {
		return nil, errors.New("imagen de fondo tiene 0px en sus dimensiones")
	}

	// Cargar imágenes
	s.imgFondo, _, err = ebitenutil.NewImageFromFile(imgFondoPath)
	if err != nil {
		return nil, err
	}
	s.panal.imgPanal, _, err = ebitenutil.NewImageFromFile("imagenes/panal-100.png")
	if err != nil {
		return nil, err
	}
	s.panal.imgOpt = &ebiten.DrawImageOptions{}
	s.panal.imgOpt.GeoM.Translate(-50, -50) // mover al centro de la imagen
	s.panal.imgOpt.GeoM.Translate(ToPixels(s.panal.x), ToPixels(s.panal.y))

	s.imgAbeja, _, err = ebitenutil.NewImageFromFile("imagenes/bee20.png")
	if err != nil {
		return nil, err
	}
	s.imgAbeja2, _, err = ebitenutil.NewImageFromFile("imagenes/bee20a.png")
	if err != nil {
		return nil, err
	}

	// Abeja manual
	// s.abejas = append(s.abejas, abeja{
	// 	id:        0,
	// 	x:         ToMetros(780),
	// 	y:         ToMetros(35),
	// 	velocidad: velocidadAbeja,
	// 	angulo:    180,
	// })

	// Abejas aleatorias
	for i := 0; i < numeroDeAbejas; i++ {
		s.abejas = append(s.abejas, abeja{
			id: i + 1,
			// x: float64(randomBetween(0, int(ToMetros(float64(s.screenW))))),
			// y: float64(randomBetween(0, int(ToMetros(float64(s.screenH))))),
			x: s.panal.x,
			y: s.panal.y,

			// velocidad: velocidadAbeja,
			angulo: float64(randomBetween(0, 360)),

			polenRecolectado: randomBetween(i*ticksPorSegundo, i*ticksPorSegundo+ticksPorSegundo),

			enPanal: true,
		})
	}

	// Texto
	cargarFuentesParaTexto()

	// Ventana
	ebiten.SetWindowTitle("Simulación abejas")
	ebiten.SetWindowSize(s.screenW, s.screenH)
	ebiten.SetTPS(ticksPorSegundo)

	return &s, nil
}
