package ArbolMerkle


import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
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
		this.construirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.construirArbol(lista)
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

func (this *ArbolComentarios) construirArbol(lista *list.List){
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
		fmt.Fprintf(&cadena, "node%p[label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolComentarios")
}

func (L *ArbolComentarios) generar(cadena *strings.Builder, padre *NodoComentario, actual *NodoComentario, izquierda bool){
	if actual != nil {
		if actual.Tienda != "" {
			fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|Tipo: %s|Respondiendo: %s|DPI: %d|Comentario: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.Respondiendo, actual.Dpi, actual.Comentario)
		}else{
			if actual.Izquierda != nil && actual.Derecha != nil {
				fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*actual), actual.Hash, actual.Izquierda.Hash, actual.Derecha.Hash)
			}else{
				fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|Agrega: -1}\" ];\n", &(*actual), actual.Hash)
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