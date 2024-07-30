package main

import (
	"context" // Importa el paquete context para el manejo de la cancelación.
	"fmt"     // Importa el paquete fmt para imprimir mensajes en la consola.
	"time"    // Importa el paquete time para utilizar temporizadores.
)

func main() {
	// Crea un contexto cancelable con context.WithCancel. Este contexto se deriva de context.Background(),
	// que es el contexto raíz para situaciones en las que no hay un contexto más específico.
	ctx, stop := context.WithCancel(context.Background())

	// Inicia una goroutine anónima que espera una entrada del usuario. Cuando el usuario presiona Enter,
	// se llama a la función stop() para cancelar el contexto.
	go func() {
		fmt.Scanln() // Espera la entrada del usuario (hasta que presione Enter).
		stop()       // Cancela el contexto, lo que enviará una señal a cualquier operación que lo esté escuchando.
	}()

	// Llama a la función heartbeat, pasando el contexto cancelable.
	heartbeat(ctx)
}

// La función heartbeat simula latidos del corazón ("heartbeat") imprimiendo "beat" en la consola cada segundo,
// hasta que el contexto se cancela.
func heartbeat(ctx context.Context) {
	tick := time.Tick(time.Second) // Crea un ticker que envía un mensaje cada segundo.

	for {
		select {
		case <-tick: // Cada vez que el ticker envía un mensaje, pasa por este caso.
		case <-ctx.Done(): // Si el contexto se cancela (por la llamada a stop() en la goroutine anónima),
			return // la función retorna, deteniendo el bucle y finalizando la simulación de los latidos.
		}
		fmt.Println("beat") // Imprime "beat" en la consola, representando un latido.
	}
}
