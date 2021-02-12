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

type Arreglo struct{
	Indice string
	Departamento string
	Magnifica []tiendaF
	Excelente []tiendaF
	MuyBuena []tiendaF
	Buena []tiendaF
	Regular []tiendaF
}

type tiendaF struct{
	NombreT string
	DescripcionT string
	ContactoT string
	Calificacion int
	Siguiente *tiendaF
	Anterior *tiendaF
}

type ListaTienda struct{
	Cabeza *tiendaF
	Cola *tiendaF
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
	linkT := &ListaTienda{}
	linkT.CrearArray(indices, departamentos, *linkM, *linkE, *linkMB, *linkB, *linkR)
	return ""
}

func (L *ListaTienda) CrearArray(indices []string, departamentos []string, m ListaM, e ListaE, mb ListaMB, b ListaB, r ListaR){
	//var Vector []Arreglo
	CM := m.Cabeza
	CE := e.Cabeza
	CMB := mb.Cabeza
	CB := b.Cabeza
	CR := r.Cabeza
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(departamentos); j++ {
			if CM.Indice == indices[i] && CM.NombreDepartamento == departamentos[j] {
				for k := 0; k < len(CM.tienda); k++ {
					fmt.Println("SI\t"+CM.Indice+"-->"+CM.NombreDepartamento+"-->"+CM.tienda[k].NombreTienda)
				}
				CM = CM.Siguiente
			}
			if CE.Indice == indices[i] && CE.NombreDepartamento == departamentos[j] {
				for k := 0; k < len(CE.tienda); k++ {
					fmt.Println("SI\t"+CE.Indice+"-->"+CE.NombreDepartamento+"-->"+CE.tienda[k].NombreTienda)
				}
				CE = CE.Siguiente
			}
			if CMB.Indice == indices[i] && CMB.NombreDepartamento == departamentos[j] {
				for k := 0; k < len(CMB.tienda); k++ {
					fmt.Println("SI\t"+CMB.Indice+"-->"+CMB.NombreDepartamento+"-->"+CMB.tienda[k].NombreTienda)
				}
				CMB = CMB.Siguiente
			}
			if CB.Indice == indices[i] && CB.NombreDepartamento == departamentos[j] {
				for k := 0; k < len(CB.tienda); k++ {
					fmt.Println("SI\t"+CB.Indice+"-->"+CB.NombreDepartamento+"-->"+CB.tienda[k].NombreTienda)
				}
				CB = CB.Siguiente
			}
			if CR.Indice == indices[i] && CR.NombreDepartamento == departamentos[j] {
				for k := 0; k < len(CR.tienda); k++ {
					fmt.Println("SI\t"+CR.Indice+"-->"+CR.NombreDepartamento+"-->"+CR.tienda[k].NombreTienda)
				}
				CR = CR.Siguiente
			}

		}
	}
}

func (L *ListaTienda) Insertar(nuevo *tiendaF) string{
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