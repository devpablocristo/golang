El patrón de diseño Builder se utiliza para separar la construcción de un objeto complejo de su representación, de manera que el mismo proceso de construcción pueda crear diferentes representaciones. Este patrón es especialmente útil cuando un objeto debe ser creado con muchas opciones posibles y no todas ellas son necesarias en cada instancia.

### Aspectos Fundamentales del Patrón Builder Explicados con el Código

#### Separación de la Construcción de su Representación
- **Cómo se Implementa**: El patrón permite construir diferentes "representaciones" (es decir, estados o configuraciones) del objeto `Person` sin cambiar el proceso subyacente. En el código, esto se logra mediante el uso de `PersonJobBuilder` y `PersonAddressBuilder`, que manejan diferentes aspectos de `Person`.
- **Beneficio**: Esta separación facilita la construcción de variaciones complejas del objeto `Person` utilizando el mismo constructor principal (`PersonBuilder`).

#### Encapsulamiento de la Construcción
- **Cómo se Implementa**: Los detalles de cómo se ensambla el objeto `Person` están ocultos dentro de los builders específicos. Por ejemplo, el cliente (`main`) no necesita conocer cómo `PersonJobBuilder` establece el campo `CompanyName`.
- **Beneficio**: El usuario del builder solo necesita interactuar con la interfaz pública del builder, sin preocuparse por los detalles internos de la construcción del objeto.

#### Flexibilidad en la Construcción
- **Cómo se Implementa**: Al permitir configurar diferentes aspectos del objeto `Person`

 de forma independiente (trabajo y dirección), el código facilita modificaciones en el objeto construido sin necesidad de alterar el código existente.
- **Beneficio**: Se pueden hacer cambios en la estructura o los atributos de `Person` con mínimas a mínimas modificaciones en el código que utiliza el patrón Builder.

#### Fluent Interface ("Method Chaining")
- **Cómo se Implementa**: Los métodos de `PersonJobBuilder` y `PersonAddressBuilder` retornan referencias a sí mismos, lo que permite encadenar llamadas de manera legible y expresiva, como se muestra en la función `main`.
- **Beneficio**: Esta característica mejora la legibilidad y simplifica el código cliente, haciendo el proceso de construcción del objeto más intuitivo y fluido.

En resumen, este código ilustra cómo el patrón Builder no solo simplifica la creación de objetos complejos, sino que también proporciona una estructura flexible, mantenible y fácil de usar para configurar objetos con múltiples atributos y representaciones.