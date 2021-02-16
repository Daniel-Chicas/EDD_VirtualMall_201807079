package TiendaEspecifica

import (
	"../Listas"
	"strconv"
	"strings"
)

type General struct{
	Departamento string `json:"Departamento"`
	Nombre string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`
}

type Buscar struct{
	Depa string
	NombreB string
	Cal int
}

func (B * Buscar) Buscar(vector []Listas.NodoArray, Indices []string, Departamentos []string) string{
	var cadena string
	indice := strings.Split(B.NombreB, "")
	posFila := Posicion(Indices, indice[0])
	posColumna := Posicion(Departamentos, B.Depa)
	Primero := posFila-0
	Segundo := Primero * len(Departamentos) + posColumna
	Tercero := Segundo*5+(B.Cal-1)
	imp := vector[Tercero].ListGA.Cabeza
	for imp != nil {
		if imp.NombreTienda == B.NombreB {
			cadena = imp.NombreTienda+"&"+imp.Descripcion+"&"+imp.Contacto+"&"+strconv.Itoa(imp.Calificacion)
		}
		imp = imp.Siguiente
	}
	return cadena
}

func Posicion(arreglo []string, busqueda string) int {
	for indice, valor := range arreglo {
		if valor == busqueda {
			return indice
		}
	}
	return -1
}