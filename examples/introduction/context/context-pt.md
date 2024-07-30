# O que é context?

O conceito de `context` em Go é fundamentalmente uma forma de gerenciar o estado e o controle de execução em um programa, especialmente quando se trata de operações concorrentes e em sistemas distribuídos.

## Definição e Propósito

- **Contexto como Interface**: Em Go, `context` é uma interface definida no pacote padrão `context`. Esta interface fornece métodos para controlar o cancelamento de operações, estabelecer limites de tempo para a execução de processos (prazos e timeouts) e transportar dados através da execução de uma aplicação.

- **Gerenciamento de Cancelamentos**: Uma das principais razões para usar contextos é gerenciar o cancelamento de operações de forma eficaz. Isso é especialmente útil em programas concorrentes, onde múltiplas threads de execução (goroutines) precisam ser coordenadas e possivelmente canceladas em resposta a certos eventos, como a conclusão de tarefas dependentes, erros ou interrupções do usuário.

- **Estabelecimento de Prazos e Timeouts**: Os contextos permitem especificar prazos ou timeouts após os quais uma operação deve ser automaticamente cancelada. Isso ajuda a evitar que as operações fiquem bloqueadas indefinidamente e facilita a criação de programas mais robustos e com melhor resposta a falhas.

- **Propagação de Dados**: Além de controlar o cancelamento e o tempo de vida das operações, os contextos podem carregar dados através das fronteiras das chamadas de função. Isso é útil para passar informações relevantes como identificadores de solicitação, detalhes de autenticação ou preferências de configuração ao longo de uma cadeia de processamento, sem ter que modificar a assinatura de cada função envolvida.

## Como Funciona

- **Criação e Derivação**: Um contexto é criado com uma função base como `context.Background()` ou `context.TODO()`. A partir daí, novos contextos são derivados utilizando funções como `context.WithCancel`, `context.WithDeadline`, `context.WithTimeout` e `context.WithValue`. Estes contextos derivados herdam o estado do contexto pai, mas podem ser cancelados ou ter dados anexados de forma independente.

- **Cancelamento Propagado**: Quando um contexto é cancelado, seja manualmente através de uma função de cancelamento ou automaticamente devido a um prazo ou timeout, esse cancelamento se propaga a todos os contextos derivados dele. Isso permite uma coordenação eficaz do cancelamento através de múltiplas goroutines e operações.

- **Uso através de Goroutines e API**: Os contextos são passados explicitamente de uma função para outra como o primeiro argumento (por convenção). Isso garante que o cancelamento, os prazos e os dados específicos estejam disponíveis ao longo de toda a operação, desde o início até as funções mais profundas na cadeia de chamadas.

## Context Base

A utilização mais simples do pacote `context` em Go poderia ser o uso de `context.Background()` para criar um contexto base que não faz nada de especial: não se cancela, não tem um prazo de validade e não carrega valores. Este contexto é então passado para uma função que o aceita como argumento, mas não utiliza nenhuma das características especiais do contexto, cumprindo com uma interface que espera um contexto mas sem aplicar controle de concorrência ou cancelamento.

```go
package main

import (
    "context"
    "fmt"
)

// Uma função simples que aceita um contexto e uma mensagem, mas apenas imprime a mensagem.
func printWithContext(ctx context.Context, message string) {
    fmt.Println(message)
}

func main() {
    // Criar um contexto básico
    ctx := context.Background()
    
    // Passar o contexto e uma mensagem para a função
    printWithContext(ctx, "Hello, world do context!")
}
```

Este exemplo é extremamente simplificado e na prática, o valor de `context` vem do seu uso para cancelar operações, estabelecer prazos, ou passar valores específicos através das fronteiras das chamadas de função em operações mais complexas e em execuções concorrentes.

## Propósito do context.Background

`context.Background()` tem um propósito muito específico e valioso em Go, embora à primeira vista pareça que "não se pode fazer nada com ele". É a raiz de qualquer árvore de contextos dentro de uma aplicação e serve como o

 ponto de partida para criar contextos mais específicos, ou "contextos filhos", com funcionalidades adicionais como cancelamento, prazos e armazenamento de valores específicos. Aqui detalhamos seus principais usos e por que é importante:

