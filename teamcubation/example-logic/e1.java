import java.util.Scanner;

public class VerificadorPrimo {

    // Função para verificar se o número é primo
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

        // Se não houve divisores encontrados, o número é primo
        return true;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        // Solicitar ao usuário que insira um número
        System.out.print("Digite um número: ");
        int numero = scanner.nextInt();

        // Verificar se o número é primo usando a função ehPrimo
        if (ehPrimo(numero)) {
            System.out.printf("%d é um número primo.\n", numero);
        } else {
            System.out.printf("%d não é um número primo.\n", numero);
        }

        scanner.close();
    }
}
