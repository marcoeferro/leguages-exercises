// Contador de Palabras en una Frase
// Descripción:
// Escribe un programa que reciba una cadena de texto, la separe en palabras (puedes utilizar funciones de la librería strings) y cuente la frecuencia de cada palabra almacenando el resultado en un map.
// Requisitos:

// Uso de maps para almacenar el conteo de palabras.
// Empleo de ciclos para recorrer las palabras y actualizar el map.
// Temas involucrados: Maps, manejo de strings, estructuras de control.

package main

import (
	// "fmt"
	"strings"
)

// func main() {
// 	// Example usage
// 	result := Frecuancia_Strings("hello world hello")
// 	fmt.Println(result) // Output: map[hello:2 world:1]
// }

func Frecuancia_Strings(cadena string) map[string]int {
	var array []string

	array = strings.Fields(cadena)

	frecuencias := make(map[string]int)

	for _, values := range array {
		value := strings.ToLower(values)
		_, exist := frecuencias[value]
		if exist {
			frecuencias[value]++
		} else {
			frecuencias[value] = 1
		}
	}
	return frecuencias
}
