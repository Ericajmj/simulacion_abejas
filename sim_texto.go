package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Para poder dibujar texto en la simulación.

var (
	fontNormal font.Face
	// fontBig    font.Face
)

func cargarFuentesParaTexto() {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const fontDPI = 72
	fontNormal, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     fontDPI,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	// fontBig, err = opentype.NewFace(tt, &opentype.FaceOptions{
	// 	Size:    48,
	// 	DPI:     fontDPI,
	// 	Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
