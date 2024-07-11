package main

import (
	"errors" // Importa el paquete 'errors' para crear errores.
	"fmt"    // Importa el paquete 'fmt' para la impresión de mensajes.
)

// addTwo añade dos a un valor dado. Si el valor es negativo, devuelve un error.
func addTwo(value int) (int, error) {
	if value < 0 {
		// Si el valor es negativo, no seguimos las reglas de negocio y devolvemos un error.
		// Usamos fmt.Sprintf para formatear el mensaje de error con el valor proporcionado.
		return value, errors.New(fmt.Sprintf("algo salió mal, el valor es %d:", value))
	} else {
		// Si el valor es positivo o cero, añadimos dos y devolvemos el resultado sin errores.
		return value + 2, nil
	}
}

func main() {
	// Llamamos a la función addTwo con un valor negativo para probar el manejo de errores.
	v, err := addTwo(-1)
	if err != nil {
		// Si hay un error (es decir, el valor era negativo), imprimimos el error.
		fmt.Print(err)
	} else {
		// Si no hay error, imprimimos el valor devuelto por la función.
		fmt.Print(v)
	}
}
