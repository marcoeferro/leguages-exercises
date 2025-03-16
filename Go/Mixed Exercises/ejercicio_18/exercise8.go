// Procesamiento Concurrente de Cadenas y Conteo de Frecuencia
// Descripción:
// Dado un slice de cadenas, lanza una goroutine para procesar cada cadena (por ejemplo, convertirla a mayúsculas). Luego, utiliza un channel para recolectar los resultados y almacena en un map la cantidad de ocurrencias de cada cadena procesada.
// Requisitos:

// Uso de slices y maps.
// Lanzar múltiples goroutines para procesar datos de forma concurrente.
// Recolectar resultados mediante channels y combinarlos en un map.
// Temas involucrados: Arrays y slices, maps, goroutines, channels.

package main

import (
	"fmt"
	"strings"
	"sync"
)

// Función para procesar una cadena (convertirla a mayúsculas)
func procesarCadena(cadena string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la goroutine como terminada al finalizar

	// Convertir la cadena a mayúsculas
	resultado := strings.ToUpper(cadena)

	// Enviar el resultado al channel
	ch <- resultado
}

func main() {
	// Slice de cadenas a procesar
	cadenas := []string{"hola", "mundo", "go", "hola", "concurrencia", "go", "go"}

	// Crear un channel para recibir los resultados procesados
	ch := make(chan string, len(cadenas))

	// WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup

	// Lanzar una goroutine para cada cadena
	for _, cadena := range cadenas {
		wg.Add(1) // Añadir una goroutine al WaitGroup
		go procesarCadena(cadena, ch, &wg)
	}

	// Goroutine adicional para cerrar el canal cuando todas las goroutines terminen
	go func() {
		wg.Wait()
		close(ch) // Cierra el canal después de que todas las goroutines hayan finalizado
	}()

	// Crear un map para almacenar la frecuencia de cada cadena procesada
	frecuencias := make(map[string]int)

	// Recolectar los resultados del channel y contar las frecuencias
	for resultado := range ch {
		frecuencias[resultado]++
	}

	// Mostrar las frecuencias
	fmt.Println("Frecuencia de cadenas procesadas:")
	for cadena, frecuencia := range frecuencias {
		fmt.Printf("%s: %d\n", cadena, frecuencia)
	}
}
