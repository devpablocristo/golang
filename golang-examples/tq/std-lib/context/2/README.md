Este código en Go ilustra el uso del paquete `context` para enriquecer un contexto con valores y controlar la ejecución de una operación (representada por la función `doStuff`) con un tiempo de espera o "timeout". A continuación, el código está explicado con comentarios detallados para entender cada parte:

- La función `enrichCtx` demuestra cómo adjuntar datos adicionales a un contexto, que luego puede ser accedido en otras partes de la aplicación que reciben este contexto. Sin embargo, el valor añadido ("joe") no se usa en ninguna parte del ejemplo proporcionado.

- La función `doStuff` realiza una operación en un bucle infinito hasta que recibe una señal de cancelación a través del contexto. Utiliza `select` para esperar a la cancelación y proceder con la lógica de salida o continuar con su operación predeterminada.

- En `main`, se crea un contexto con un timeout de 2 segundos y se enriquece con un valor adicional. Este contexto se pasa a `doStuff` ejecutado en una goroutine. La ejecución de `main` espera a que el contexto se cancele debido al timeout y luego procede a imprimir un mensaje final antes de finalizar.

Este ejemplo muestra el uso de contextos para controlar el tiempo de ejecución de operaciones concurrentes y cómo pasar datos a través del contexto. La inmutabilidad de los contextos (es decir, el hecho de que `enrichCtx` realmente crea un nuevo contexto en lugar de modificar el existente) es crucial para entender cómo trabajar con ellos de manera efectiva en Go.