package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// fmt.Println() se usa para imprimir mensajes simples en la consola.
	// Es útil para mensajes de depuración o para informar el estado del programa.
	fmt.Println("Inicio del programa")

	// log.Println() imprime mensajes con fecha y hora.
	// Es más informativo que fmt.Println, útil para registros (logs) más detallados.
	log.Println("Este es un mensaje de log estándar")

	// Ejemplo de una función que puede generar un error.
	err := doSomethingRisky()
	if err != nil {
		// log.Fatalln() imprime un mensaje de error y luego llama a os.Exit(1),
		// terminando el programa inmediatamente.
		// Utilizado para errores críticos donde no se puede continuar la ejecución.
		log.Fatalln("Error fatal:", err)
	}

	// os.Exit() termina el programa con un código de estado.
	// No se llaman a las funciones defer pendientes, por lo que se debe usar con cuidado.
	// os.Exit(1)

	// log.Panicln() imprime un mensaje de error y luego ejecuta un panic.
	// A diferencia de log.Fatalln, los defer se ejecutarán antes de terminar.
	// log.Panicln("Error de pánico:", err)

	// panic() genera un pánico inmediatamente.
	// Utilizado en situaciones donde el programa no puede recuperarse.
	// panic("Algo fue terriblemente mal")

	// Recuperación de un pánico. Este bloque defer se ejecutará si hay un panic.
	defer func() {
		if r := recover(); r != nil {
			// recover() detiene el pánico y permite que el programa continúe.
			// Es útil para manejar errores en tiempo de ejecución y evitar la terminación del programa.
			fmt.Println("Recuperado de un pánico:", r)
		}
	}()

	// Generando un pánico para la demostración de recover.
	panic("Demostración de pánico")

	// fmt.Errorf() y errors.New() se utilizan para crear errores personalizados.
	// fmt.Errorf() es útil cuando necesitas formatear el mensaje de error.
	// errors.New() se usa para crear un error simple con un mensaje específico.
	err = doAnotherRiskyThing()
	if err != nil {
		// Imprimiendo el error creado con fmt.Errorf() o errors.New().
		fmt.Println("Error encontrado:", err)
	}

	fmt.Println("Fin del programa")
}

// doSomethingRisky simula una operación que puede fallar y devolver un error.
func doSomethingRisky() error {
	// Aquí se podría incluir lógica real que pueda fallar.
	return errors.New("algo salió mal en doSomethingRisky")
}

// doAnotherRiskyThing simula otra operación que puede fallar.
func doAnotherRiskyThing() error {
	// Esta función utiliza fmt.Errorf para añadir más contexto al error.
	return fmt.Errorf("error en doAnotherRiskyThing: %w", errors.New("falla específica"))
}
