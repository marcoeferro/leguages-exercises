// FizzBuzz usando switch
// Descripción:
// Implementa el ejercicio clásico FizzBuzz:

// Imprime los números del 1 al 100.
// Si un número es divisible por 3, imprime "Fizz".
// Si es divisible por 5, imprime "Buzz".
// Si es divisible por ambos, imprime "FizzBuzz".
// Requisitos:
// Uso de un ciclo for para iterar sobre los números.
// Empleo de la estructura switch (o if/else) para determinar qué imprimir en cada caso.
// Temas involucrados: Estructuras de control (for, switch).

package main

import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 15, 15, 15, 10, 6}
// 	FizzBuzz(numbers)
// }

func FizzBuzz(numbers []int) {
	for _, value := range numbers {
		if value%3 == 0 {
			fmt.Print("Fizz")
		}
		if value%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Println()
	}

}
