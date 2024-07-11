**Enunciado del Ejercicio: Suma Máxima en Hourglasses**

Dada una matriz bidimensional de tamaño 6x6, implementa una función en Go que encuentre y devuelva la suma máxima de todos los posibles "hourglasses" dentro de la matriz. Un "hourglass" se define como un conjunto de números en forma de reloj de arena, como se muestra a continuación:

```
a b c
  d
e f g
```

Por ejemplo, en la siguiente matriz 6x6:

```
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
```

Hay varios "hourglasses" posibles, y sus sumas son:

```
2 4 4
  2
1 2 4

1 1 0
  4
0 2 4

0 0 0
  2
0 1 2

0 0 2
  0
0 1 2

0 0 0
  2
0 1 2

0 0 1
  2
0 4 0
```

La función debería devolver la suma máxima entre todas estas "hourglasses".

**Firma de la Función:**
```go
func HighestHourglassSum(matrix [6][6]int) int {
    // Tu implementación aquí
}
```

**Ejemplo de Uso:**
```go
matrix := [6][6]int{
    {1, 1, 1, 0, 0, 0},
    {0, 1, 0, 0, 0, 0},
    {1, 1, 1, 0, 0, 0},
    {0, 0, 2, 4, 4, 0},
    {0, 0, 0, 2, 0, 0},
    {0, 0, 1, 2, 4, 0},
}

result := HighestHourglassSum(matrix)
fmt.Println(result) // Debería imprimir la suma máxima de los "hourglasses"
```

**Restricciones:**
- La matriz siempre será de tamaño 6x6.
- Los elementos de la matriz serán números enteros.