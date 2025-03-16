// Intercambio de Valores (Swap) con Punteros
// Descripción:
// Crea una función que reciba dos punteros a enteros y que intercambie sus valores.
// Requisitos:

// Uso de punteros para modificar directamente el contenido de las variables pasadas por referencia.
// Demostrar el paso de parámetros por referencia en una función.
// Temas involucrados: Punteros, funciones y parámetros.
package main

// import "fmt"

// func main() {
// 	point1, point2 := 1, 25
// 	fmt.Println("Before", point1, point2)

// 	SwapPointers(&point1, &point2)

// 	fmt.Println("After", point1, point2)
// }

func SwapPointers(pointer1, pointer2 *int) {
	*pointer1, *pointer2 = *pointer2, *pointer1

}
