### **6️⃣ Manejo de errores (`errors.New`)**
**Ejercicio:**  
Escribe una función `Divide(a, b float64) (float64, error)` que realice la división `a / b`, pero retorne un error si `b == 0`.  

**Ejemplo:**  
```go
Divide(10, 2)  // Devuelve (5, nil)
Divide(5, 0)   // Devuelve (0, error "no se puede dividir por cero")
```
