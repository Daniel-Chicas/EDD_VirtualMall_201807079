package ArbolMerkle

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strings"
)

type NodoComentarioProducto struct{
	Hash string
	Tipo string
	Tienda string
	Departamento string
	Calificacion int
	CodigoProducto int
	Respondiendo string
	Dpi int
	Fecha string
	Comentario string
	Derecha *NodoComentarioProducto
	Izquierda *NodoComentarioProducto
}

type ArbolComentariosProducto struct{
	Raiz *NodoComentarioProducto
}

func NuevoNodoComentariosProducto(hash string, tipo string, tienda string, departamento string, calificacion int, respondiendo string, Codigo int, dpi int, fecha string, comentario string, derecha *NodoComentarioProducto, izquierda *NodoComentarioProducto) *NodoComentarioProducto{
	return &NodoComentarioProducto{hash, tipo, tienda, departamento, calificacion, Codigo,respondiendo, dpi,fecha,comentario , derecha, izquierda}
}

func NuevoArbolComentariosProducto()*ArbolComentariosProducto{
	return &ArbolComentariosProducto{}
}
//NodoComentarioProducto
func (this *ArbolComentariosProducto) Insertar(hash string, tipo string, tienda string, departamento string, calificacion int, codigo int, respondiendo string, dpi int, fecha string, comentario string){
	n := NuevoNodoComentariosProducto(hash, tipo, tienda, departamento, calificacion, respondiendo, codigo, dpi,fecha,comentario, nil,nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodoComentariosProducto(uno.String(), "", "", "",-1,"",-1,-1,"", "",nil,nil))
		this.construirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.construirArbol(lista)
	}
}

func (this *ArbolComentariosProducto) ObtenerLista() *list.List{
	lista := list.New()
	obtenerListaComentariosProductos(lista, this.Raiz.Izquierda)
	obtenerListaComentariosProductos(lista, this.Raiz.Derecha)
	return lista
}

func obtenerListaComentariosProductos(lista *list.List, actual *NodoComentarioProducto){
	if actual != nil {
		obtenerListaComentariosProductos(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		obtenerListaComentariosProductos(lista, actual.Derecha)
	}
}

func (this *ArbolComentariosProducto) construirArbol(lista *list.List){
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
		lista.PushBack(NuevoNodoComentariosProducto(uno.String(),"", "","",-1,"", -1, -1, "", "", nil, nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoComentarioProducto)
		nodo2 := segundo.Value.(*NodoComentarioProducto)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodoComentariosProducto(Hash.String(),"", "","",-1,"", -1, -1, "", "", nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoComentarioProducto)
}

func (L *ArbolComentariosProducto) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph ArbolTiendas{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolComentariosProductos")
}

func (L *ArbolComentariosProducto) generar(cadena *strings.Builder, padre *NodoComentarioProducto, actual *NodoComentarioProducto, izquierda bool){
	if actual != nil {
		if actual.Tienda != "" {
			fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|Tipo: %s|Código: %d|Respondiendo: %s|DPI: %d|Comentario: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.CodigoProducto, actual.Respondiendo, actual.Dpi, actual.Comentario)
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
