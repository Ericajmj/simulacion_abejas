package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ================================================================ //
// ========== DIBUJAR FRAME ======================================= //

func (s *Simulacion) Draw(screen *ebiten.Image) {
	if !s.pausa {
		s.frames++
	}

	// Fondo
	screen.DrawImage(s.imgFondo, &ebiten.DrawImageOptions{})

	// Panal
	screen.DrawImage(s.panal.imgPanal, s.panal.imgOpt)

	// Abejas
	for _, abeja := range s.abejas {
		if abeja.enPanal {
			continue
		}

		img := s.imgAbeja
		op := &ebiten.DrawImageOptions{}
		escala := 1.0

		if abeja.enFlor {
			escala = 1.5
			op.GeoM.Scale(escala, escala)
			img = s.imgAbeja2
		}

		op.GeoM.Translate(-10*escala, -10*escala)    // Preparar para rotar. La imagen de la abeja mide 20x20px, entonces se pone enmedio con -10,-10.
		op.GeoM.Rotate(abeja.angulo * math.Pi / 180) // Usar radianes para rotar.
		op.GeoM.Translate(ToPixels(abeja.x), ToPixels(abeja.y))
		screen.DrawImage(img, op)

	}

	// Texto
	vector.DrawFilledRect(screen, 10, 10, 500, 60, color.White, false)
	text.Draw(screen, s.String(), fontNormal, 20, 30, color.Black)
	text.Draw(screen, s.abejas[0].String(), fontNormal, 20, 50, color.Black)

	// Escala
	vector.DrawFilledRect(screen, 10, float32(s.screenH-20), ToPixels32(100), 10, color.Black, false)
	text.Draw(screen, "100m", fontNormal, 12, s.screenH-25, color.Black)

	// Pausa
	if s.pausa {
		vector.DrawFilledRect(screen, 10, float32(s.screenH)-50, 110, 40, color.RGBA{A: 150}, false)
		text.Draw(screen, "Pausa [ESC]", fontNormal, 15, s.screenH-25, color.White)
	}
}

// ================================================================ //

// func posicion(x, y float64) *ebiten.DrawImageOptions {
// 	op := &ebiten.DrawImageOptions{}
// 	// op.GeoM.Translate(MetrosToPx64(x), MetrosToPx64(y))
// 	op.GeoM.Translate(x, y)
// 	return op
// }
