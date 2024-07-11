Ejercicio 10: 

Objetivo: Encontrar la salida de un laberinto

Busqueda.

Entrada.

Salida.

arriba, abajo, izq, der


-> izq a der = avanzar 

como sabemos cuales son las paredes de del laberinto? 

el laberidnto es una matriz de 10x10 

'x' respresneta bloqueado (borde laberinto)
'0' representa bloqueado

'1' representa desbloqueado

'p' player
's' representa salida <- meta
    c
   xxxxxxxxxxxxxxxxxxxxxxxxxx
 f x000000111111111111111111x 0
   p111111100000000100000000x 1
   x000000000000000100000000x 2
   x000000000000000100000000x 3
   x000000000000000100000000x 4
   x000000001111111111111111x 5
   x000000000000000100000000x 6
   x000000000000000100000000x 7 
   x000000000000000100000000x 8
   x000000000000000s00000000x 9
   xxxxxxxxxxxxxxxxxxxxxxxxxx
    0123456789012345678901234
est

laberinto[10][25]

dividir el problema en 2 partes

1 creador del laberinto <----- ya esta listo

el laberinto es una matriz de 10x10
el vertice izq, superior (de los 0, las x no se tienen en cuenta, solo como bloqueante)
laberinto[0][0]
el vertice der, inferior
laberinto [10][25]

laberinto[x][y] <--- donde estamos parados


2 buscar la salida del laberinto  <----- ESTE ES QUE RESOLVEMOS


solo podemos avanzar en los casillers con valor 1
cuando se avanza el casillero que se deja pasa de valer a valer 1

010
1p1
010

cada turno se avanza hacia adelante,
    si es posible, se avanza, se resetea a 1 la posicion anterior, se sobreescribe con p la posicion actual
si no es posible, se avanza hacia abajo,
    si es posible, se avanza, se resetea a 1 la posicion anterior, se sobreescribe con p la posicion actual
si no es posible, se avanza hacia arriba,
    si es posible, se avanza, se resetea a 1 la posicion anterior, se sobreescribe con p la posicion actual



si no, se retroce <--- CASO especial
    se retrocede hasta ultimaPosMasDeUnCamino 





     I
     0  
1111p10 ->
     0
     I


cada vez que el player da un paso, escanea los caminos disponibles, y si hay mas de 1 los guarda, 
estos caminos disponibles, se marcaran como "probados" si ya se camino por ellos.



type Player struct {
    x int 
    y int 
    ultimaPos [2]int //desde donde venimos, para tener en cuenta el caso especial de retroceder
    ultimaPosMasDeUnCamino 
}

type  ultimaPosMasDeUnCamino struct {
    arriba bool
    abajo bool 
    izq   bool
    der  bool 
    ultimaPos [2]int
}


col y fil
rand 5 - 20
 

00000
00000
00000
00000
00000
00000

00000000000
00000000000
00000000000
00000000000
00000000000
00000000000



xxxxxxxxxxxxxxxxxx
000000000100000000x
000000000100000000x
000000000100000000x
11111111111111111sx
000000000000000000x
000000000000000000x
xxxxxxxxxxxxxxxxxx



Tarea 1: dibujar matriz ceros, con col y fil random, entre 5 y 20 chars



Tarea 2: sobre escribir en la matriz de 0, "camino" de 1s, en linea recta
Tarea 3: sobre escribir en la matriz de 0, "camino" de 1s, con 1 esquina 
Tarea 4: sobre escribir en la matriz de 0, "camino" de 1s, con 2 esquina 
Tarea 5: sobre escribir en la matriz de 0, "camino" de 1s, con 3 esquina
Tarea 6: colocar al "player" en la entrada del laberinto.
Tarea 7: dibujar bordes ("x") del laberinto. 


