package ArbolMerkle


import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoComentario struct{
	Hash string
	Tipo string
	Tienda string
	Departamento string
	Calificacion int
	Respondiendo string
	Dpi int
	Fecha string
	Comentario string
	Derecha *NodoComentario
	Izquierda *NodoComentario
}

type ArbolComentarios struct{
	Raiz *NodoComentario
}

func NuevoNodoComentarios(hash string, tipo string, tienda string, departamento string, calificacion int, respondiendo string, dpi int, fecha string, comentario string, derecha *NodoComentario, izquierda *NodoComentario) *NodoComentario{
	return &NodoComentario{hash, tipo, tienda, departamento, calificacion, respondiendo, dpi,fecha,comentario , derecha, izquierda}
}

func NuevoArbolComentarios()*ArbolComentarios{
	return &ArbolComentarios{}
}

func (this *ArbolComentarios) Insertar(hash string, tipo string, tienda string, departamento string, calificacion int, respondiendo string, dpi int, fecha string, comentario string){
	n := NuevoNodoComentarios(hash, tipo, tienda, departamento, calificacion, respondiendo, dpi,fecha,comentario, nil,nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodoComentarios(uno.String(), "", "", "",-1,"",-1,"", "",nil,nil))
		this.ConstruirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.ConstruirArbol(lista)
	}
}

func (this *ArbolComentarios) ObtenerLista() *list.List{
	lista := list.New()
	obtenerListaComentarios(lista, this.Raiz.Izquierda)
	obtenerListaComentarios(lista, this.Raiz.Derecha)
	return lista
}

func obtenerListaComentarios(lista *list.List, actual *NodoComentario){
	if actual != nil {
		obtenerListaComentarios(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		obtenerListaComentarios(lista, actual.Derecha)
	}
}

func (this *ArbolComentarios) ConstruirArbol(lista *list.List){
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
		lista.PushBack(NuevoNodoComentarios(uno.String(),"", "","",-1,"", -1, "", "", nil, nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoComentario)
		nodo2 := segundo.Value.(*NodoComentario)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodoComentarios(Hash.String(),"", "","",-1,"", -1, "", "", nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoComentario)
}

func (L *ArbolComentarios) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph ArbolTiendas{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		if L.Raiz.Izquierda != nil && L.Raiz.Derecha != nil {
			var hash strings.Builder
			fmt.Fprintf(&hash, "%x", sha256.Sum256([]byte(L.Raiz.Izquierda.Hash+L.Raiz.Derecha.Hash)))
			if hash.String() == L.Raiz.Hash {
				fmt.Fprintf(&cadena, "node%p[color=\"green\" label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
			}else{
				fmt.Fprintf(&cadena, "node%p[color=\"red\" label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
			}
		}
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolComentarios")
}

func (L *ArbolComentarios) generar(cadena *strings.Builder, padre *NodoComentario, actual *NodoComentario, izquierda bool){
	if actual != nil {
		if actual.Tienda != "" && actual.Izquierda == nil && actual.Derecha == nil{

			var Hash strings.Builder
			fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(actual.Tienda+actual.Departamento+strconv.Itoa(actual.Calificacion)+strconv.Itoa(actual.Dpi)+actual.Comentario+actual.Fecha)))
			if actual.Hash == Hash.String() {
				fmt.Fprintf(cadena, "node%p[color=\"green\" label=\"{Hash: %v|Tipo: %s|Respondiendo: %s|DPI: %d|Comentario: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.Respondiendo, actual.Dpi, actual.Comentario)
			}else{
				fmt.Fprintf(cadena, "node%p[color=\"green\" label=\"{Hash: %v|Tipo: %s|Respondiendo: %s|DPI: %d|Comentario: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.Respondiendo, actual.Dpi, actual.Comentario)
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

func (L *ArbolComentarios) Arreglar(raiz *NodoComentario, lista *list.List) *list.List{
	if raiz != nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil {
			if raiz.Tipo != "" {
				var Hash strings.Builder
				fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(raiz.Tienda+raiz.Departamento+strconv.Itoa(raiz.Calificacion)+strconv.Itoa(raiz.Dpi)+raiz.Comentario+raiz.Fecha)))
				raiz.Hash = Hash.String()
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