### Ponto de Partida Universal

- **Contexto Base:** `context.Background()` fornece um contexto base quando não há outro contexto mais específico disponível. É especialmente útil no início de uma aplicação, na função `main`, ou ao começar uma nova goroutine para a qual não existe um contexto de entrada.

### Criação de Contextos Específicos

- **Criar Contextos Filhos:** A partir de `context.Background()`, você pode derivar contextos com características adicionais. Por exemplo, você pode usar `context.WithCancel`, `context.WithDeadline` e `context.WithTimeout` para criar contextos que podem ser cancelados ou que expiram após um certo tempo. Isso é essencial para gerenciar o cancelamento de operações e timeouts em processos concorrentes.

### Neutralidade

- **Contexto Neutro:** Quando você está escrevendo uma biblioteca ou um pacote que será usado em diversos contextos, começar com `context.Background()` permite que o usuário da sua biblioteca decida como e quando introduzir contextos mais específicos, proporcionando flexibilidade.

### Casos de Uso

- **Operações de Longa Duração:** Para operações que são essenciais para a vida útil de toda a aplicação e que não se espera que sejam canceladas (por exemplo, processos de inicialização ou serviços que correm indefinidamente).

### Exemplo Prático

Suponha que você tem uma aplicação que inicia vários servidores internos (HTTP, gRPC, etc.) e serviços de fundo ao arrancar:

```go
func main() {
    ctx := context.Background()

    // Iniciar um servidor HTTP
    go startHTTPServer(ctx)

    // Iniciar um serviço de fundo
    go startBackgroundService(ctx)

    // Esperar sinais do sistema ou condições de saída
    waitForShutdown(ctx)
}
```

Neste cenário, `context.Background()` atua como o contexto raiz do qual podem ser derivados outros contextos mais específicos, conforme as necessidades de cada goroutine ou serviço. Por exemplo, você pode querer cancelar todas as operações iniciadas por `main` em caso de um sinal de desligamento; isso é facilitado passando um contexto cancelável derivado de `context.Background()` para suas funções.

Embora `context.Background()` por si só não forneça funcionalidades de cancelamento, prazos ou armazenamento de valores, seu valor reside em ser o ancestral universal de todos os contextos em uma aplicação Go. Facilita a estrutura e o manejo de operações concorrentes de maneira ordenada e previsível, permitindo a derivação de contextos mais específicos conforme necessário.

## `context.Background()` e `context.TODO()`

São duas funções do pacote `context` em Go que servem para criar contextos base. Estes contextos são pontos de partida para a derivação de novos contextos com mais funcionalidades, como cancelamento, prazos, timeouts e valores adicionais.

### `context.Background()`

`context.Background()` devolve um contexto vazio. Este contexto não se cancela, não tem valores e não tem prazo de validade. É o contexto mais "puro" e serve como o ponto de partida para criar contextos mais específicos em uma aplicação. É utilizado como o contexto raiz de uma cadeia de operações quando não há outro contexto mais apropriado que deva ser utilizado.

**Quando e por que usar `context.Background()`**:

- **Início de uma aplicação/serviço**: Ideal para usar no `main()` do seu programa ou ao iniciar goroutines a nível da aplicação, especialmente quando não há um contexto de entrada de outra operação.
- **Testes**: Frequentemente usado em testes para inicializar contextos necessários para executar operações que requerem um contexto.
- **Operações de Longa Duração**: Para operações que se espera que sejam executadas durante toda a vida útil da aplicação e não precisam ser canceladas.

### `context.TODO()`

`context.TODO()` também devolve um contexto vazio, similar a `context.Background()`. A diferença é semântica e serve como uma indicação de que o contexto deve ser definido mais tarde. `context.TODO()` é utilizado em lugares do código onde ainda não está claro que contexto usar ou se o código eventualmente precisará

 de um contexto mais específico.

**Quando, como e por que usar `context.TODO()`**:

