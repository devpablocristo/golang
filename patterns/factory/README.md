El patrón de diseño Factory es utilizado para crear objetos sin necesidad de especificar la clase exacta del objeto que será creado. Este patrón es especialmente útil en situaciones donde se necesita crear diferentes tipos de objetos, que comparten una clase base común o interfaz, pero tienen diferentes características o datos iniciales.

### Aspectos Fundamentales del Patrón Factory

#### Separación de la Creación del Uso
- **Cómo se Implementa**: Ambos enfoques, funcional y estructural, permiten crear objetos `Employee` especificando ciertos detalles (como `Position` y `AnnualIncome`) sin necesidad de conocer la implementación exacta de la creación de `Employee`.
- **Beneficio**: Esto permite cambiar la implementación de cómo se crean los objetos `Employee` sin afectar el código que los utiliza.

#### Flexibilidad en la Creación de Objetos
- **Cómo se Implementa**: El enfoque estructural, en particular, permite ajustes post-creación antes de la instancia final del objeto (como modificar `AnnualIncome` de `bossFactory`).
- **Beneficio**: Proporciona una manera flexible de ajustar los detalles de los objetos antes de su creación final.

#### Encapsulamiento de la Lógica de Instanciación
- **Cómo se Implementa**: Tanto el enfoque funcional como el estructural encapsulan el proceso de creación de objetos `Employee`, ocultando los detalles de su instanciación.
- **Beneficio**: Simplifica el proceso de creación de objetos y centraliza la lógica de instanciación, facilitando su mantenimiento y modificación.

Este ejemplo muestra claramente cómo el patrón Factory puede ser utilizado para crear objetos de manera flexible y desacoplada, permitiendo la fácil extensión y mantenimiento del código.