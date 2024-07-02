package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Usar imagen PNG con color depth de 8 bits.
//
// Si el color es por ejemplo rgb(150,0 ,0 ) = #960000.
// Corresponde a un byte por canal: r = 0x96, g = 0x00 , b = 0x00
// El blanco es 0x00 y el negro 0xFF
//
// Al llamar img.At(x, y).RGBA() se retorna un RGBA de 16 bits dentro de un uint32.
// el valor máximo de cada canal sí son 16 bits FFFF (65525), pero se usa uint32 para evitar rollover.
// Para nuestro ejemplo el 0x96 del red se convertirá en 0x9696.
//
// Por esto basta con truncar el número uint38 a un uint8 para convertirlo a un color hex.

func cargarFondoSimulacion(filePath string) (w int, h int, img image.Image) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err = png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	pixelesTotal := width * height

	fmt.Println("Dimensiones:", width, "x", height, "pixeles:", pixelesTotal)

	// for i := 0; i < pixelesTotal; i++ {
	// 	x := i % width
	// 	y := i / width
	// 	color := rgbaToHex(img.At(x, y))
	// 	switch color {
	// 	case "960000":
	// 		// fmt.Printf("Pixel en (%v,%v) es rojo\n", x, y)
	// 	case "0000AA":
	// 		// fmt.Printf("Pixel en (%v,%v) es azul\n", x, y)
	// 	}
	// }

	return width, height, img
}

func rgbaToHex(c color.Color) string {
	var r, g, b uint8

	r1, g1, b1, _ := c.RGBA()

	r = uint8(r1)
	g = uint8(g1)
	b = uint8(b1)

	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
