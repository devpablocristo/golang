package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type carta struct {
	Numero int
	Palo   string
}

type mazo []carta

func nuevoMazo() mazo {
	var m mazo

	palos := []string{"Copas", "Espadas", "Oros", "Bastos"}
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, p := range palos {
		for _, n := range numeros {
			m = append(m, carta{
				Numero: n,
				Palo:   p,
			})
		}
	}
	return m
}

func (m mazo) mostrar() {
	palo := ""
	for _, v := range m {
		if palo != v.Palo {
			palo = v.Palo
			fmt.Println("----------------------------------")
		}
		fmt.Println("Palo:", v.Palo, "- Numero: ", v.Numero)
	}
}

func (m mazo) mezclar() {

	// para cread randomness hay que hacer antes todo esto
	// no es tan simple como poner random() y listo
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range m {
		newPos := r.Intn(len(m) - 1)

		m[i], m[newPos] = m[newPos], m[i]
	}

	/*rand.Shuffle(len(m), func(i, j int) {
		m[i], m[j] = m[j], m[i]
	})*/
}

func repartir(mano int, m mazo) (mazo, mazo) {
	return m[:mano], m[mano+1:]
}

/*
func (m mazo) aByteSlice() []byte {
	var bs []byte
	for _, c := range m {
		bs = append(bs, []byte(strconv.Itoa(c.Numero))...)
		bs = append(bs, ","...)
		bs = append(bs, []byte(c.Palo)...)
		bs = append(bs, ","...)
	}
	return bs
}
*/

/*
func (m mazo) guardarEnArchivoTxt(nombreArchivo string) error {
	var err error
	bs := m.aByteSlice()
	err = ioutil.WriteFile(nombreArchivo, bs, 0666)
	return err
}
*/

/*
esta solución poco eficiente, MEJORAR.
*/
func (m mazo) guardarEnArchivoCsv(nombreArchivo string) error {
	archivoCsv, err := os.Create(nombreArchivo + ".csv")
	csvWriter := csv.NewWriter(archivoCsv)

	datos := []string{}
	datos = append(datos, "Numeros")
	datos = append(datos, "Palos")
	_ = csvWriter.Write(datos)

	for _, c := range m {
		datos = nil
		datos = append(datos, strconv.Itoa(c.Numero))
		datos = append(datos, c.Palo)
		_ = csvWriter.Write(datos)
	}
	csvWriter.Flush()

	return err
}

/*
func (m mazo) leerArchivoTxt(nombreArchivo string) error {
	var err error
	contenido, err := ioutil.ReadFile(nombreArchivo)
	fmt.Printf("\n\nContenido del archivo:\n%s", contenido)
	return err
}
*/

func (m mazo) leerArchivoCsv(nombreArchivo string) error {
	var err error
	contenido, err := ioutil.ReadFile(nombreArchivo + ".csv")
	fmt.Printf("\n\nContenido del archivo:\n%s", contenido)
	return err
}

/*
func nuevoMazoDesdeArchivoTxt(nombreArchivo string) (mazo, error) {
	var m mazo
	var err error
	contenido, err := ioutil.ReadFile(nombreArchivo)
	ss := strings.Split(string(contenido), ",")

	var ssAux []string
	var cartaAux carta
	for i, c := range ss {
		ssAux = strings.Split(c, ",")
		cartaAux.Numero, _ = strconv.Atoi(string(ssAux[i])) // falta manejo error
		i++
		cartaAux.Palo = ssAux[i]
		m = append(m, cartaAux)
	}
	return m, err
}
*/

func nuevoMazoDesdeArchivoCsv(nombreArchivo string) mazo {
	var m mazo

	archivoCsv, err := os.Open(nombreArchivo + ".csv")
	if err != nil {
		fmt.Println("Error en 'nuevoMazoDesdeArchivoCsv':", err)
		os.Exit(1)
	}

	// supongo qu es mas eficiente se pone aqui, investigar
	defer archivoCsv.Close()

	lineasCsv, err := csv.NewReader(archivoCsv).ReadAll()
	if err != nil {
		fmt.Println("Error en 'nuevoMazoDesdeArchivoCsv':", err)
		os.Exit(1)

	}

	var cartaAux carta
	for _, line := range lineasCsv {
		cartaAux.Numero, _ = strconv.Atoi(line[0]) // falta manejo error
		cartaAux.Palo = line[1]
		m = append(m, cartaAux)
	}

	//archivoCsv.Close()

	// eliminar primera línea, la cabecera
	m = m[1:]
	return m
}