- **Código em Desenvolvimento**: Durante o desenvolvimento inicial, quando ainda está decidindo como lidar com o cancelamento ou a passagem de valores.
- **Refatoração**: Quando está refatorando código existente para usar contextos mas ainda não decidiu como eles serão integrados em toda a aplicação.
- **Marcador de Posição**: Como um sinal para você mesmo ou para outros desenvolvedores de que o contexto precisa de revisão e provavelmente deve ser substituído por um contexto mais específico no futuro.

Os contextos em Go são imutáveis, o que significa que uma vez que um contexto é criado, ele não pode ser modificado. Toda vez que você precisa adicionar informações ou alterar o comportamento de um contexto (por exemplo, adicionando um timeout, um prazo ou valores específicos), o que você realmente faz é criar um novo contexto baseado no anterior. Este novo contexto herda as propriedades do contexto original, além das modificações ou adições que você fez.

A imutabilidade dos contextos tem várias vantagens importantes:

### Simplificação do Manejo de Concorrência

Como os contextos são imutáveis, você pode passá-los de forma segura entre goroutines sem se preocupar com problemas de concorrência, como condições de corrida. Não há risco de uma goroutine modificar o contexto de maneira que afete outras goroutines que possam estar usándo-lo.

### Segurança e Previsibilidade

Sendo imutáveis, o comportamento de um contexto é previsível. Você sabe que ele não mudará uma vez criado, o que facilita o raciocínio sobre seu código e reduz a probabilidade de efeitos colaterais inesperados.

### Encadeamento e Derivação

Quando você deriva um novo contexto de um existente, você está criando uma cadeia de contextos. Isso é útil para cancelamentos e prazos, já que cancelar um contexto pai automaticamente cancela todos os contextos derivados dele, o que é uma forma eficaz de propagar sinais de cancelamento através de sua aplicação.

### Exemplo Prático

Vejamos como isso se aplica em código. Suponha que você tem um servidor web que lida com solicitações de usuários. Para uma solicitação particular, você quer estabelecer um timeout para garantir que ela não demore demais para responder.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Criar um contexto com um timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Assegurar-se de cancelar o contexto para liberar recursos

	// Simular um trabalho que pode levar mais tempo do que permitido pelo timeout
	work(ctx)

	fmt.Fprintf(w, "Resposta enviada ao cliente")
}

func work(ctx context.Context) {
	select {
	case <-time.After(200 * time.Millisecond): // Simular trabalho
		fmt.Println("Trabalho concluído")
	case <-ctx.Done(): // O contexto foi cancelado ou o timeout expirou
		fmt.Println("Trabalho cancelado:", ctx.Err())
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

Neste exemplo, embora `context.WithTimeout` modifique o comportamento do contexto original (`context.Background()`), ele faz isso criando um novo contexto que é uma versão modificada do original. Se a operação `work` demorar muito, o contexto é cancelado (devido ao timeout), e a execução pode ser interrompida antes que a operação seja concluída, demonstrando como os contextos imutáveis podem ser usados para controlar o fluxo de operações em aplicações concorrentes.

### Exemplos de Uso

#### `context.Background()` em uma aplicação web

```go
func main() {
    ctx := context.Background()
    setupServer(ctx)
}

func setupServer(ctx context.Context) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Aqui poderíamos derivar um novo contexto para cada solicitação
        doWork(ctx)
    })
    http.ListenAndServe(":8080", nil)
}
```

#### `context.TODO()` em um código em processo de integrar contextos

```go
func fetchDataFromDB() {
    // Suponha que este código precisa ser refatorado para usar contextos
    ctx := context.TODO() //

 Marcador de que este contexto precisa ser substituído
    // Imagine uma chamada ao banco de dados aqui usando ctx
}
```

A escolha entre `context.Background()` e `context.TODO()` se reduz à intenção e clareza do código. Use `context.Background()` como um contexto raiz ou quando você está seguro de que uma operação não precisa ser cancelada. Use `context.TODO()` como um sinal de que o contexto está pendente de revisão e provavelmente precisará ser especificado com mais detalhes no futuro. Ambos os contextos servem como pontos de partida claros e limpos para a criação de contextos mais complexos em aplicações Go.