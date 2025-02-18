package ArbolMerkle

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Nodo struct{
	Hash string
	Tipo string
	Indice string
	Departamento string
	NombreTienda string
	DescripcionTienda string
	ContactoTienda string
	Calificacion int
	Logo string
	Derecha *Nodo
	Izquierda *Nodo
}

type Arbol struct{
	Raiz *Nodo
}

func NuevoNodo(hash string, tipo string, indice string, depa string, nombre string, descripcion string, contacto string, calificacion int, logo string, derecha *Nodo, izquierda *Nodo) *Nodo{
	return &Nodo{hash, tipo, indice, depa, nombre, descripcion, contacto, calificacion, logo, derecha, izquierda}
}

func NuevoArbol()*Arbol{
	return &Arbol{}
}

func (this *Arbol) Insertar(hash string, tipo string, indice string, depa string, nombre string, descripcion string, contacto string, calificacion int, logo string){
	n := NuevoNodo(hash, tipo, indice, depa, nombre, descripcion, contacto, calificacion, logo, nil, nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodo(uno.String(), "", "","","","","",-1,"", nil, nil))
		this.ConstruirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.ConstruirArbol(lista)
	}
}

func (this *Arbol) ObtenerLista() *list.List{
	lista := list.New()
	obtenerLista(lista, this.Raiz.Izquierda)
	obtenerLista(lista, this.Raiz.Derecha)
	return lista
}

func obtenerLista(lista *list.List, actual *Nodo){
	if actual != nil {
		obtenerLista(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.Derecha)
	}
}

func (this *Arbol) ConstruirArbol(lista *list.List){
	size := float64(lista.Len())
	cant := 1
	for (size/2) > 1 {
		cant++
		size = size/2
	}
	nodostotales := math.Pow(2, float64(cant))
	for lista.Len() < int(nodostotales) {
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodo(uno.String(), "", "","","","","",-1,"", nil, nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*Nodo)
		nodo2 := segundo.Value.(*Nodo)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodo(Hash.String(), "", "", "", "", "", "", -1, "", nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*Nodo)
}

func (L *Arbol) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph ArbolTiendas{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		if L.Raiz.Izquierda != nil && L.Raiz.Derecha != nil {
			var hash strings.Builder
			fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte(L.Raiz.Izquierda.Hash+L.Raiz.Derecha.Hash)))
			if hash.String() == L.Raiz.Hash {
				fmt.Fprintf(&cadena, "node%p[color=\"green\"  label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
			}else{
				fmt.Fprintf(&cadena, "node%p[color=\"red\" label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
			}
		}
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolTiendas")
}

func (L *Arbol) generar(cadena *strings.Builder, padre *Nodo, actual *Nodo, izquierda bool){
	if actual != nil {
		if actual.NombreTienda != "" && actual.Izquierda == nil && actual.Derecha == nil{
			var hash strings.Builder
			fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte(actual.Indice+actual.Departamento+actual.NombreTienda+actual.DescripcionTienda+actual.ContactoTienda+strconv.Itoa(actual.Calificacion)+actual.Logo)))
			if actual.Hash == hash.String() {
				fmt.Fprintf(cadena, "node%p[color=\"green\"  label=\"{Hash: %v|Tienda: %s|Contacto: %s|Calificacion: %d}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.NombreTienda, actual.Calificacion)
			}else{
				fmt.Fprintf(cadena, "node%p[color=\"red\" label=\"{Hash: %v|Tienda: %s|Contacto: %s|Calificacion: %d}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.NombreTienda, actual.Calificacion)
			}
		}else{
			if actual.Izquierda != nil && actual.Derecha != nil {
				var hash strings.Builder
				fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte(actual.Izquierda.Hash+actual.Derecha.Hash)))
				if actual.Hash == hash.String() {
					fmt.Fprintf(cadena, "node%p[color = \"green\" label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*actual), actual.Hash, actual.Izquierda.Hash, actual.Derecha.Hash)
				}else{
					fmt.Fprintf(cadena, "node%p[color = \"red\" label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*actual), actual.Hash, actual.Izquierda.Hash, actual.Derecha.Hash)
				}
			}else{
				var hash strings.Builder
				fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte("-1")))
				if actual.Hash == hash.String() {
					fmt.Fprintf(cadena, "node%p[color=\"green\" label=\"{Hash: %v|Agrega: -1}\" ];\n", &(*actual), actual.Hash)
				}else{
					fmt.Fprintf(cadena, "node%p[color=\"red\" label=\"{Hash: %v|Agrega: -1}\" ];\n", &(*actual), actual.Hash)
				}
			}
		}
		if izquierda{
			fmt.Fprintf(cadena, "node%p:f0->node%p:f1\n", &(*padre), &(*actual))
		}else{
			fmt.Fprintf(cadena, "node%p:f2->node%p:f1\n", &(*padre), &(*actual))
		}
		L.generar(cadena, actual, actual.Izquierda, true)
		L.generar(cadena, actual, actual.Derecha, false)
	}
}

func guardarArchivo(cadena string, nombreArchivo string) {
	f, err := os.Create("..\\cliente\\src\\ArbolesMerkle\\diagrama.dot")
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
	fmt.Println(strconv.Itoa(l)+" bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "..\\cliente\\src\\ArbolesMerkle\\diagrama.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("..\\cliente\\src\\ArbolesMerkle\\"+nombreArchivo+".pdf", cmd, os.FileMode(mode))
}

func (L *Arbol) Arreglar(raiz *Nodo, lista *list.List) *list.List{
	if raiz != nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil {
			if raiz.NombreTienda != "" {
				var hash strings.Builder
				fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte(raiz.Indice+raiz.Departamento+raiz.NombreTienda+raiz.DescripcionTienda+raiz.ContactoTienda+strconv.Itoa(raiz.Calificacion)+raiz.Logo)))
				raiz.Hash = hash.String()
				lista.PushBack(raiz)
			}else{
				var hash strings.Builder
				fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte("-1")))
				raiz.Hash = hash.String()
				lista.PushBack(raiz)
			}
		}
		L.Arreglar(raiz.Izquierda, lista)
		L.Arreglar(raiz.Derecha, lista)
	}
	return lista
}








