package Reportes

import (
	"../Listas"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Lista struct {
}

func (L *Lista) Arreglo(vector []Listas.NodoArray) string{
	var cadena strings.Builder
	var contador int = 10
	fmt.Fprintf(&cadena, "digraph "+vector[contador-1].Departamento+"{\n")
	fmt.Fprintf(&cadena, "node[shape=record];\n")
	for i := 0; i < len(vector); i++ {
		if i<contador{
			posicion := vector[i].ListGA
			if posicion.Cabeza == nil {
				fmt.Fprintf(&cadena, "node"+vector[i].Indice+vector[i].Departamento+strconv.Itoa(vector[i].Calificacion)+"[label=\"%v|%v|%v|%v\"];\n", "", "", "", vector[i].Calificacion)
				//grafico(posicion.Cabeza, &cadena, nil)
			}else{
				grafico(posicion.Cabeza, &cadena, posicion.Cabeza.Siguiente)
			}
		}
		if i == contador-1 {
			fmt.Fprintf(&cadena, "}")
			guardarArchivo(cadena.String(), "Archivo_"+vector[i].Departamento+".dot")
			cadena.Reset()
			fmt.Fprintf(&cadena, "digraph "+vector[contador-1].Departamento+"{\n")
			fmt.Fprintf(&cadena, "node[shape=record];\n")
			contador = contador +10
		}
	}

	return "Los archivos han sido creados."
}

func grafico(anterior *Listas.NodoTienda, s *strings.Builder, actual *Listas.NodoTienda) {
	if anterior != nil {
		fmt.Fprintf(s, "node%p[label=\"%v|%v|%v|%v\"];\n", &(*anterior), anterior.NombreTienda, anterior.Descripcion, anterior.Contacto, anterior.Calificacion)
		if actual != nil && actual!=anterior{
			fmt.Fprintf(s, "node%p->node%p;\n", &(*actual), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*anterior), &(*actual))
		}
		grafico(anterior.Siguiente, s, anterior)
	}
}

func guardarArchivo(cadena string, nombreArchivo string) {
	f, err := os.Create(nombreArchivo)
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
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./"+nombreArchivo).Output()
	mode := int(0777)
	ioutil.WriteFile(nombreArchivo+".png", cmd, os.FileMode(mode))
}