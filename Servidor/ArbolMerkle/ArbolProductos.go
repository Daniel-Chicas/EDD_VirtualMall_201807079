package ArbolMerkle

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strings"
)

type NodoProductos struct{
	Hash string
	Tipo string
	Tienda string
	Departamento string
	Calificacion int
	NombreProducto string
	Codigo int
	Descripcion string
	Precio float64
	Cantidad int
	Imagen string
	Almacenamiento string
	Derecha *NodoProductos
	Izquierda *NodoProductos
}

type ArbolProductos struct{
	Raiz *NodoProductos
}

func NuevoNodoProducto(hash string, tipo string, tienda string, departamento string, calificacion int, nombreProducto string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string, derecha *NodoProductos, izquierda *NodoProductos) *NodoProductos{
	return &NodoProductos{hash, tipo, tienda, departamento, calificacion, nombreProducto,codigo,descripcion,precio,cantidad,imagen,almacenamiento, derecha, izquierda}
}

func NuevoArbolProducto()*ArbolProductos{
	return &ArbolProductos{}
}

func (this *ArbolProductos) Insertar(hash string, tipo string, tienda string, departamento string, calificacion int, nombreProducto string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string){
	n := NuevoNodoProducto(hash, tipo, tienda, departamento, calificacion, nombreProducto,codigo,descripcion,precio,cantidad,imagen,almacenamiento,nil,nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodoProducto(uno.String(), "", "", "",-1,"",-1,"",-1,0, "", "", nil,nil))
		this.construirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.construirArbol(lista)
	}
}

func (this *ArbolProductos) ObtenerLista() *list.List{
	lista := list.New()
	obtenerListaProductos(lista, this.Raiz.Izquierda)
	obtenerListaProductos(lista, this.Raiz.Derecha)
	return lista
}

func obtenerListaProductos(lista *list.List, actual *NodoProductos){
	if actual != nil {
		obtenerListaProductos(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		obtenerListaProductos(lista, actual.Derecha)
	}
}

func (this *ArbolProductos) construirArbol(lista *list.List){
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
		lista.PushBack(NuevoNodoProducto(uno.String(),"", "", "",-1,"",-1,"",-1,0, "", "", nil,nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoProductos)
		nodo2 := segundo.Value.(*NodoProductos)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodoProducto(Hash.String(),"", "","",-1,"",-1,"",-1,-1,"","", nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoProductos)
}

func (L *ArbolProductos) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph ArbolTiendas{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"{Hash: %v|HashIzquierdo: %v|HashDerecho: %v}\" ];\n", &(*L.Raiz), L.Raiz.Hash, L.Raiz.Izquierda.Hash, L.Raiz.Derecha.Hash)
		L.generar(&cadena, L.Raiz, L.Raiz.Izquierda, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolProductos")
}

func (L *ArbolProductos) generar(cadena *strings.Builder, padre *NodoProductos, actual *NodoProductos, izquierda bool){
	if actual != nil {
		if actual.Tienda != "" {
			fmt.Fprintf(cadena, "node%p[label=\"{Hash: %v|Tipo: %s|Tienda: %s|Codigo: %d}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.Tienda, actual.Codigo)
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