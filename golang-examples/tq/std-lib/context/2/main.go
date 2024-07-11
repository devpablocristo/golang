package main

import (
	"context" // Importa el paquete context para controlar cancelaciones y pasar valores.
	"fmt"
	"time"
)

// enrichCtx toma un contexto y le añade un valor, en este caso un "user" con el nombre "joe".
// Retorna un nuevo contexto derivado que incluye este valor.
func enrichCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", "joe")
}

// doStuff es una función que simula una tarea que se repite hasta que el contexto se cancela debido a un timeout.
func doStuff(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Si el contexto se cancela (por ejemplo, debido a un timeout),
			fmt.Println("timeout") // imprime "timeout" y
			return                 // sale de la función.
		default:
			fmt.Println("doing stuff") // En el caso por defecto, imprime "doing stuff".
		}
		time.Sleep(500 * time.Millisecond) // Espera 500 milisegundos antes de continuar con el siguiente ciclo.
	}
}

func main() {

	ctx := context.Background() // Comienza con un contexto vacío.

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // Crea un contexto con un timeout de 2 segundos.
	defer cancel()                                         // Asegura que la cancelación del contexto se llama para liberar recursos.

	ctx = enrichCtx(ctx) // Enriquece el contexto con un valor adicional, en este caso, un "user".
	go doStuff(ctx)      // Ejecuta doStuff en una goroutine, pasando el contexto enriquecido.

	select {
	case <-ctx.Done(): // Si el contexto se "hace", lo que sucederá tras alcanzar el timeout,
		fmt.Println("final timeout") // imprime "final timeout" y
		fmt.Println(ctx.Err())       // el error asociado al contexto, típicamente "context deadline exceeded".
	}

	time.Sleep(2 * time.Second) // Espera 2 segundos antes de finalizar el programa.
	// Este Sleep está aquí para asegurar que el programa no finalice antes de que la goroutine pueda imprimir su salida,
	// pero en este caso, es redundante debido al bloqueo anterior en ctx.Done() y podría ser eliminado.
}
