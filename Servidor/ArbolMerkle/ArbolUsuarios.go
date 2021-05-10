package ArbolMerkle

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoUsuarios struct{
	Hash string
	Tipo string
	DPI int
	Nombre string
	Correo string
	Password string
	Cuenta string
	Derecha *NodoUsuarios
	Izquierda *NodoUsuarios
}

type ArbolUsuarios struct{
	Raiz *NodoUsuarios
}

func NuevoNodoUsuarios(hash string, tipo string, dpi int, nombre string, correo string, password string, cuenta string, derecha *NodoUsuarios, izquierda *NodoUsuarios) *NodoUsuarios{
	return &NodoUsuarios{hash, tipo, dpi, nombre, correo, password, cuenta, derecha, izquierda}
}

func NuevoArbolUsuarios()*ArbolUsuarios{
	return &ArbolUsuarios{}
}

func (this *ArbolUsuarios) Insertar(hash string, tipo string, dpi int, nombre string, correo string, password string, cuenta string){
	n := NuevoNodoUsuarios(hash, tipo, dpi,nombre,correo,password,cuenta, nil,nil)
	if this.Raiz == nil{
		lista := list.New()
		lista.PushBack(n)
		var uno strings.Builder
		fmt.Fprintf(&uno, "%x", sha256.Sum256([]byte("-1")))
		lista.PushBack(NuevoNodoUsuarios(uno.String(), "", -1, "","","","",nil,nil))
		this.ConstruirArbol(lista)
	}else{
		lista := this.ObtenerLista()
		lista.PushBack(n)
		this.ConstruirArbol(lista)
	}
}

func (this *ArbolUsuarios) ObtenerLista() *list.List{
	lista := list.New()
	obtenerListaUsuarios(lista, this.Raiz.Izquierda)
	obtenerListaUsuarios(lista, this.Raiz.Derecha)
	return lista
}

func obtenerListaUsuarios(lista *list.List, actual *NodoUsuarios){
	if actual != nil {
		obtenerListaUsuarios(lista, actual.Izquierda)
		if actual.Derecha == nil && actual.DPI != -1 {
			lista.PushBack(actual)
		}
		obtenerListaUsuarios(lista, actual.Derecha)
	}
}

func (this *ArbolUsuarios) ConstruirArbol(lista *list.List){
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
		lista.PushBack(NuevoNodoUsuarios(uno.String(), "", -1, "","","","",nil,nil))
	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoUsuarios)
		nodo2 := segundo.Value.(*NodoUsuarios)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nodo1.Hash+nodo2.Hash)))
		nuevo := NuevoNodoUsuarios(Hash.String(),"", -1,"","","","", nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.Raiz = lista.Front().Value.(*NodoUsuarios)
}

func (L *ArbolUsuarios) Generar(){
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
	guardarArchivo(cadena.String(), "ArbolUsuarios")
}

func (L *ArbolUsuarios) generar(cadena *strings.Builder, padre *NodoUsuarios, actual *NodoUsuarios, izquierda bool){
	if actual != nil {
		if actual.Nombre != "" && actual.Izquierda == nil && actual.Derecha == nil{

			var Hash strings.Builder
			fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(strconv.Itoa(actual.DPI)+actual.Nombre+actual.Correo+actual.Password)))
			if actual.Hash == Hash.String() {
				fmt.Fprintf(cadena, "node%p[color=\"green\" label=\"{Hash: %v|Tipo: %s|DPI: %d|Nombre: %s|Correo: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.DPI, actual.Nombre, actual.Correo)
			}else{
				fmt.Fprintf(cadena, "node%p[color=\"red\" label=\"{Hash: %v|Tipo: %s|DPI: %d|Nombre: %s|Correo: %s}\" ];\n", &(*actual), actual.Hash, actual.Tipo, actual.DPI, actual.Nombre, actual.Correo)
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

func (L *ArbolUsuarios) Arreglar(raiz *NodoUsuarios, lista *list.List) *list.List{
	if raiz != nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil {
			if raiz.Tipo != "" {
				var Hash strings.Builder
				fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(strconv.Itoa(raiz.DPI)+raiz.Nombre+raiz.Correo+raiz.Password)))
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
