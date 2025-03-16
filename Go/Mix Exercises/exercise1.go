// Calculadora de áreas de un círculo
// Descripción:
// Escribe un programa que calcule el área de un círculo. Declara una constante para el valor de π (pi) y crea una función que reciba el radio (asegurándote de que sea un valor positivo) y retorne el área.
// Requisitos:

// Uso de variables y constantes.
// Manejo de tipos de datos primitivos (por ejemplo, float64 para el radio).
// Uso de condicionales (if) para validar que el radio sea mayor que cero.
// Implementación de una función que realice el cálculo.
// Temas involucrados: Sintaxis, variables y constantes, tipos primitivos, estructuras de control, funciones.

package main

import (
	"math"
)

const pi float64 = 3.14

func Area_circulo(radio float64) float64 {
	if radio <= 0 {
		return 0.0
	}
	area := pi * math.Pow(radio, 2)
	return area

}
