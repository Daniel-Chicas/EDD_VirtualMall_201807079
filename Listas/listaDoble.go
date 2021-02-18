package Listas

import (
	"fmt"
)

var Vector []NodoArray
var Indices []string
var Depar []string

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
	Tienda []TiendasN
	Siguiente *NodoG
	Anterior *NodoG
}

type TiendasN struct{
	NombreTienda string
	Descripcion string
	Contacto string
	Calificacion int
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

type NodoArray struct{
	Indice string
	Departamento string
	Calificacion int
	ListGA ListaGA
}

type NodoTienda struct{
	NombreTienda string
	Descripcion string
	Contacto string
	Calificacion int
	Siguiente *NodoTienda
	Anterior *NodoTienda
}

type ListaGA struct{
	Cabeza *NodoTienda
	Cola *NodoTienda
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

func (L *ListaGA) InsertarGA(nuevo *NodoTienda) string{
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

func (L *Lista) CrearMatriz() []NodoArray{
	imp := L.Cabeza
	for imp != nil{
		existencia := existe(Indices, imp.Indice)
		existes := existe(Depar, imp.Departamento.NombreDepartamento)
		if existencia == true && existes == true{
			imp = imp.Siguiente
		}else {
			if existencia == false{
				Indices = append(Indices, imp.Indice)
			}
			if existes == false{
				Depar = append(Depar, imp.Departamento.NombreDepartamento)
			}
			imp = imp.Siguiente
		}
	}
	linkM := &ListaM{}
	linkE := &ListaE{}
	linkMB := &ListaMB{}
	linkB := &ListaB{}
	linkR := &ListaR{}
	imp = L.Cabeza
	for imp != nil {
		for j := 0; j < len(Indices); j++ {
			listaTiend := []TiendasN{}
			if imp.Indice == Indices[j] && imp != nil {
				for k := 0; k < len(Depar); k++ {
					if imp.Departamento.NombreDepartamento == Depar[k] {
						tienda := imp.Departamento.Tienda
						listaTiend = append(listaTiend, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
						nuevo := NodoG{Indice: imp.Indice, NombreDepartamento: imp.Departamento.NombreDepartamento, Tienda: listaTiend}
						if tienda.Calificacion == 1 {
							impr := linkR.Cabeza
							if impr != nil {
								if len(impr.Tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.Tienda = append(impr.Tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
									} else {
										linkR.InsertarR(&nuevo)
									}
								}
							} else {
								linkR.InsertarR(&nuevo)
							}
						} else if tienda.Calificacion == 2 {
							impr := linkB.Cabeza
							if impr != nil {
								if len(impr.Tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.Tienda = append(impr.Tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
									} else {
										linkB.InsertarB(&nuevo)
									}
								}
							} else {
								linkB.InsertarB(&nuevo)
							}
						} else if tienda.Calificacion == 3 {
							impr := linkMB.Cabeza
							if impr != nil {
								if len(impr.Tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.Tienda = append(impr.Tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
									} else {
										linkMB.InsertarMB(&nuevo)
									}
								}
							} else {
								linkMB.InsertarMB(&nuevo)
							}
						} else if tienda.Calificacion == 4 {
							impr := linkE.Cabeza
							if impr != nil {
								if len(impr.Tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.Tienda = append(impr.Tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
									} else {
										linkE.InsertarE(&nuevo)
									}
								}
							} else {
								linkE.InsertarE(&nuevo)
							}
						} else if tienda.Calificacion == 5 {
							impr := linkM.Cabeza
							if impr != nil {
								if len(impr.Tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.Tienda = append(impr.Tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
									} else {
										linkM.InsertarM(&nuevo)
									}
								}
							} else {
								linkM.InsertarM(&nuevo)
							}
						}
					}
				}
			}
		}
		imp = imp.Siguiente
	}
	linkT := &Lista{}
	linkT.CrearArray(Indices, Depar, *linkM, *linkE, *linkMB, *linkB, *linkR)
	return Vector
}

func (L *Lista) CrearArray(indices []string, departamentos []string, m ListaM, e ListaE, mb ListaMB, b ListaB, r ListaR){
	CM := m.Cabeza
	CE := e.Cabeza
	CMB := mb.Cabeza
	CB := b.Cabeza
	CR := r.Cabeza
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(departamentos); j++ {
			linkGA := &ListaGA{}
			if CR != nil {
				if CR.Indice == indices[i] && CR.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CR.Tienda); k++ {
						tienda := CR.Tienda[k]
						nuevo := NodoTienda{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkGA.InsertarGA(&nuevo)
					}
					CR = CR.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], Calificacion: 1, ListGA: *linkGA})
			linkGA = &ListaGA{}
			if CB != nil {
				if CB.Indice == indices[i] && CB.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CB.Tienda); k++ {
						tienda := CB.Tienda[k]
						nuevo := NodoTienda{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkGA.InsertarGA(&nuevo)
					}
					CB = CB.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], Calificacion: 2, ListGA: *linkGA})
			linkGA = &ListaGA{}
			if CMB != nil {
				if CMB.Indice == indices[i] && CMB.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CMB.Tienda); k++ {
						tienda := CMB.Tienda[k]
						nuevo := NodoTienda{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkGA.InsertarGA(&nuevo)
					}
					CMB = CMB.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], Calificacion: 3, ListGA: *linkGA})
			linkGA = &ListaGA{}
			if CE != nil {
				if CE.Indice == indices[i] && CE.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CE.Tienda); k++ {
						tienda := CE.Tienda[k]
						nuevo := NodoTienda{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkGA.InsertarGA(&nuevo)
					}
					CE = CE.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], Calificacion: 4, ListGA: *linkGA})
			linkGA = &ListaGA{}
			if CM != nil {
				if CM.Indice == indices[i] && CM.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CM.Tienda); k++ {
						tienda := CM.Tienda[k]
						nuevo := NodoTienda{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkGA.InsertarGA(&nuevo)
					}
					CM = CM.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], Calificacion: 5, ListGA: *linkGA})
		}
	}
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

func (L *Lista) Indi() []string{
	imp := L.Cabeza
	for imp != nil{
		existencia := existe(Indices, imp.Indice)
		if existencia == true{
			imp = imp.Siguiente
		}else {
			if existencia == false{
				Indices = append(Indices, imp.Indice)
			}
			imp = imp.Siguiente
		}
	}
	return Indices
}

func (L *Lista) Departa() []string{
	imp := L.Cabeza
	for imp != nil{
		existes := existe(Depar, imp.Departamento.NombreDepartamento)
		if existes == true{
			imp = imp.Siguiente
		}else {
			Depar = append(Depar, imp.Departamento.NombreDepartamento)
			imp = imp.Siguiente
		}
	}
	return Depar
}