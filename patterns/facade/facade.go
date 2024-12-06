
// Patrón Facade
// El patrón Facade proporciona una interfaz simplificada a una biblioteca, un marco o cualquier otro grupo complejo de clases.

// Supongamos que tiene varios componentes como Disk, Memory y CPU en un sistema de computadora.

// Definimos la estructura Disk, que representa un componente del sistema de computadora.

// En este caso, Computer es una fachada que proporciona una interfaz simplificada (Start()) para iniciar varios subsistemas (Disk, Memory, CPU).

// El patrón Facade es un patrón de diseño estructural. Los patrones estructurales se ocupan de la composición de clases y objetos para formar estructuras más grandes. El patrón Facade se utiliza para simplificar las interfaces y ocultar la complejidad del sistema subyacente, proporcionando una interfaz unificada que facilita la interacción con el sistema.
type Disk struct{}

// Definimos un método Start para la estructura Disk, que imprime una línea en la consola.
func (d *Disk) Start() {
	fmt.Println("Iniciando disco")
}

// Definimos la estructura Memory, que representa otro componente del sistema de computadora.
type Memory struct{}

// Definimos un método Load para la estructura Memory, que imprime una línea en la consola.
func (m *Memory) Load() {
	fmt.Println("Cargando memoria")
}

// Definimos la estructura CPU, que representa otro componente del sistema de computadora.
type CPU struct{}

// Definimos un método Execute para la estructura CPU, que imprime una línea en la consola.
func (c *CPU) Execute() {
	fmt.Println("Ejecutando CPU")
}

// Definimos la estructura Computer, que representa un sistema de computadora que incluye un Disk, Memory y CPU.
type Computer struct {
	cpu    *CPU
	memory *Memory
	disk   *Disk
}

// Creamos un constructor para la estructura Computer que inicializa un sistema de computadora con Disk, Memory y CPU.
func NewComputer() *Computer {
	return &Computer{
		cpu:    &CPU{},
		memory: &Memory{},
		disk:   &Disk{},
	}
}

// Definimos un método Start para la estructura Computer, que inicia los componentes Disk, Memory y CPU en ese orden.
func (c *Computer) Start() {
	c.disk.Start()
	c.memory.Load()
	c.cpu.Execute()
}

// La función main crea una nueva instancia de Computer y comienza a iniciar sus componentes.
func main() {
	computer := NewComputer()
	computer.Start()
	// Iniciando disco
	// Cargando memoria
	// Ejecutando CPU
}

// Ejemplo:
// imagina que estás creando una aplicación de comercio electrónico.

// Esta aplicación necesita interactuar con varios sistemas diferentes para realizar una compra, como el sistema de inventario (para verificar la disponibilidad del producto), el sistema de facturación (para crear una factura), el sistema de pago (para procesar el pago del cliente) y el sistema de envío (para organizar el envío del producto al cliente).

// Cada uno de estos sistemas tiene su propia interfaz y lógica de negocio, lo que puede hacer que la realización de una compra sea una operación compleja.

// Aquí es donde puede ser útil el patrón Facade. Podrías crear una clase CompraFacade que proporciona un método simple realizarCompra. Internamente, este método se encargaría de interactuar con todos los sistemas necesarios (inventario, facturación, pago y envío) en el orden correcto y manejando los posibles errores.

// De esta manera, el resto de tu aplicación de comercio electrónico no necesita conocer los detalles de cómo interactuar con estos sistemas individuales. En su lugar, simplemente puede llamar al método realizarCompra de la clase CompraFacade, lo que simplifica mucho el código y lo hace más fácil de entender y mantener.
