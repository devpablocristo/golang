### Resumen

- **Qué estamos haciendo**: Inyectando funciones en nuestras estructuras para simular comportamientos específicos durante las pruebas.
- **Dónde lo hacemos**: En `MonitorUsecase`, agregando el campo `PlatformOrBrandFunc`.
- **Por qué lo hacemos**: Para poder controlar y simular diferentes escenarios sin depender de las implementaciones reales de nuestras dependencias, haciendo nuestras pruebas más predecibles y rápidas.
- **Cómo prevenimos errores de referencia a punteros nil**: Verificando si los punteros son nil antes de acceder a sus campos o métodos.

### Contexto y Objetivo

Cuando desarrollamos aplicaciones en Go, queremos asegurarnos de que nuestro código sea fácil de probar. Las pruebas unitarias nos permiten verificar que cada parte del código funciona correctamente de manera aislada. Para lograr esto, necesitamos una manera de controlar las dependencias externas durante las pruebas.

### Inyección de Dependencias

La inyección de dependencias es una técnica que permite proporcionar las dependencias de un componente desde fuera del componente mismo. Esto hace que el código sea más modular y fácil de probar. En este contexto, utilizamos funciones o interfaces inyectables (como `PlatformOrBrandFunc`) para facilitar las pruebas unitarias.

### Caso Específico: MonitorUsecase

Tienes una estructura `MonitorUsecase` que se encarga de crear monitores. Esta estructura depende de `MonitorRepositoryPort` y `MonitorDatadogPort`, que interactúan con una base de datos y un servicio de Datadog, respectivamente.

#### Problema

Queremos probar el método `CreateMonitor` de `MonitorUsecase` sin interactuar realmente con la base de datos o el servicio de Datadog. Necesitamos simular estas interacciones.

#### Solución: Uso de Funciones Inyectables

Para lograr esto, usamos funciones inyectables. En este caso, añadimos una función inyectable llamada `PlatformOrBrandFunc` en `MonitorUsecase`. Esta función determina la query basada en el producto del monitor (`Platform` o `Brand`).

### Implementación Paso a Paso

1. **Definición de la Estructura**

   Definimos la estructura `MonitorUsecase` con las dependencias necesarias y la función inyectable.

   ```go
   type MonitorUsecase struct {
       repository          MonitorRepositoryPort
       datadog             MonitorDatadogPort
       PlatformOrBrandFunc func(ctx context.Context, m *Monitor, query []string) ([]string, error)
   }
   ```

2. **Constructor**

   Creamos un constructor para inicializar `MonitorUsecase` sin la función inyectable por defecto.

   ```go
   func NewMonitorUsecase(repo MonitorRepositoryPort, datadog MonitorDatadogPort) MonitorUsecasePort {
       return &MonitorUsecase{
           repository: repo,
           datadog:    datadog,
       }
   }
   ```

3. **Método `CreateMonitor`**

   Implementamos el método `CreateMonitor`, utilizando la función inyectable si está definida y verificando punteros nil.

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

           m.DatadogID = int(r.Body[0]) // simplificado para la demo

           err = u.repository.CreateMonitor(m)
           if err != nil {
               return fmt.Errorf("error saving monitor to db: %w", err)
           }
       }
       return nil
   }
   ```

4. **Método `PlatformOrBrand`**

   Implementamos el método `PlatformOrBrand`, que utiliza `PlatformOrBrandFunc` si está definida y verificando punteros nil.

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

### Pruebas

Ahora que tenemos nuestra estructura preparada, podemos escribir pruebas unitarias para `CreateMonitor`.

1. **Preparar Mocks**

   Usamos una librería como `gomock` para crear mocks de las dependencias `MonitorRepositoryPort` y `MonitorDatadogPort`.

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

2. **Escribir Pruebas**

   Escribimos diferentes escenarios de pruebas inyectando una función `PlatformOrBrandFunc` diferente según el caso.

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
                   // No esperamos que CreateMonitor sea llamado en este caso, ya que debe fallar antes.
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

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               ctrl := gomock.NewController(t)
               defer ctrl.Finish()

               repo := tt.repo(ctrl)
               datadog := tt.datadog(ctrl)

               usecase := monitor.NewMonitorUsecase(repo, datadog)

               // Mock de PlatformOrBrand a nivel de prueba
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

### Explicación de la Línea Crítica

```go
usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand
```

#### Paso a Paso

1. **Interfaz y Estructura**:

   ```go
   type MonitorUsecasePort interface {
       CreateMonitor(ctx context.Context, m *Monitor) error


       PlatformOrBrand(ctx context.Context, m *Monitor, query []string) ([]string, error)
   }

   type MonitorUsecase struct {
       repository          MonitorRepositoryPort
       datadog             MonitorDatadogPort
       PlatformOrBrandFunc func(ctx context.Context, m *Monitor, query []string) ([]string, error)
   }
   ```

   `NewMonitorUsecase` devuelve un tipo `MonitorUsecasePort` que es una interfaz. La implementación concreta es `MonitorUsecase`.

2. **Asignación de Función Inyectable**:

   En nuestras pruebas, queremos inyectar una función específica en el campo `PlatformOrBrandFunc` de `MonitorUsecase` para controlar su comportamiento durante las pruebas.

3. **Conversión de Tipo (Type Assertion)**:

   ```go
   usecase.(*monitor.MonitorUsecase)
   ```

   `usecase` es del tipo `MonitorUsecasePort`, que es una interfaz. Para acceder a los campos específicos de la estructura `MonitorUsecase`, necesitamos convertir (hacer un type assertion) `usecase` a su tipo concreto `*MonitorUsecase`. Esto nos permite acceder directamente a los campos y métodos específicos de `MonitorUsecase`.

4. **Asignación de Función**:

   ```go
   usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand
   ```

   Aquí estamos asignando la función `tt.platformOrBrand` al campo `PlatformOrBrandFunc` de nuestra instancia concreta `MonitorUsecase`. `tt.platformOrBrand` es una función específica definida en cada caso de prueba que simula el comportamiento que queremos verificar.

### Prevenir Errores de Referencia a Punteros Nil

#### ¿Qué es un Error de Referencia a Punteros Nil?

Un error de referencia a punteros nil ocurre cuando intentas acceder a un campo o método de un puntero que no ha sido inicializado (es nil). Esto causa un pánico en tiempo de ejecución y hace que el programa falle.

#### Cómo Prevenirlo

1. **Verificaciones de nil**: Siempre verifica si los punteros son nil antes de acceder a sus campos o métodos.
2. **Inicialización adecuada**: Asegúrate de que los punteros se inicialicen correctamente antes de usarlos.

En nuestro código, prevenimos estos errores verificando si el monitor (`m`) es nil al inicio de los métodos `CreateMonitor` y `PlatformOrBrand`.

```go
if m == nil {
    return errors.New("monitor cannot be nil")
}
```

