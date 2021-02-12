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
	tienda []TiendasN
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
	listM ListaMA
	listE ListaEA
	listMB ListaMBA
	listB ListaBA
	listR ListaRA
	Siguiente *Tiendas
	Anterior *Tiendas
}

type TiendasAr struct{
	NombreTienda string
	Descripcion string
	Contacto string
	Calificacion int
	Siguiente *TiendasAr
	Anterior *TiendasAr
}


type ListaMA struct{
	Cabeza *TiendasAr
	Cola *TiendasAr
}

type ListaEA struct{
	Cabeza *TiendasAr
	Cola *TiendasAr
}

type ListaMBA struct{
	Cabeza *TiendasAr
	Cola *TiendasAr
}

type ListaBA struct{
	Cabeza *TiendasAr
	Cola *TiendasAr
}

type ListaRA struct{
	Cabeza *TiendasAr
	Cola *TiendasAr
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

func (L *ListaMA) InsertarMA(nuevo *TiendasAr) string{
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

func (L *ListaEA) InsertarEA(nuevo *TiendasAr) string{
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

func (L *ListaMBA) InsertarMBA(nuevo *TiendasAr) string{
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

func (L *ListaBA) InsertarBA(nuevo *TiendasAr) string{
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

func (L *ListaRA) InsertarRA(nuevo *TiendasAr) string{
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
	linkM := &ListaM{}
	linkE := &ListaE{}
	linkMB := &ListaMB{}
	linkB := &ListaB{}
	linkR := &ListaR{}
	imp = L.Cabeza
	for imp != nil {
		for j := 0; j < len(indices); j++ {
			listaTiend := []TiendasN{}
			if imp.Indice == indices[j] && imp != nil {
				for k := 0; k < len(departamentos); k++ {
					if imp.Departamento.NombreDepartamento == departamentos[k] {
						tienda := imp.Departamento.Tienda
						listaTiend = append(listaTiend, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
						nuevo := NodoG{Indice: imp.Indice, NombreDepartamento: imp.Departamento.NombreDepartamento, tienda: listaTiend}
						if tienda.Calificacion == 1 {
							impr := linkR.Cabeza
							if impr != nil {
								if len(impr.tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.tienda = append(impr.tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
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
								if len(impr.tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.tienda = append(impr.tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
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
								if len(impr.tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.tienda = append(impr.tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
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
								if len(impr.tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.tienda = append(impr.tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
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
								if len(impr.tienda) != 0 {
									if imp.Indice == impr.Indice && imp.Departamento.NombreDepartamento == impr.NombreDepartamento {
										impr.tienda = append(impr.tienda, TiendasN{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion})
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
	vector := linkT.CrearArray(indices, departamentos, *linkM, *linkE, *linkMB, *linkB, *linkR)
	fmt.Println(vector)
	return ""
}

func (L *Lista) CrearArray(indices []string, departamentos []string, m ListaM, e ListaE, mb ListaMB, b ListaB, r ListaR) []NodoArray{
	CM := m.Cabeza
	CE := e.Cabeza
	CMB := mb.Cabeza
	CB := b.Cabeza
	CR := r.Cabeza
	var Vector []NodoArray
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(departamentos); j++ {
			linkMA := &ListaMA{}
			linkEA := &ListaEA{}
			linkMBA := &ListaMBA{}
			linkBA := &ListaBA{}
			linkRA := &ListaRA{}
			if CM != nil {
				if CM.Indice == indices[i] && CM.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CM.tienda); k++ {
						tienda := CM.tienda[k]
						nuevo := TiendasAr{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkMA.InsertarMA(&nuevo)
					}
					CM = CM.Siguiente
				}
			}
			if CE != nil {
				if CE.Indice == indices[i] && CE.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CE.tienda); k++ {
						tienda := CE.tienda[k]
						nuevo := TiendasAr{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkEA.InsertarEA(&nuevo)
					}
					CE = CE.Siguiente
				}
			}
			if CMB != nil {
				if CMB.Indice == indices[i] && CMB.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CMB.tienda); k++ {
						tienda := CMB.tienda[k]
						nuevo := TiendasAr{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkMBA.InsertarMBA(&nuevo)
					}
					CMB = CMB.Siguiente
				}
			}
			if CB != nil {
				if CB.Indice == indices[i] && CB.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CB.tienda); k++ {
						tienda := CB.tienda[k]
						nuevo := TiendasAr{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkBA.InsertarBA(&nuevo)
					}
					CB = CB.Siguiente
				}
			}
			if CR != nil {
				if CR.Indice == indices[i] && CR.NombreDepartamento == departamentos[j] {
					for k := 0; k < len(CR.tienda); k++ {
						tienda := CR.tienda[k]
						nuevo := TiendasAr{NombreTienda: tienda.NombreTienda, Descripcion: tienda.Descripcion, Contacto: tienda.Contacto, Calificacion: tienda.Calificacion}
						linkRA.InsertarRA(&nuevo)
					}
					CR = CR.Siguiente
				}
			}
			Vector = append(Vector, NodoArray{Indice: indices[i], Departamento: departamentos[j], listM: *linkMA, listE: *linkEA, listMB: *linkMBA, listB: *linkBA, listR: *linkRA})
		}
	}
	return Vector
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