**Enunciado del Problema:**

Implementa un sistema que simule el trabajo de **Santa Claus**, sus **renos** y los **elfos** en un taller, utilizando gorutinas, canales y sincronización en Go. El sistema debe cumplir con los siguientes requerimientos:

1. **Santa Claus:**
   - Comienza en estado dormido y solo se despierta cuando:
     - **Tres elfos** tienen problemas al fabricar juguetes y necesitan ayuda.
     - **Todos los renos** han llegado al taller.
   - Una vez despierto, Santa resuelve los problemas de los elfos o prepara los renos.

2. **Elfos:**
   - Hay 12 elfos trabajando en la fabricación de juguetes.
   - Cada elfo tiene un 33% de probabilidad de tener problemas mientras trabaja.
   - Cuando tres elfos tienen problemas simultáneamente, despiertan a Santa para que los ayude.

3. **Renos:**
   - Hay 9 renos que llegan al taller con un intervalo de tiempo aleatorio (entre 5 y 7 segundos).
   - Una vez que todos los renos han llegado, despiertan a Santa.

4. **Sincronización:**
   - Santa Claus no puede atender a los elfos y a los renos al mismo tiempo.
   - Utiliza mecanismos de sincronización (`sync.Mutex`) y comunicación mediante canales para manejar las interacciones.

5. **Mensajes de registro:**
   - El sistema debe registrar en tiempo real las acciones y estados:
     - Llegada de renos.
     - Progreso de los elfos.
     - Estados de Santa (dormido, despierto, ayudando).

6. **Finalización:**
   - El programa termina cuando todos los renos han llegado y Santa ha ayudado a los elfos según lo requerido.

**Ejemplo de salida esperada:**

```plaintext
Santa is sleeping...
elf 1 crafting a toy
toy of elf 1 is done
elf 2 crafting a toy
The toy of elf 2 is broken, he needs help
...
reindeer 1 arrived!
...
Santa is up now...
Santa fixing the toys
Problem fixed in 3.5 seconds
Santa is sleeping...
...
all the reindeers are here
Santa is ready to go!
```