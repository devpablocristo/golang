## ¿Qué es una rune?

* En Go, una rune es un alias de `int32` que ocupa 32 bits y representa un punto de código Unicode.
* Se define mediante literales entre comillas simples, por ejemplo `'a'` o `'😊'`, y su tipo es `rune`.

## Diferencia entre byte y rune

* Un **byte** es un alias de `uint8` que representa un octeto de datos, ideal para texto ASCII o para manipulaciones de bajo nivel.
* Una **rune** puede ocupar de 1 a 4 bytes al codificarse en UTF-8, permitiendo representar caracteres multibyte de diferentes idiomas y símbolos complejos.
* Al indexar directamente una cadena con `s[i]`, Go devuelve un valor de tipo `byte`, no una rune.

## Cómo manejar runas en Go

* Para procesar texto Unicode correctamente, primero convierte la cadena a un slice de runas con `[]rune(s)`, de modo que cada elemento corresponda a un carácter completo y no a un byte suelto.
* También puedes iterar directamente sobre la cadena usando `for i, r := range s`; en cada iteración `r` es una rune (punto de código) y `i` es la posición en bytes.

## Ejemplos básicos

* **Convertir a slice de runas:**

  ```go
  runes := []rune(s) // convierte s a []rune
  ```

  Esto garantiza que cada índice de `runes` sea un carácter Unicode completo.
* **Recorrer runas con range:**

  ```go
  for i, r := range s {
      fmt.Printf("Índice: %d, rune: %c\n", i, r)
  }
  ```

  Así puedes imprimir cada carácter Unicode en su forma legible.


  ---


Portugues:


## O que é uma rune?

* Em Go, uma **rune** é um alias de `int32` que ocupa 32 bits e representa um ponto de código Unicode.
* É definida por literais entre aspas simples, por exemplo `'a'` ou `'😊'`, e seu tipo é `rune`.

## Diferença entre byte e rune

* Um **byte** é um alias de `uint8` que representa um octeto de dados, ideal para texto ASCII ou para manipulações de baixo nível.
* Uma **rune** pode ocupar de 1 a 4 bytes ao ser codificada em UTF-8, permitindo representar caracteres multibyte de diferentes idiomas e símbolos complexos.
* Ao indexar diretamente uma string com `s[i]`, Go retorna um valor do tipo `byte`, não uma rune.

## Como lidar com runas em Go

* Para processar texto Unicode corretamente, primeiro converta a string em um slice de runas com `[]rune(s)`, de modo que cada elemento corresponda a um caractere completo, e não a um byte isolado.
* Você também pode iterar diretamente sobre a string usando `for i, r := range s`; em cada iteração, `r` é uma rune (ponto de código) e `i` é a posição em bytes.

## Exemplos básicos

* **Converter para slice de runas:**

  ```go
  runes := []rune(s) // converte s para []rune
  ```

  Isso garante que cada índice de `runes` seja um caractere Unicode completo.

* **Percorrer runas com range:**

  ```go
  for i, r := range s {
      fmt.Printf("Índice: %d, rune: %c\n", i, r)
  }
  ```

  Assim você pode imprimir cada caractere Unicode de forma legível.

