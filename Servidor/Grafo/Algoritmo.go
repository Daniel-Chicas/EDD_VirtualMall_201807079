package GrafoRecorrido

import "fmt"

type NodoDijkstra struct {
	Vertice string
	Final float64
	Temporal float64
	Siguiente *NodoDijkstra
	Anterior *NodoDijkstra
}

type ListaDijkstra struct {
	Cabeza *NodoDijkstra
	Cola *NodoDijkstra
}

func (L *ListaDijkstra) InsertarVertice(nuevo *NodoDijkstra){
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
}

func (L *ListaAdyacencia) Dijkstra(inicio string, final string, recorrer []Nodos) *ListaRecorrido{
	if inicio == final {
		llenando := &ListaRecorrido{}
		nuevo := &NodoRecorrido{Viene: inicio, Va: final, Costo: 0, Siguiente: nil, Anterior: nil}
		llenando.InsertarRecCabeza(nuevo)
		return llenando
	}

	linkRe := &ListaRecorrido{}
	linkDijk := ListaDijkstra{}
	if recorrer != nil {
		for i := 0; i < len(recorrer); i++ {
			if recorrer[i].Nombre == inicio {
				for j := 0; j < len(recorrer[i].Enlaces); j++ {
					nuevo := &NodoRecorrido{Viene: recorrer[i].Nombre, Va: recorrer[i].Enlaces[j].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					nuevoRe := &NodoRecorrido{Viene: recorrer[i].Enlaces[j].Nombre, Va: recorrer[i].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					linkRe.InsertarRec(nuevo)
					linkRe.InsertarRec(nuevoRe)
				}
			}
		}

		for i := 0; i < len(recorrer); i++ {
			if recorrer[i].Nombre != inicio && recorrer[i].Nombre != final {
				for j := 0; j < len(recorrer[i].Enlaces); j++ {
					nuevo := &NodoRecorrido{Viene: recorrer[i].Nombre, Va: recorrer[i].Enlaces[j].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					nuevoRe := &NodoRecorrido{Viene: recorrer[i].Enlaces[j].Nombre, Va: recorrer[i].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					linkRe.InsertarRec(nuevo)
					linkRe.InsertarRec(nuevoRe)
				}
			}
		}

		for i := 0; i < len(recorrer); i++ {
			if recorrer[i].Nombre == final {
				for j := 0; j < len(recorrer[i].Enlaces); j++ {
					nuevo := &NodoRecorrido{Viene: recorrer[i].Nombre, Va: recorrer[i].Enlaces[j].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					nuevoRe := &NodoRecorrido{Viene: recorrer[i].Enlaces[j].Nombre, Va: recorrer[i].Nombre, Costo: recorrer[i].Enlaces[j].Distancia, Siguiente: nil, Anterior: nil}
					linkRe.InsertarRec(nuevo)
					linkRe.InsertarRec(nuevoRe)
				}
			}
		}
	}
	linkDijk.InsertarVertice(&NodoDijkstra{Vertice: inicio, Final: 0, Temporal: 0, Siguiente: nil, Anterior: nil})
	imp := linkRe.Cabeza
	for imp != nil{
		existe := false
		impd := linkDijk.Cabeza
		for impd != nil{
			if impd.Vertice == imp.Viene{
				existe = true
			}
			impd = impd.Siguiente
		}
		if existe == false {
			linkDijk.InsertarVertice(&NodoDijkstra{Vertice: imp.Viene,Final: 0.09, Temporal: 0.09, Siguiente: nil, Anterior: nil})
		}
		imp = imp.Siguiente
	}

	llenando := &ListaRecorrido{}
	for i := linkDijk.Cabeza; i != nil ; i = i.Siguiente {
		if i.Vertice == final {
			for i.Final == 0.09{
			 	CaminoDijkstra := caminoCorto(inicio, &linkDijk, linkRe, 0)
			 	caminoFinal(final, CaminoDijkstra, linkRe, llenando)
			}
		}
	}

	return llenando
}

func caminoCorto (inicio string, lista *ListaDijkstra, recorrido *ListaRecorrido, cuenta float64) *ListaDijkstra{
	for i := recorrido.Cabeza; i != nil ; i = i.Siguiente {
		if i.Viene == inicio {
			for e := lista.Cabeza; e != nil ; e = e.Siguiente {
				if i.Va == e.Vertice {
					nuevo := cuenta+i.Costo
					if e.Temporal==0.09 || e.Temporal>nuevo {
						e.Temporal = nuevo
					}
				}
			}
		}
	}
	nuevo := &NodoDijkstra{}
	for e := lista.Cabeza; e != nil ; e = e.Siguiente {
		if e.Vertice != inicio && e.Final == 0.09 && e.Temporal!=0.09{
			if nuevo.Temporal == 0 {
				nuevo = e
			}else{
				if e.Temporal < nuevo.Temporal && nuevo.Final==0.09 {
					nuevo = &NodoDijkstra{e.Vertice, e.Temporal, e.Temporal, nil, nil}
				}
			}
		}
	}
	for e := lista.Cabeza; e != nil ; e = e.Siguiente {
		if e.Vertice == nuevo.Vertice {
			e.Final = nuevo.Temporal
		}
	}
	if nuevo.Vertice == "" {
		return lista
	}else{
		caminoCorto(nuevo.Vertice, lista, recorrido, nuevo.Final)
	}
	return lista
}

func caminoFinal(final string, lista *ListaDijkstra, recorrido *ListaRecorrido, llenando *ListaRecorrido){
	for e := lista.Cabeza; e != nil ; e = e.Siguiente {
		if e.Vertice == final {
			for i := recorrido.Cabeza; i != nil ; i = i.Siguiente {
				if i.Va == final {
					prueba := e.Final-i.Costo
					for j := lista.Cabeza; j != nil ; j = j.Siguiente {
						if i.Viene == j.Vertice {
							if j.Final == prueba {
								nuevo := &NodoRecorrido{Viene: j.Vertice, Va: i.Va, Costo: i.Costo, Siguiente: nil, Anterior: nil}
								llenando.InsertarRecCabeza(nuevo)
								caminoFinal(j.Vertice, lista, recorrido, llenando)
							}
						}
					}
				}
			}
		}
	}
}

func (L *ListaAdyacencia) DFS(){
	aux := &Lista{}
	for e := L.Lista.Cabeza; e != nil ; e=e.Siguiente {
		L.DFS2(aux, e.Vertice)
	}
	for i := aux.Cabeza; i != nil ; i = i.Siguiente {
		fmt.Println(i.Vertice.Departamento)
	}
	fmt.Println("---------------------------------------------------------------------------")
}

func (L *ListaAdyacencia) DFS2(aux *Lista, actual *Vertice){
	if Contiene(aux, actual) == false {
		aux.InsertarVertice(&NodoLista{Vertice: actual, Siguiente: nil, Anterior: nil})
	}else{
		return
	}
	for e := actual.Adyacentes.Cabeza; e != nil; e=e.Siguiente{
		L.DFS2(aux, e.Vertice)
	}
}

