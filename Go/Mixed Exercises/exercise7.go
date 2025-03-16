// // Cálculo Concurrente de Factoriales con Goroutines y Channels
// // Descripción:
// // Escribe un programa que, dada una lista de números, lance una goroutine para calcular el factorial de cada uno. Cada goroutine debe enviar su resultado a través de un channel, y la función principal debe recolectar y mostrar todos los resultados.
// // Requisitos:

// // Uso de goroutines para realizar cálculos en paralelo.
// // Empleo de channels para la comunicación entre las goroutines y la función principal.
// // Temas involucrados: Goroutines, channels, funciones.

package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Función para calcular el factorial de un número
// func factorial(n int, ch chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Marca la goroutine como terminada al finalizar

// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}
// 	ch <- result // Envía el resultado al channel
// }

// func main() {
// 	// Lista de números para calcular factoriales
// 	numbers := []int{5, 7, 3, 10, 4}

// 	// Crear un channel para recibir los resultados
// 	ch := make(chan int, len(numbers))

// 	// WaitGroup para esperar a que todas las goroutines terminen
// 	var wg sync.WaitGroup

// 	// Lanzar una goroutine para cada número
// 	for _, num := range numbers {
// 		wg.Add(1) // Añadir una goroutine al WaitGroup
// 		go factorial(num, ch, &wg)
// 	}

// 	// Goroutine adicional para cerrar el canal cuando todas las goroutines terminen
// 	go func() {
// 		wg.Wait()
// 		close(ch) // Cierra el canal después de que todas las goroutines hayan finalizado
// 	}()

// 	// Recolectar y mostrar los resultados
// 	fmt.Println("Resultados de los factoriales:")
// 	for result := range ch {
// 		fmt.Println(result)
// 	}
// }
