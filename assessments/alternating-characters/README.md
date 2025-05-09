**Enunciado del Problema:**

Dado un string `s`, se requiere desarrollar un algoritmo que cumpla con las siguientes funcionalidades relacionadas con los caracteres `'A'` y `'B'`:

1. **Filtrar el string**: Crear una función que elimine todos los caracteres del string excepto `'A'` y `'B'`. Los caracteres restantes deben conservar su orden relativo.

2. **Contar pares consecutivos**: Implementar una función que cuente cuántos pares consecutivos de caracteres `'A'` o `'B'` existen en el string resultante del filtrado. Un par consecutivo se define como dos caracteres iguales seguidos, por ejemplo: `'AA'` o `'BB'`.

3. **Eliminar pares consecutivos**: Diseñar una función que elimine todos los pares consecutivos de caracteres `'A'` y `'B'` en el string filtrado, dejando un solo carácter en su lugar. Por ejemplo, `'AABBA'` se transformaría en `'ABA'`.

4. **Optimizar con una segunda versión** *(opcional)*: Crear una implementación alternativa para las funciones anteriores usando métodos avanzados o estructuras eficientes, como `strings.Builder` o `strings.Map`, para mejorar la claridad y/o el rendimiento.

### Ejemplo de entrada y salida:

#### Entrada:
```plaintext
abcdefghAijklmnoApqBrstuvwBxyzAawqwqeBBjjj<
```

#### Salida esperada:
```plaintext
Cleaned String: AABBAABB
Count of Extras AB: 3
String after removing extras AB: ABAB
```

### Restricciones:
1. El string de entrada puede contener caracteres alfabéticos, numéricos y símbolos.
2. Se debe ignorar la distinción entre mayúsculas y minúsculas para los caracteres `'A'` y `'B'` (si se desea agregar esta funcionalidad).
3. Se deben implementar funciones independientes para cada paso.

El algoritmo debe ser eficiente y fácil de mantener, utilizando buenas prácticas de codificación.