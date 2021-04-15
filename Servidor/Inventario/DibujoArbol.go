package Inventario

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var colores = []string{"red", "yellow", "gray", "blue", "violet", "green", "orange"}
func (L *Arbol) Generar(){
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph Arbol{")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if L.Raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"{Nombre: %v|Código: %v}|{Cantidad Disponible: %v|Precio: %v}\" color=\""+colores[rand.Intn(7)]+"\"];\n", &(*L.Raiz), L.Raiz.NombreProducto, L.Raiz.Codigo, L.Raiz.Cantidad, L.Raiz.Precio)
		L.generar(&cadena, L.Raiz, L.Raiz.Izq, true)
		L.generar(&cadena, L.Raiz, L.Raiz.Der, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "ArbolProductos")
}

func (L *Arbol) generar(cadena *strings.Builder, padre *NodoArbol, actual *NodoArbol, izquierda bool){
	if actual != nil {
		fmt.Fprintf(cadena, "node%p[label=\"{Nombre: %v|Código: %v}|{Cantidad Disponible: %v|Precio: %v}\" color=\""+colores[rand.Intn(7)]+"\"];\n", &(*actual), actual.NombreProducto, actual.Codigo, actual.Cantidad, actual.Precio)
		if izquierda{
			fmt.Fprintf(cadena, "node%p:f0->node%p:f1\n", &(*padre), &(*actual))
		}else{
			fmt.Fprintf(cadena, "node%p:f2->node%p:f1\n", &(*padre), &(*actual))
		}
		L.generar(cadena, actual, actual.Izq, true)
		L.generar(cadena, actual, actual.Der, false)
	}
}

func guardarArchivo(cadena string, nombreArchivo string) {
	f, err := os.Create("..\\cliente\\src\\ImagenArbol\\"+nombreArchivo+".dot")
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
	fmt.Println(l)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "..\\cliente\\src\\ImagenArbol\\"+nombreArchivo+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile("..\\cliente\\src\\ImagenArbol\\"+nombreArchivo+".pdf", cmd, os.FileMode(mode))
}