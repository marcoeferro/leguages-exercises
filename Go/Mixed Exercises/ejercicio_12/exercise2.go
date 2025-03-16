// Máximo y Mínimo en un Slice de Enteros
// Descripción:
// Crea una función que reciba un slice de enteros y retorne dos valores: el máximo y el mínimo encontrados en el slice.
// Requisitos:

// Recorrer el slice utilizando un ciclo for.
// Uso de estructuras de control (if) para comparar y actualizar el valor máximo y mínimo.
// Manejo de arrays/slices y funciones con parámetros y retorno múltiple.
// Temas involucrados: Arrays y slices, estructuras de control, funciones y parámetros.

package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, -6565, 9999999999999999}
// 	max, min := Max_Min(slice)
// 	fmt.Println("MAXIMO", max, "Y MINIMO", min)
// }

func Max_Min(slice []int) (int, int) {
	max := slice[0]
	min := slice[0]
	for _, value := range slice {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max, min
}
