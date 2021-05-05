package ArbolMerkle

import (
	"../Grafo"
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strings"
)

type NodoPedidos struct{
	Hash string
	Tipo string
	Fecha string
	Tienda string
	Departamento string
	Calificacion int
	Cliente int
	Producto int
	Cantidad int
	Recorrido []GrafoRecorrido.NodoRecorrido
	Derecha *NodoPedidos
	Izquierda *NodoPedidos
}

type ArbolPedidos struct{
	Raiz *NodoPedidos
}

func NuevoNodoPedido(hash string, tipo string,fecha string, tienda string, departamento string, calificacion int, dpi int, producto int, cantidad int, recorrido []GrafoRecorrido.NodoRecorrido, derecha *NodoPedidos, izquierda *NodoPedidos) *NodoPedidos{
	return &NodoPedidos{hash, tipo, fecha, tienda, departamento, calificacion,dpi,producto , cantidad, recorrido, derecha, izquierda}
}

func NuevoArbolPedidos()*ArbolPedidos{
	return &ArbolPedidos{}
}

func (this *ArbolPedidos) Insertar(hash string, tipo string,fecha string, tienda string, departamento string, calificacion int, dpi int, producto int, cantidad int, recorrido []GrafoRecorrido.NodoRecorrido){
	n := NuevoNodoPedido(hash, tipo, fecha, tienda, departamento, calificacion, dpi, producto, cantidad, recorrido, nil,nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodoPedido(uno.String(), "", "", "-1","",-1,-1,-1,0,nil,nil,nil))
		this.construirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.construirArbol(lista)
	}
}

func (this *ArbolPedidos) ObtenerLista() *list.List{
	lista := list.New()
	obtenerListaPedidos(lista, this.Raiz.Izquierda)
	obtenerListaPedidos(lista, this.Raiz.Derecha)
	return lista
}

func obtenerListaPedidos(lista *list.List, actual *NodoPedidos){
	if actual != nil {
		obtenerListaPedidos(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		obtenerListaPedidos(lista, actual.Derecha)
	}
}

func (this *ArbolPedidos) construirArbol(lista *list.List){
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
		lista.PushBack(NuevoNodoPedido(uno.String(), "", "", "","",-1,-1,-1, 0,nil,nil,nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoPedidos)
		nodo2 := segundo.Value.(*NodoPedidos)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodoPedido(Hash.String(),"", "","","",-1,-1,-1, 0,nil, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoPedidos)
}

func (L *ArbolPedidos) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph ArbolTiendas{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolPedidos")
}

func (L *ArbolPedidos) generar(cadena *strings.Builder, padre *NodoPedidos, actual *NodoPedidos, izquierda bool){
	if actual != nil {
		if actual.Tienda != "" {
			fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|Tipo: %s|Cliente: %d|Codigo: %d}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.Cliente, actual.Producto)
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