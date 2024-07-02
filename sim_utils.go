package main

import (
	"math/rand"
)

// ================================================================ //
// ========== RANDOM ============================================== //

// Genera número aleatorio entre min y max
func randomBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// ================================================================ //
// ========== DISTANCIA =========================================== //

// Convierte metros a pixeles.
func ToPixels(mtrs float64) float64 {
	return mtrs / metrosPorPixel
}

func ToPixels32(mtrs float32) float32 {
	return mtrs / metrosPorPixel
}

func ToPixelsInt(mtrs float64) int {
	return int(mtrs / metrosPorPixel)
}

// Convierte metros desde pixeles.
func ToMetros(px float64) float64 {
	return px * metrosPorPixel
}

// ================================================================ //
// ========== VELOCIDAD =========================================== //

// Retorna m/s desde km/h
func KilometrosPorHora(MetrosPorSegundo float64) int {
	return int(MetrosPorSegundo * 18 / 5)
}
