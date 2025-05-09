package main

// 12 elfos
// 1 santa claus
// 9 renos

// 1. **Santa Claus:**
//    - Comienza en estado dormido y solo se despierta cuando:
//      - **Tres elfos** tienen problemas al fabricar juguetes y necesitan ayuda.
//      - **Todos los renos** han llegado al taller.
//    - Una vez despierto, Santa resuelve los problemas de los elfos o prepara los renos.


var (
	elvesTotal := 12
	santaclausTotal := 1
	reindeersTotal := 9
}

type status string

const (
	StatusSleeping   Status = "sleeping"
	StatusAwakeInactive Status = "awake"
)

func santaclauseInitialState(santa *santaClause) {
	santa.state = sleeping
}

type santaClause struct {
	status Status
}

func main() {

	var santa santaClause 
	
	


}	
