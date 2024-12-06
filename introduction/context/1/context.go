package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//ctx := getContext()
	// value := ctx.Value(1)
	// fmt.Println(value)

	fmt.Println("server started!")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greeting", middleware(greetingHandler))
	http.ListenAndServe(":9191", nil)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	traceID := r.Context().Value("traceID")
	w.Write([]byte(traceID.(string)))
}

func middleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		//De esta forma se puede pasar cualquier valor dentro de un middleware a la logica de negocios
		ctx = context.WithValue(ctx, "traceID", "a1s2d3f4g5")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// La Request tiene un contexto que es muy importante, es el ciclo
// de vida de la llamada que tenemos en el webserver.
// Comienza desde le momento que llega al webserver, con toda la metadata
// de una Request.
// La Request pasa por toda la logica de negocios hasta que termina y se devuelve un
// devuelve al client un valor con un estado (200, 201, 400, etc).
// Ahi es donde termina el ciclo de vida de la Request.
// Durante todo el ciclo de vida, tenemos acceso al ese contexto.
// Se puede escribir, recibir valores y tambien recibir señales que se
// mande a ese contexto.
// Si se una middleware, la Request siempre es la misma, por eso se manda
// un puntero. Dentro del middleware se puede propagar el contexto. Un
// ejemplo comun es el trace id de una llamada.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// contexto de la request
	ctx := r.Context()
	var err error
	// defer, se ejecuta al finalizar la función
	defer fmt.Println("afer signal interrupt")

	// donde esta la "magia"
	select {
	// se esperan 10 segundos para poder apretar ctrl+c
	case <-time.After(10 * time.Second):
	// Cuando se interrupe la señal (ctrl+c) se selecciona este caso.
	case <-ctx.Done():
		// se obtiene el error
		err = ctx.Err()
	}

	if err != nil {
		// se imprime el error por stdout
		fmt.Printf("error is %s \n", err.Error())
	}
}

func getContext() context.Context {

	// manda el contexto de la goroutiene que
	// es dueña de la función
	ctx := context.Background()

	// sirve para hacer pruebas
	// es un empty context
	// si en un test se necesita pasar el contexto
	// se puede pasar el TODO
	// siempre que se llamam, nunca es nil y siempre estará vacio
	// ctxDone := context.TODO()

	// necesita un contexto padre
	// key y value funciona como un map, aunque no se si es una map internamente
	ctxWithValue := context.WithValue(ctx, 1, 999)

	return ctxWithValue
}

func testGetContext() {
	// goroutine main is the holder
	ctx := getContext()
	value := ctx.Value(1)
	fmt.Println(value)
}
