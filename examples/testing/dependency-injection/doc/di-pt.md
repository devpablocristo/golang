### Resumo

- **O que estamos fazendo**: Injetando funções nas nossas estruturas para simular comportamentos específicos durante os testes.
- **Onde fazemos isso**: Em `MonitorUsecase`, adicionando o campo `PlatformOrBrandFunc`.
- **Por que fazemos isso**: Para poder controlar e simular diferentes cenários sem depender das implementações reais das nossas dependências, tornando nossos testes mais previsíveis e rápidos.
- **Como prevenimos erros de referência a ponteiros nil**: Verificando se os ponteiros são nil antes de acessar seus campos ou métodos.

### Contexto e Objetivo

Quando desenvolvemos aplicações em Go, queremos garantir que nosso código seja fácil de testar. Os testes unitários nos permitem verificar se cada parte do código funciona corretamente de maneira isolada. Para isso, precisamos de uma maneira de controlar as dependências externas durante os testes.

### Injeção de Dependências

A injeção de dependências é uma técnica que permite fornecer as dependências de um componente de fora do próprio componente. Isso torna o código mais modular e fácil de testar. Nesse contexto, utilizamos funções ou interfaces injetáveis (como `PlatformOrBrandFunc`) para facilitar os testes unitários.

### Caso Específico: MonitorUsecase

Você tem uma estrutura `MonitorUsecase` que é responsável por criar monitores. Esta estrutura depende de `MonitorRepositoryPort` e `MonitorDatadogPort`, que interagem com um banco de dados e um serviço do Datadog, respectivamente.

#### Problema

Queremos testar o método `CreateMonitor` de `MonitorUsecase` sem interagir realmente com o banco de dados ou o serviço do Datadog. Precisamos simular essas interações.

#### Solução: Uso de Funções Injetáveis

Para conseguir isso, usamos funções injetáveis. Neste caso, adicionamos uma função injetável chamada `PlatformOrBrandFunc` em `MonitorUsecase`. Esta função determina a query com base no produto do monitor (`Platform` ou `Brand`).

### Implementação Passo a Passo

1. **Definição da Estrutura**

   Definimos a estrutura `MonitorUsecase` com as dependências necessárias e a função injetável.

   ```go
   type MonitorUsecase struct {
       repository          MonitorRepositoryPort
       datadog             MonitorDatadogPort
       PlatformOrBrandFunc func(ctx context.Context, m *Monitor, query []string) ([]string, error)
   }
   ```

2. **Construtor**

   Criamos um construtor para inicializar `MonitorUsecase` sem a função injetável por padrão.

   ```go
   func NewMonitorUsecase(repo MonitorRepositoryPort, datadog MonitorDatadogPort) MonitorUsecasePort {
       return &MonitorUsecase{
           repository: repo,
           datadog:    datadog,
       }
   }
   ```

3. **Método `CreateMonitor`**

   Implementamos o método `CreateMonitor`, utilizando a função injetável se estiver definida e verificando ponteiros nulos.

   ```go
   func (u *MonitorUsecase) CreateMonitor(ctx context.Context, m *Monitor) error {
       if m == nil {
           return errors.New("monitor cannot be nil")
       }

       exists, _ := u.repository.CheckMonitorExists(m)
       if exists {
           return errors.New("monitor already exists")
       }

       query, err := u.PlatformOrBrand(ctx, m, nil)
       if err != nil {
           return fmt.Errorf("error checking platform or brand: %w", err)
       }

       for _, q := range query {
           r, err := u.datadog.CreateMonitor(ctx, q)
           if err != nil {
               return fmt.Errorf("error creating datadog monitor: %w", err)
           }

           m.DatadogID = int(r.Body[0]) // simplificado para a demo

           err = u.repository.CreateMonitor(m)
           if err != nil {
               return fmt.Errorf("error saving monitor to db: %w", err)
           }
       }
       return nil
   }
   ```

4. **Método `PlatformOrBrand`**

   Implementamos o método `PlatformOrBrand`, que utiliza `PlatformOrBrandFunc` se estiver definida e verificando ponteiros nulos.

   ```go
   func (u *MonitorUsecase) PlatformOrBrand(ctx context.Context, m *Monitor, query []string) ([]string, error) {
       if m == nil {
           return nil, errors.New("monitor cannot be nil")
       }

       if u.PlatformOrBrandFunc != nil {
           return u.PlatformOrBrandFunc(ctx, m, query)
       }

       if m.Product == "Platform" {
           query = append(query, "platform_query")
       } else if m.Product == "Brand" {
           query = append(query, "brand_query")
       } else {
           return nil, errors.New("unknown product")
       }
       return query, nil
   }
   ```

### Testes

Agora que temos nossa estrutura preparada, podemos escrever testes unitários para `CreateMonitor`.

1. **Preparar Mocks**

   Usamos uma biblioteca como `gomock` para criar mocks das dependências `MonitorRepositoryPort` e `MonitorDatadogPort`.

   ```go
   package test

   import (
       "context"
       "testing"

       "github.com/golang/mock/gomock"
       "github.com/pkg/errors"
       "github.com/stretchr/testify/assert"

       "demo-project/internal/monitor"
       mock_monitor "demo-project/internal/monitor/mock"
   )
   ```

