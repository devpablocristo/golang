**Enunciado del Problema:**

Dadas dos cadenas de texto, desarrollar un programa que permita analizar las diferencias y coincidencias entre ambas, cumpliendo con los siguientes requerimientos:

1. **Ordenar las cadenas:** Procesar las cadenas para ordenarlas alfabéticamente.

2. **Encontrar coincidencias:** Identificar los caracteres comunes entre ambas cadenas, considerando todas sus ocurrencias.

3. **Eliminar duplicados:** Filtrar los caracteres comunes para que cada uno aparezca solo una vez, y calcular cuántas ocurrencias duplicadas fueron eliminadas.

4. **Calcular un contador final:** Determinar el total de caracteres únicos combinados entre las dos cadenas, después de eliminar duplicados de las coincidencias.

5. **Mostrar resultados:**
   - Los caracteres únicos comunes de la primera cadena en la segunda.
   - Los caracteres únicos comunes de la segunda cadena en la primera.
   - El valor del contador final.

---

### **Ejemplo de entrada:**

```plaintext
Entrada:
Cadena 1: "aabcd"
Cadena 2: "aefgha"
```

### **Salida esperada:**

```plaintext
Salida:
Comunes en Cadena 1: "aa"
Comunes en Cadena 2: "aa"
Contador final: 8
```

---

### **Restricciones:**
1. Las cadenas pueden contener caracteres alfabéticos en mayúsculas o minúsculas y deben tratarse como sensibles a mayúsculas.
2. No se garantiza que las cadenas estén libres de duplicados inicialmente.
3. Los resultados no requieren conservar el orden original de las cadenas.