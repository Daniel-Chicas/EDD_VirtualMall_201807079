package GrafoRecorrido

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Archivo struct{
	PosicionInicialRobot string `json:"PosicionInicialRobot"`
	Entrega string `json:"Entrega"`
	General []Nodos `json:"Nodos"`
}

type Nodos struct{
	Nombre string `json:"Nombre"`
	Enlaces []Enlace `json:"Enlaces"`
}

type Enlace struct{
	Nombre string `json:"Nombre"`
	Distancia float64 `json:"Distancia"`
}

type Lista struct {
	Cabeza *NodoLista
	Cola *NodoLista
}

type NodoLista struct {
	Vertice *Vertice
	Siguiente *NodoLista
	Anterior *NodoLista
}

func (L *Lista) InsertarVertice(nuevo *NodoLista) string{
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

type Vertice struct{
	Departamento string
	Distancia float64
	Adyacentes *Lista
}

type ListaAdyacencia struct{
	Lista *Lista
}

func NuevoVertice (depa string, precio float64) *Vertice{
	a := &Lista{}
	return &Vertice{depa, precio, a}
}

func NuevaListaAdyacencia() *ListaAdyacencia{
	a := &Lista{}
	return &ListaAdyacencia{a}
}

func (L *ListaAdyacencia)obtenerVertice(depa string) *Vertice{
	for e:=L.Lista.Cabeza; e!=nil;e=e.Siguiente{
		if e.Vertice.Departamento == depa {
			return e.Vertice
		}
	}
	return nil
}

func (A *ListaAdyacencia) Insertar(depa string, precio float64){
	if A.obtenerVertice(depa)==nil {
		n := NuevoVertice(depa, precio)
		x := &NodoLista{n, nil, nil}
		A.Lista.InsertarVertice(x)
	}
}

func (L *ListaAdyacencia) Enlazar(a string, b string){
	var origen *Vertice
	var destino *Vertice
	origen = L.obtenerVertice(a)
	destino = L.obtenerVertice(b)
	if origen == nil || destino == nil {
		fmt.Println("No se encontró el vértice")
		return
	}
	nuevoOrigen := &NodoLista{destino, nil, nil}
	nuevoDestino := &NodoLista{origen, nil, nil}
	origen.Adyacentes.InsertarVertice(nuevoOrigen)
	destino.Adyacentes.InsertarVertice(nuevoDestino)
}

func Contiene(buscando *Lista, elemento *Vertice) bool{
	for i := buscando.Cabeza; i != nil ; i=i.Siguiente {
		if i.Vertice == elemento {
			return true
		}
	}
	return false
}

func (L *ListaAdyacencia) Dibujar(inicio string, fin string, enlaces []Nodos, recorrido *ListaRecorrido){

	aux := &Lista{}
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	imp := L.Lista.Cabeza
	for imp != nil{
		tmp := imp
		if Contiene(aux, tmp.Vertice) == false {
			aux.InsertarVertice(&NodoLista{tmp.Vertice, nil, nil })
			if tmp.Vertice.Departamento == inicio {
				fmt.Fprintf(&cadena, "node%p[label=\"%v\" style=\"filled\" color=\"yellow\"]\n", &(*tmp.Vertice), tmp.Vertice.Departamento)
			}else{
				fmt.Fprintf(&cadena, "node%p[label=\"%v\"]\n", &(*tmp.Vertice), tmp.Vertice.Departamento)
			}
			if tmp.Vertice.Departamento == fin {
				fmt.Fprintf(&cadena, "node%p[label=\"%v\" style=\"filled\" color=\"green\"]\n", &(*tmp.Vertice), tmp.Vertice.Departamento)
			}else{
				fmt.Fprintf(&cadena, "node%p[label=\"%v\"]\n", &(*tmp.Vertice), tmp.Vertice.Departamento)
			}
		}
		temporal := tmp.Vertice.Adyacentes.Cabeza
		for temporal != nil{
			for i := 0; i < len(enlaces); i++ {
				if tmp.Vertice.Departamento == enlaces[i].Nombre {
					for j := 0; j < len(enlaces[i].Enlaces); j++ {
						if temporal.Vertice.Departamento == enlaces[i].Enlaces[j].Nombre{
							existe := false
							for k := recorrido.Cabeza; k != nil ; k = k.Siguiente {
								if tmp.Vertice.Departamento == k.Viene && enlaces[i].Enlaces[j].Nombre == k.Va{
									existe = true
								}
								if tmp.Vertice.Departamento == k.Va && enlaces[i].Enlaces[j].Nombre == k.Viene{
									existe = true
								}
							}
							if existe == true {
								fmt.Fprintf(&cadena, "node%p->node%p[label=\"%v\" dir=\"both\" color=\"red\"]\n", &(*tmp.Vertice), &(*temporal.Vertice), enlaces[i].Enlaces[j].Distancia)
							}else{
								fmt.Fprintf(&cadena, "node%p->node%p[label=\"%v\" dir=\"both\"]\n", &(*tmp.Vertice), &(*temporal.Vertice), enlaces[i].Enlaces[j].Distancia)
							}
						}
					}
				}
			}

			if Contiene(aux, temporal.Vertice) == false {
				aux.InsertarVertice(&NodoLista{temporal.Vertice, nil, nil })
				if temporal.Vertice.Departamento == inicio {
					fmt.Fprintf(&cadena, "node%p[label=\"%v\" style=\"filled\" color=\"yellow\"]\n", &(*temporal.Vertice), temporal.Vertice.Departamento)
				}else{
					fmt.Fprintf(&cadena, "node%p[label=\"%v\"]\n", &(*temporal.Vertice), temporal.Vertice.Departamento)
				}
				if temporal.Vertice.Departamento == fin {
					fmt.Fprintf(&cadena, "node%p[label=\"%v\" style=\"filled\" color=\"green\"]\n", &(*temporal.Vertice), temporal.Vertice.Departamento)
				}else{
					fmt.Fprintf(&cadena, "node%p[label=\"%v\"]\n", &(*temporal.Vertice), temporal.Vertice.Departamento)
				}
			}
			temporal = temporal.Siguiente
		}
		imp = imp.Siguiente
	}
	if recorrido.Cabeza != nil {
		fmt.Fprintf(&cadena, "Tabla [shape = none, fontsize=10, label=<\n")
		fmt.Fprintf(&cadena, "<TABLE BORDER=\"1\">\n")
		fmt.Fprintf(&cadena, "<tr><td>ITERACION</td><td>ACTUAL</td><td>SIGUIENTE</td><td>CUENTA</td></tr>\n")
		var cuenta float64 = 0
		var iteracion = 1
		for i := recorrido.Cabeza; i != nil ; i=i.Siguiente {
			cuenta = cuenta+i.Costo
			s := fmt.Sprintf("%v", cuenta)
			fmt.Fprintf(&cadena, "<tr><td>"+strconv.Itoa(iteracion)+"</td><td>"+i.Viene+"</td><td>"+i.Va+"</td><td>"+s+"</td></tr>\n")
			iteracion++
		}
		fmt.Fprintf(&cadena, "</TABLE>\n")
		fmt.Fprintf(&cadena, ">];\n")
	}

	fmt.Fprintf(&cadena, "}")
	guardarArchivo(cadena.String())
	Imagen("GrafoRecorrido.pdf")
}

func Imagen ( nombre string){
	path, _:= exec.LookPath("circo")
	cmd, _ := exec.Command(path, "-Tpdf", "..\\cliente\\src\\Recorrido\\diagrama.circo").Output()
	mode := int(0777)
	ioutil.WriteFile("..\\cliente\\src\\Recorrido\\"+nombre, cmd, os.FileMode(mode))
}

func guardarArchivo(cadena string) {
	f, err := os.Create("..\\cliente\\src\\Recorrido\\diagrama.circo")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type NodoRecorrido struct {
	Viene string
	Va string
	Costo float64
	Siguiente *NodoRecorrido
	Anterior *NodoRecorrido
}

type ListaRecorrido struct {
	Cabeza *NodoRecorrido
	Cola *NodoRecorrido
}

func (L *ListaRecorrido) InsertarRec(nuevo *NodoRecorrido){
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
}

func (L *ListaRecorrido) InsertarRecCabeza(nuevo *NodoRecorrido){
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cabeza.Anterior = nuevo
		nuevo.Siguiente = L.Cabeza
		L.Cabeza = nuevo
	}
}




