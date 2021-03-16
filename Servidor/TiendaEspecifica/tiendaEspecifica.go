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

type GeneralEliminar struct{
	Categoria string `json:"Departamento"`
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
	Tercero := B.posicionTercero(Indices, Departamentos)
	imp := vector[Tercero].ListGA.Cabeza
	for imp != nil {
		if imp.NombreTienda == B.NombreB {
			cadena = imp.NombreTienda+"&"+imp.Descripcion+"&"+imp.Contacto+"&"+strconv.Itoa(imp.Calificacion)+"&"+imp.Logo
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

func (B *Buscar) BusquedaPosicion (vector []Listas.NodoArray, posicion int) string{
	var cadena string
	if posicion >= len(vector) {
		cadena = "Posición inválida."
	}else if vector != nil {
		imp := vector[posicion].ListGA.Cabeza
		if imp == nil{
			cadena = "No hay tienda."
		}else{
			for imp != nil {
				cadena = cadena+"%"+imp.NombreTienda+"&"+imp.Descripcion+"&"+imp.Contacto+"&"+strconv.Itoa(imp.Calificacion)+"&"+imp.Logo
				imp = imp.Siguiente
			}
		}
	}
	return cadena
}

func (BE *Buscar) Eliminar(vector []Listas.NodoArray, Indices []string, Departamentos []string) []Listas.NodoArray{
	Tercero := BE.posicionTercero(Indices, Departamentos)
	lista := vector[Tercero].ListGA
	impC := vector[Tercero].ListGA.Cabeza
	for impC != nil {
		if impC.NombreTienda == BE.NombreB {
			if impC == lista.Cabeza {
				lista.Cabeza = impC.Siguiente
				if impC.Siguiente != nil {
					impC.Siguiente.Anterior = nil
				}
			}else if impC.Siguiente != nil {
				impC.Anterior.Siguiente = impC.Siguiente
				impC.Siguiente.Anterior = impC.Anterior
			}else{
				impC.Anterior.Siguiente = nil
			}
			if impC == lista.Cola {
				lista.Cola = impC.Anterior
				if impC.Anterior != nil {
					impC.Anterior.Siguiente = nil
				}
			}
			impC = nil
		}else {
			impC = impC.Siguiente
		}
	}
	vector[Tercero].ListGA = lista
	return vector
}

func (B *Buscar) posicionTercero(Indices []string, Departamentos []string) int{
	indice := strings.Split(B.NombreB, "")
	posFila := Posicion(Indices, indice[0])
	posColumna := Posicion(Departamentos, B.Depa)
	Primero := posFila-0
	Segundo := Primero * len(Departamentos) + posColumna
	Tercero := Segundo*5+(B.Cal-1)
	return Tercero
}