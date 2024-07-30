/*
package = proyecto = area de trabajo
La primer linea de cada package (paquete, programa), debe declarar a que paquete pertence.

package main <- pertenece al paquete main.

¿Por que llamar 'main' al paquete?

	Hay 2 tipos de paquetes:
		· Ejecutables: Crean un ejecutable, un programa que hace algo.
		· Reusables: Dependecias, librerias.
	El nombre del paquete determina si es ejecutable o resusable. La palabra 'main' crea un paquete de tipo ejecutable.

	package main -> go build -> main.exe
	package lalala -> go build -> nada!

	La palabra 'main' es SAGRADA.

Los ejectubles, ademas de tener la linea:
· package main
deben tener la funcion:
· func main () {}

¿Que signfica 'import fmt'?

	Importa un paquete, el paquete 'fmt' en este caso. O sea, permite que 'main' tenga acceso a todo el codigo y todas las funcionalidades del paquete 'fmt'.
	Por defecto 'main' NO tiene acceso a ningún otro paquete, hay que importar cada paquete que se desee utilizar.
	Para encontrar mas informacion sobre la libreria estandar: golang.org/pkg.

¿Que significa 'func'?

	func: Indica que esuna funcion.
	main: Nombre de la funcion.
	(): Lista de argumentos pasasdos a la funcion.
	{}: Cuerpo de la función.

	func main () {

	}

¿Como esta organizado main.go?

	package main	<- declaracion del paquete
	import 'fmt' 	<- importacion de paquetes   necesarios.
	func main () { 	<- declaracion de funciones.
		fmt.Println (hola mundo!)
	}
*/
package main

import "fmt"

func main() {
	fmt.Printf("Hola Mundo!\n")
}

/*
	¿Como se corre un programa?
		directorio-del-programa$ go build programa.go 	-> Compila el programa, genera ejecutable.
		directorio-del-programa$ go run programa.go		-> Compila y ejecuta, no genera ejecutable.
		directorio-del-programa$ go fmt programa.go		-> Formatea todo el codigo de todos los archivos del directorio.
		directorio-del-programa$ go install programa.go	-> Compila e instala un paquete.
		directorio-del-programa$ go get programa.go 	-> Descarga un paquete.
		directorio-del-programa$ go test programa.go 	-> Corre un test asociado con un paquete.
*/
