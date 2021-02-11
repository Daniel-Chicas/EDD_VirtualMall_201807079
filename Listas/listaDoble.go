package listas

import (
	"fmt"
)

type General struct{
	Inicio []Datos `json:"Datos"`
}

type Datos struct{
	Indice string `json:"Indice"`
	Departamentos []Departamento `json:"Departamentos"`
}

type Departamento struct{
	Nombre string `json:"Nombre"`
	Tiendas []Tienda `json:"Tiendas"`
}

type Tienda struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
}

type Nodo struct{
	Indice string
	Departamento Departamentos
	Siguiente *Nodo
	Anterior *Nodo
}

type Departamentos struct{
	NombreDepartamento string
	Tienda Tiendas
}

type Tiendas struct{
	NombreTienda string
	Descripcion string
	Contacto string
	Calificacion int
}

type Lista struct {
	Cabeza *Nodo
	Cola *Nodo
}

type NodoG struct{
	Indice string
	NombreDepartamento string
	NombreTienda string
	Descripcion string
	Contacto string
	Calificacion int
	Siguiente *NodoG
	Anterior *NodoG
}

type ListaM struct{
	Cabeza *NodoG
	Cola *NodoG
}

type ListaE struct{
	Cabeza *NodoG
	Cola *NodoG
}

type ListaMB struct{
	Cabeza *NodoG
	Cola *NodoG
}

type ListaB struct{
	Cabeza *NodoG
	Cola *NodoG
}

type ListaR struct{
	Cabeza *NodoG
	Cola *NodoG
}

func (L *Lista) Insertar(nuevo *Nodo) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *ListaM) InsertarM(nuevo *NodoG) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *ListaE) InsertarE(nuevo *NodoG) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *ListaMB) InsertarMB(nuevo *NodoG) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *ListaB) InsertarB(nuevo *NodoG) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *ListaR) InsertarR(nuevo *NodoG) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{

		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
	return ""
}

func (L *Lista) CrearMatriz()string{
	var indices []string
	var departamentos []string
	imp := L.Cabeza
	for imp != nil{
		existencia := existe(indices, imp.Indice)
		existes := existe(departamentos, imp.Departamento.NombreDepartamento)
		if existencia == true && existes == true{

			imp = imp.Siguiente
		}else {
			if existencia == false{
				indices = append(indices, imp.Indice)
			}
			if existes == false{
				departamentos = append(departamentos, imp.Departamento.NombreDepartamento)
			}
			imp = imp.Siguiente
		}
	}
	tam := len(indices)+ len(departamentos)
	linkM := &ListaM{}
	linkE := &ListaE{}
	linkMB := &ListaMB{}
	linkB := &ListaB{}
	linkR := &ListaR{}
		for j := 0; j < len(indices); j++ {
			imp = L.Cabeza
			for i := 0; i < tam; i++ {
			if imp.Indice == indices[j] {
				for k := 0; k < len(departamentos); k++ {
					if imp.Departamento.NombreDepartamento == departamentos[k]{
						tienda := imp.Departamento.Tienda
						nuevo := NodoG{Indice: imp.Indice, NombreDepartamento: imp.Departamento.NombreDepartamento, NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						if tienda.Calificacion == 1 {
							linkR.InsertarR(&nuevo)
						}else if tienda.Calificacion == 2 {
							linkB.InsertarB(&nuevo)
						}else if tienda.Calificacion == 3 {
							linkMB.InsertarMB(&nuevo)
						}else if tienda.Calificacion == 4 {
							linkE.InsertarE(&nuevo)
						}else if tienda.Calificacion == 5 {
							linkM.InsertarM(&nuevo)
						}
					}
				}

			}
			imp = imp.Siguiente
		}
	}
	return ""
}

func existe(arreglo []string, busqueda string) bool{
	for _, numero := range arreglo{
		if numero == busqueda{
			return true
		}
	}
	return false
}

func (L *Lista) Imprimir()string{
	imp := L.Cabeza
	for imp != nil{
		tienda := imp.Departamento.Tienda
		fmt.Println("indice: "+imp.Indice+"\t Departamento: "+imp.Departamento.NombreDepartamento+"\t Tienda: "+tienda.NombreTienda)
		imp = imp.Siguiente
	}
	return "Ya se ingresó"
}