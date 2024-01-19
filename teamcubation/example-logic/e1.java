public class VerificadorPrimo {

    // Função para verificar se um número é primo
    public static boolean ehPrimo(int numero) {
        // Se o número for menor ou igual a 1, não é primo
        if (numero <= 1) {
            return false;
        }

        // Iterar de 2 até a raiz quadrada do número
        for (int i = 2; i <= Math.sqrt(numero); i++) {
            // Se o número for divisível por algum valor entre 2 e a raiz quadrada, não é primo
            if (numero % i == 0) {
                return false;
            }
        }

        // Se nenhum divisor foi encontrado, o número é primo
        return true;
    }

    public static void main(String[] args) {

        // Atribuir um valor diretamente à variável número
        int numero = 15; // Altere este valor para testar com diferentes números

        // Verificar se o número é primo usando a função ehPrimo
        if (ehPrimo(numero)) {
            // Se o número for primo, imprime que é um número primo
            System.out.printf("%d é um número primo.\n", numero);
        } else {
            // Se não for primo, imprime que não é um número primo
            System.out.printf("%d não é um número primo.\n", numero);
        }
    }
}