2. **Escrever Testes**

   Escrevemos diferentes cenários de testes injetando uma função `PlatformOrBrandFunc` diferente conforme o caso.

   ```go
   func TestCreateMonitor(t *testing.T) {
       type args struct {
           ctx context.Context
           m   *monitor.Monitor
       }
       tests := []struct {
           name            string
           repo            func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort
           datadog         func(ctrl *gomock.Controller) monitor.MonitorDatadogPort
           platformOrBrand func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error)
           args            args
           wantErr         bool
           err             error
       }{
           {
               name: "monitor already exists",
               repo: func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort {
                   mockRepo := mock_monitor.NewMockMonitorRepositoryPort(ctrl)
                   mockRepo.EXPECT().CheckMonitorExists(gomock.Any()).Return(true, nil)
                   return mockRepo
               },
               datadog: func(ctrl *gomock.Controller) monitor.MonitorDatadogPort {
                   return nil
               },
               platformOrBrand: func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error) {
                   return nil, errors.New("unknown product")
               },
               args: args{
                   ctx: context.Background(),
                   m: &monitor.Monitor{
                       ID:          123,
                       Link:        "https://example.com",
                       DatadogID:   0,
                       DowntimeID:  "456",
                       Product:     "Platform",
                       BrandName:   "ExampleBrand",
                       Site:        "MLA",
                       Description: "Monitoring service for ExampleBrand",
                   },
               },
               wantErr: true,
               err:     errors.New("monitor already exists"),
           },
           {
               name: "error checking platform or brand",
               repo: func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort {
                   mockRepo := mock_monitor.NewMockMonitorRepositoryPort(ctrl)
                   mockRepo.EXPECT().CheckMonitorExists(gomock.Any()).Return(false, nil)
                   return mockRepo
               },
               datadog: func(ctrl *gomock.Controller) monitor.MonitorDatadogPort {
                   mockDatadog := mock_monitor.NewMockMonitorDatadogPort(ctrl)
                   // Não esperamos que CreateMonitor seja chamado neste caso, pois deve falhar antes.
                   mockDatadog.EXPECT().CreateMonitor(gomock.Any(), gomock.Any()).Times(0)
                   return mockDatadog
               },
               platformOrBrand: func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error) {
                   return nil, errors.New("unknown product")
               },
               args: args{
                   ctx: context.Background(),
                   m: &monitor.Monitor{
                       ID:          123,
                       Link:        "https://example.com",
                       DatadogID:   0,
                       DowntimeID:  "456",
                       Product:     "Platform",
                       BrandName:   "ExampleBrand",
                       Site:        "MLA",
                       Description: "Monitoring service for ExampleBrand",
                   },
               },
               wantErr: true,
               err:     errors.New("error checking platform or brand: unknown product"),
           },
       }

       for _, tt := vai tt.name, func(t *testing.T) {
               ctrl := gomock.NewController(t)
               defer ctrl.Finish()

               repo := tt.repo(ctrl)
               datadog := tt.datadog(ctrl)

               usecase := monitor.NewMonitorUsecase(repo, datadog)

               // Mock de PlatformOrBrand a nível de teste
               usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand

               err := usecase.CreateMonitor(tt.args.ctx, tt.args.m)

               if (err != nil) != tt.wantErr {
                   t.Errorf("CreateMonitor() error = %v, wantErr %v", err, tt.wantErr)
               }
               if tt.wantErr {
                   assert.Error(t, err)
                   assert.Equal(t, tt.err.Error(), err.Error())
               } else {
                   assert.NoError(t, err)
               }
           })
       }
   }
   ```

### Explicação da Linha Crítica

```go
usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand
```

#### Passo a Passo

1. **Interface e Estrutura**:

   ```go
   type MonitorUsecasePort interface {
       CreateMonitor(ctx context.Context, m *Monitor) error
       PlatformOrBrand(ctx context.Context, m *Monitor

, query []string) ([]string, error)
   }

   type MonitorUsecase struct {
       repository          MonitorRepositoryPort
       datadog             MonitorDatadogPort
       PlatformOrBrandFunc func(ctx context.Context, m *Monitor, query []string) ([]string, error)
   }
   ```

   `NewMonitorUsecase` retorna um tipo `MonitorUsecasePort`, que é uma interface. A implementação concreta é `MonitorUsecase`.

2. **Atribuição de Função Injetável**:

   Nos nossos testes, queremos injetar uma função específica no campo `PlatformOrBrandFunc` de `MonitorUsecase` para controlar seu comportamento durante os testes.

3. **Conversão de Tipo (Type Assertion)**:

   ```go
   usecase.(*monitor.MonitorUsecase)
   ```

   `usecase` é do tipo `MonitorUsecasePort`, que é uma interface. Para acessar os campos específicos da estrutura `MonitorUsecase`, precisamos converter (`type assertion`) `usecase` para seu tipo concreto `*MonitorUsecase`. Isso nos permite acessar diretamente os campos e métodos específicos de `MonitorUsecase`.

4. **Atribuição de Função**:

   ```go
   usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand
   ```

   Aqui estamos atribuindo a função `tt.platformOrBrand` ao campo `PlatformOrBrandFunc` da nossa instância concreta `MonitorUsecase`. `tt.platformOrBrand` é uma função específica definida em cada caso de teste que simula o comportamento que queremos verificar.

### Prevenir Erros de Referência a Ponteiros Nil

#### O Que é um Erro de Referência a Ponteiros Nil?

Um erro de referência a ponteiros nil ocorre quando você tenta acessar um campo ou método de um ponteiro que não foi inicializado (é nil). Isso causa um pânico em tempo de execução e faz com que o programa falhe.

#### Como Preveni-lo

1. **Verificações de nil**: Sempre verifique se os ponteiros são nil antes de acessar seus campos ou métodos.
2. **Inicialização adequada**: Certifique-se de que os ponteiros sejam inicializados corretamente antes de usá-los.

No nosso código, prevenimos esses erros verificando se o monitor (`m`) é nil no início dos métodos `CreateMonitor` e `PlatformOrBrand`.

```go
if m == nil {
    return errors.New("monitor cannot be nil")
}
```

