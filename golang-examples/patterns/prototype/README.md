Prototipo: Un objeto parcial o completamente inicializado que copias (clonas) y del cual haces uso.

El patrón de prototipo es un patrón de diseño creacional que se utiliza para crear objetos duplicados mientras se mantiene el rendimiento en mente. Es especialmente útil en situaciones donde la creación de un objeto es más costosa (en términos de recursos y tiempo) que copiar un objeto existente.

Este patrón implica la creación de un objeto prototipo que sirve como una plantilla para crear nuevas instancias. Las nuevas instancias se crean clonando este objeto prototipo en lugar de crear nuevos objetos desde cero. Esto permite una inicialización más eficiente de los nuevos objetos, especialmente si tienen configuraciones o estados iniciales complejos que serían costosos de replicar mediante métodos convencionales.

El patrón de prototipo se utiliza comúnmente en situaciones donde:

Los objetos a crear son similares en estructura pero variarán en la configuración de sus estados. Al tener un objeto prototipo preconfigurado, puedes clonarlo y modificar solo los estados necesarios para la nueva instancia, ahorrando recursos.
La creación directa de un objeto es ineficiente o complicada debido a restricciones como accesos a bases de datos extensos, llamadas a APIs, o cálculos complejos. Usar un prototipo permite omitir estos pasos repetitivos.
Necesitas ocultar la complejidad de la creación de instancias de ciertas clases a los usuarios de estas. Los usuarios pueden simplemente clonar el prototipo en lugar de lidiar con la complejidad de su inicialización.
En la programación orientada a objetos, este patrón se implementa proporcionando una interfaz que permite a los objetos ser clonados, comúnmente a través de un método clone o similar. Los lenguajes de programación como Java ofrecen soporte integrado para clonación de objetos a través de la interfaz Cloneable, mientras que otros lenguajes requieren que los desarrolladores implementen manualmente la lógica de clonación.