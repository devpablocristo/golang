**Assessment: Saltar en las nubes**

---

### Descripción

Tienes un arreglo de nubes representado por enteros, donde:

* `0` = nube segura
* `1` = nube peligrosa (no se puede pisar)

Tu objetivo es llegar desde la primera nube hasta la última realizando el **mínimo número de saltos** posible. Desde una nube segura puedes saltar **una** o **dos** posiciones hacia adelante, pero **solo** si la nube de destino también es segura.

Implementa en Go la función:

```go
// saltarEnNubes recibe un slice de enteros c y devuelve
// el número mínimo de saltos necesarios para alcanzar
// el final. Se asume que c[0] = 0 y c[n-1] = 0.
func saltarEnNubes(c []int) int
```

Y un programa `main` que:

1. Lee de la entrada estándar:

   * Un entero `n` (2 ≤ n ≤ 100), la cantidad de nubes.
   * Una línea con `n` enteros `0` o `1`, separados por espacios.
2. Llama a `saltarEnNubes(c)` y
3. Imprime el resultado (un entero) en la salida estándar.

---

### Formato de entrada

```
n
c[0] c[1] c[2] … c[n-1]
```

* `n`: número de nubes.
* `c[i]`: estado de la nube i (0 = segura, 1 = peligrosa).

### Formato de salida

Un único entero: la **cantidad mínima de saltos** para ir desde la nube 0 hasta la nube n-1.

---

### Ejemplo

#### Entrada

```
7
0 1 0 0 0 1 0
```

#### Salida

```
4
```

**Explicación**
Con `c = [0,1,0,0,0,1,0]` la secuencia de saltos mínima es:

* Salto de índice 0 → 2
* Salto de índice 2 → 4
* Salto de índice 4 → 6
  Total = 3 saltos.

*(En este ejemplo, la salida es 3; el ejemplo de arriba ilustra el formato.)*

---

### Restricciones y casos de prueba

* Siempre empieza y termina en nubes seguras (`c[0] = c[n-1] = 0`).
* Garantiza que exista al menos una forma de llegar al final.
* Prueba tu solución con casos extremos, por ejemplo:

  * `n = 2`, `c = [0,0]` → respuesta `1`
  * Todos ceros: `c = [0,0,0,0,…]`
  * Patrones alternados de 0 y 1.

---

### Puntos de evaluación

* **Correctitud**: devuelve el número mínimo de saltos en todos los casos.
* **Claridad**: código legible, nombres descriptivos, comentarios si es necesario.
* **Eficiencia**: recorrido en O(n) y uso O(1) de memoria adicional.
* **Buenas prácticas Go**: manejo de errores de lectura, formateo con `gofmt`, tests unitarios opcionales.

---

¡Mucho éxito! Implementa `saltarEnNubes`, escribe tu `main` y verifica tu solución con diferentes casos de prueba.


---

---

**Avaliação: Saltar nas Nuvens**

---

### Descrição

Você recebe um slice de nuvens representado por inteiros, onde:

* `0` = nuvem segura
* `1` = nuvem perigosa (não pode pisar)

Seu objetivo é ir da primeira à última nuvem com o **mínimo número de saltos** possível. De uma nuvem segura você pode saltar **1** ou **2** posições para frente, mas somente se a nuvem de destino também for segura.

Implemente em Go a função:

```go
// saltarEmNuvens recebe um slice c e retorna
// o número mínimo de saltos para alcançar o final.
// Assume-se que c[0] = 0 e c[len(c)-1] = 0.
func saltarEmNuvens(c []int) int
```

E um programa `main` que:

1. Lê da entrada padrão:

   * Um inteiro `n` (2 ≤ n ≤ 100): quantidade de nuvens.
   * Uma linha com `n` inteiros (`0` ou `1`), separados por espaço.
2. Chama `saltarEmNuvens(c)`.
3. Imprime o resultado (um inteiro) na saída padrão.

---

### Formato de entrada

```
n
c[0] c[1] c[2] … c[n-1]
```

* `n`: número de nuvens.
* `c[i]`: estado da nuvem i (0 = segura, 1 = perigosa).

### Formato de saída

Um único inteiro: a **quantidade mínima de saltos** para ir da nuvem 0 até a nuvem n-1.

---

### Exemplo

**Entrada**

```
7
0 1 0 0 0 1 0
```

**Saída**

```
3
```

**Explicação**
Para `c = [0,1,0,0,0,1,0]`, a sequência mínima de saltos é:

* 0 → 2
* 2 → 4
* 4 → 6
  Total = 3 saltos.

---

### Restrições e casos de teste

* Inicia e termina em nuvens seguras (`c[0] = c[n-1] = 0`).
* Garante-se que exista ao menos um caminho até o final.
* Teste sua solução com casos extremos, por exemplo:

  * `n = 2`, `c = [0,0]` → resposta `1`
  * Todos zeros: `c = [0,0,0,0,…]`
  * Padrões alternados de 0 e 1.

---

### Critérios de avaliação

* **Corretude**: retorna o número mínimo de saltos em todos os cenários.
* **Clareza**: código legível, nomes descritivos e comentários, se necessário.
* **Eficiência**: complexidade O(n) e uso extra de memória O(1).
* **Boas práticas em Go**:

  * Tratamento de erros de leitura.
  * Código formatado com `gofmt`.
  * (Opcional) testes unitários e benchmarks.

---

Boa implementação e bons saltos!
