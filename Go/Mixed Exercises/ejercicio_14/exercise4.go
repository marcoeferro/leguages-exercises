// Gestión de Productos con Structs y Métodos
// Descripción:
// Define una estructura Producto con campos como nombre (string), precio (float64) y stock (int). Implementa métodos asociados para:

// Actualizar el stock (por ejemplo, al vender o reponer unidades).
// Aplicar un descuento al precio.
// Requisitos:
// Uso de structs para modelar datos.
// Implementación de métodos, utilizando punteros cuando sea necesario para modificar la estructura original.
// Temas involucrados: Structs, métodos, punteros.

package main

import "fmt"

type Product struct {
	name   string
	precio float64
	stock  int
}

// func main() {
// 	p1 := Product{"Cereal", 100, 5}
// 	p1.UpdateStock(-5)
// 	p1.Discount(35)
// }

func (p *Product) UpdateStock(stock int) {
	p.stock += stock
	fmt.Println("The stock of the product", p.name, "it's", p.stock)
}

func (p *Product) Discount(dis float64) {
	p.precio -= dis
	fmt.Println("El nuevo precio es", p.precio)
}
