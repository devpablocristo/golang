// Patrón State
// El patrón de diseño State permite que un objeto cambie su comportamiento cuando su estado interno cambia. Esto parece como si el objeto cambia de clase.

// Por ejemplo, imagine que está trabajando en una aplicación de música y tiene un objeto Player que tiene varios estados, como Playing, Paused y Stopped. En lugar de tener un montón de condicionales, puede usar el patrón State.

// Definimos una interfaz llamada 'State' que especifica los comportamientos que deben tener todos los estados.

// Por otro lado, el patrón State es un patrón de diseño de comportamiento. Los patrones de comportamiento se centran en los algoritmos y la asignación de responsabilidades entre objetos. El patrón State se utiliza para gestionar cambios en el comportamiento de un objeto cuando su estado interno cambia. En lugar de utilizar declaraciones condicionales para cambiar el comportamiento, el patrón State encapsula los comportamientos dentro de objetos de estado y delega la ejecución del comportamiento a los objetos de estado.

type State interface {
	Play()
	Stop()
}

// Creamos una estructura 'PlayingState' que representa el estado de reproducción.
type PlayingState struct{}

// Implementamos el método 'Play()' para 'PlayingState'. En este estado, si intentamos reproducir música, se muestra un mensaje diciendo que ya estamos reproduciendo música.
func (p *PlayingState) Play() {
	fmt.Println("Ya estamos reproduciendo música")
}

// Implementamos el método 'Stop()' para 'PlayingState'. En este estado, si intentamos detener la música, se muestra un mensaje diciendo que estamos deteniendo la música.
func (p *PlayingState) Stop() {
	fmt.Println("Deteniendo la música")
}

// Creamos una estructura 'StoppedState' que representa el estado de detención.
type StoppedState struct{}

// Implementamos el método 'Play()' para 'StoppedState'. En este estado, si intentamos reproducir música, se muestra un mensaje diciendo que estamos reproduciendo música.
func (s *StoppedState) Play() {
	fmt.Println("Reproduciendo música")
}

// Implementamos el método 'Stop()' para 'StoppedState'. En este estado, si intentamos detener la música, se muestra un mensaje diciendo que ya estamos detenidos.
func (s *StoppedState) Stop() {
	fmt.Println("Ya estamos detenidos")
}

// Creamos una estructura 'Player' que mantiene el estado actual en que se encuentra el reproductor de música.
type Player struct {
	state State
}

// Implementamos el método 'Play()' para 'Player' que simplemente invoca el método 'Play()' del estado actual.
func (p *Player) Play() {
	p.state.Play()
}

// Implementamos el método 'Stop()' para 'Player' que simplemente invoca el método 'Stop()' del estado actual.
func (p *Player) Stop() {
	p.state.Stop()
}

// Creamos una función para crear un nuevo reproductor de música. Inicialmente, el reproductor está en estado de detención ('StoppedState').
func NewPlayer() *Player {
	return &Player{
		state: &StoppedState{},
	}
}

// La función principal donde se crea un nuevo reproductor y se prueba la funcionalidad.
func main() {
	// Creamos un nuevo reproductor.
	player := NewPlayer()

	// Intentamos reproducir música. Como el reproductor está en estado de detención, el mensaje 'Reproduciendo música' se imprime en la consola.
	player.Play() // Reproduciendo música

	// Cambiamos el estado del reproductor a 'PlayingState'.
	player.state = &PlayingState{}

	// Intentamos reproducir música de nuevo. Como el reproductor ahora está en estado de reproducción, el mensaje 'Ya estamos reproduciendo música' se imprime en la consola.
	player.Play() // Ya estamos reproduciendo música

	// Intentamos detener la música. Como el reproductor está en estado de reproducción, el mensaje 'Deteniendo la música' se imprime en la consola.
	player.Stop() // Deteniendo la música
}

func main() {
	player := NewPlayer()

	player.Play() // Reproduciendo música
	player.state = &PlayingState{}
	player.Play() // Ya estamos reproduciendo música
	player.Stop() // Deteniendo la música
}

// Ejemplo
// Una aplicación de semáforo.

// Un semáforo típico tiene tres estados: verde, amarillo y rojo. Cada estado tiene reglas específicas. Por ejemplo, si el semáforo está en verde, puede cambiar a amarillo; si está en amarillo, puede cambiar a rojo; y si está en rojo, puede cambiar a verde.

// Podrías representar cada uno de estos estados como un objeto en tu aplicación, cada uno con un método cambiarEstado. El estado verde sabe cómo cambiar a amarillo, el estado amarillo sabe cómo cambiar a rojo, y el estado rojo sabe cómo cambiar a verde.

// Además, podrías tener una clase Semáforo que mantiene un estado actual y tiene un método cambiarEstado. Este método simplemente delega la operación al estado actual, que sabe cómo hacer el cambio apropiado.

// De esta manera, tu clase Semáforo no necesita tener una lógica compleja para decidir qué hacer en cada estado. En su lugar, puede delegar esa decisión a los objetos de estado, que conocen el comportamiento apropiado para su estado específico. Esto hace que el código sea más fácil de entender y mantener, y más flexible si necesitas agregar nuevos estados en el futuro.