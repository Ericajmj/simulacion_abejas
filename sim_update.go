package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Ejecutado cada tick de la simulación.
func (s *Simulacion) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.pausa = !s.pausa
	}

	// Control manual de abeja
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.abejas[0].x -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.abejas[0].x += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.abejas[0].y -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.abejas[0].y += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		s.abejas[0].angulo -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		s.abejas[0].angulo += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		s.abejas[0].angulo -= 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		s.abejas[0].angulo += 10
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		s.abejas[0].DirigirHacia(ToMetros(800), ToMetros(30))
	}

	if s.pausa {
		return nil
	}

	s.ticks++

	// =============================== //
	// === Lógica a partir de aquí === //

	// if s.SegundosTranscurridos()-s.ultimaAbejaSalio == 1 {
	// 	s.abejas[s.siguienteAbejaPorSalir].enPanal = false
	// 	if s.siguienteAbejaPorSalir < numeroDeAbejas {
	// 		s.siguienteAbejaPorSalir++
	// 		s.ultimaAbejaSalio = s.SegundosTranscurridos()
	// 	}
	// }

	for i := range s.abejas {
		abeja := &s.abejas[i]

		// Descargar todo el pólen en el panal antes de volver a salir.
		if abeja.enPanal {
			if abeja.viajesRealizados == viajesPorAbeja {
				continue
			}
			if abeja.polenRecolectado == 0 {
				abeja.enPanal = false
				abeja.regresando = false
				abeja.recorrido = 0
				fmt.Printf("Abeja[%d] salió del panal por %vª vez hacia %0.1fº\n", abeja.id, abeja.viajesRealizados+1, abeja.angulo)
			} else {
				abeja.polenRecolectado--
			}
			continue
		}

		// Cargar pólen y no moverse si se está en una flor.
		if abeja.enFlor {
			if abeja.polenRecolectado == limitePolenRecolectado {
				abeja.DirigirHacia(s.panal.x, s.panal.y)
				abeja.enFlor = false
				fmt.Printf("Abeja[%d] recolectó polen y regresará\n", abeja.id)
			} else {
				abeja.polenRecolectado++
			}
			continue
		}

		// Avanzar abeja volando
		abeja.Mover()

		// Giro aleatorio durante el vuelo
		if direccionAleatoriaEnVuelo {
			if !abeja.regresando {
				abeja.cooldownGiroAleatorio++
				if abeja.cooldownGiroAleatorio == 30 { //cada 30 ticks la abeja cambia de posicion aleatoriamente
					abeja.angulo += float64(randomBetween(0, 30) - 15) //el angulo incremente o decrementar entre 0 y 15grados
					abeja.cooldownGiroAleatorio = 0
				}
			}
		}

		// Limitar vuelo si no encuentra una flor. Deberá regresar al panal
		if abeja.recorrido > radioDeExploracion {
			if !abeja.regresando {
				fmt.Printf("Abeja[%v] voló demasiado y regresará\n", abeja.id)
				abeja.DirigirHacia(s.panal.x, s.panal.y)
				abeja.regresando = true
			}
		}

		if abeja.regresando { //función asegura la orientación de la abeja hacia el panal
			abeja.cooldownGiroHaciaPanal++
			if abeja.cooldownGiroHaciaPanal == 150 {
				// fmt.Printf("Abeja[%v] busca el panal con polen %v\n", abeja.id, abeja.polenRecolectado)
				abeja.DirigirHacia(s.panal.x, s.panal.y)
				abeja.cooldownGiroHaciaPanal = 0
			}
		}

		// if ToPixelsInt(abeja.x) == ToPixelsInt(s.panal.x) && ToPixelsInt(abeja.y) == ToPixelsInt(s.panal.y) {
		// 	fmt.Printf("Abeja[%v] sobre panal con %v pólen y %0.1fº\n", abeja.id, abeja.polenRecolectado, abeja.angulo)
		// }

		// Detectar si llegó al panal y ya recorrió distancia mínima
		if s.panal.x-toleranciaPanal < abeja.x && abeja.x < s.panal.x+toleranciaPanal &&
			s.panal.y-toleranciaPanal < abeja.y && abeja.y < s.panal.y+toleranciaPanal &&
			abeja.recorrido > radioDeExploracion {

			abeja.x = s.panal.x
			abeja.y = s.panal.y
			abeja.enPanal = true
			abeja.regresando = false
			abeja.viajesRealizados++
			fmt.Printf("Abeja[%v] regresó al panal con %v pólen y %0.1fº por %va vez\n", abeja.id, abeja.polenRecolectado, abeja.angulo, abeja.viajesRealizados)
			if abeja.polenRecolectado == 0 {
				abeja.polenRecolectado = randomBetween(100, limitePolenRecolectado*2)
				abeja.angulo = float64(randomBetween(0, 360)) // salir en otra dirección porque no encontró nada
			} else {
				abeja.angulo -= 180 // regresar a misma flor
			}
		}

		// Detectar abeja necesita pólen.  //Límite del polen recolectado
		if abeja.polenRecolectado < limitePolenRecolectado {
			if EsFlorDeInteres(rgbaToHex(s.espacio.At(ToPixelsInt(abeja.x), ToPixelsInt(abeja.y)))) {
				abeja.enFlor = true
				// abeja.velocidad = 0
			} else {
				abeja.enFlor = false
				s.OlerFloresCercanas(abeja)
			}
		}
	}
	return nil
}
