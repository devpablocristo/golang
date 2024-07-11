/*
	Neste código, temos duas partes principais:
	1. A função `ehPrimo`, que recebe um número inteiro e retorna `true` se o número fosse primo e `false` se não fosse.  A função usa uma abordagem eficiente que verifica apenas divisores até a raiz quadrada do número.
	2. A função `main`, onde um número é definido e a função `ehPrimo` é chamada para verificar se é primo ou não, imprimindo o resultado no console.

	Você pode modificar o valor de `numero` na função `main` para conferir diferentes números.

	Um número primo é aquele que é divisível apenas por 1 e por ele mesmo.

	Por exemplo: 1, 3, 7, 11, ...
*/

package main

import (
	"fmt"  // Importa o pacote fmt para operações de entrada/saída
	"math" // Importa o pacote math para operações matemáticas
)

// ehPrimo é uma função que determina se um número é primo.
func ehPrimo(numero int) bool {
	// Se o número fosse menor ou igual a numero um, não é primo.
	if numero <= 1 {
		return false
	}

	// Calcula a raiz quadrada do número para reduzir o número de divisões necessárias.
	// Em vez de verificar todos os números de 2 até n - 1 para ver se n é divisível por algum deles
	// (o que seria muito ineficiente para números grandes), você só precisa verificar até a raiz
	// quadrada de n. Se não encontrar divisores até esse ponto, pode ter certeza de que n é primo.
	limite := int(math.Sqrt(float64(numero)))

	// Começa do numero dois até o valor do limite (raiz quadrada do número).
	for i := 2; i <= limite; i++ {
		fmt.Printf("numero: %d mod i: %d\n", numero, i)
		fmt.Println("limite:", limite)
		fmt.Println("---------------")

		// Se o número fosse divisível por i, não é primo.
		if numero%i == 0 {
			return false
		}
	}

	// Se nenhum divisor foi encontrado, o número é primo.
	return true
}

func main() {
	//var numero int
	//fmt.Print("Digite um número: ")
	//fmt.Scan(&numero)

	// Aqui um número estático é definido para conferir a função ehPrimo.
	//numero := 10 // Altere este valor para testar com diferentes números.
	numero := 11 // Exemplo de como testar outro número.

	// Chama a função ehPrimo e exibe o resultado.
	if ehPrimo(numero) {
		fmt.Printf("%d é um número primo.\n", numero)
	} else {
		fmt.Printf("%d não é um número primo.\n", numero)
	}
}
