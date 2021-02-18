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
	respuesta := "No se ha cargado ningun listado de tiendas."
	if vector != nil {
		var cadena strings.Builder
		var contador int = 10
		fmt.Fprintf(&cadena, "digraph "+vector[contador-1].Departamento+"{\n")
		fmt.Fprintf(&cadena, "node[shape=record];\n")
		for i := 0; i < len(vector); i++ {
			if i<contador{
				posicion := vector[i].ListGA
				if posicion.Cabeza == nil {
					fmt.Fprintf(&cadena, "node"+vector[i].Indice+vector[i].Departamento+strconv.Itoa(vector[i].Calificacion)+"[label=\"{%v|%v}|%v|%v|%v\"];\n", vector[i].Indice,vector[i].Departamento, "", "", vector[i].Calificacion)
				}else{
					grafico(posicion.Cabeza, &cadena, posicion.Cabeza.Siguiente, vector[i].Departamento,vector[i].Indice)
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
		respuesta = "Los archivos han sido creados"
	}
	return respuesta
}

func grafico(anterior *Listas.NodoTienda, s *strings.Builder, actual *Listas.NodoTienda, departamento string, indice string) {
	if anterior != nil {
		fmt.Fprintf(s, "node%p[label=\"{%v|%v}|%v|%v|%v\"];\n", &(*anterior), indice,departamento, anterior.NombreTienda, anterior.Contacto, anterior.Calificacion)
		if actual != nil && actual!=anterior{
			fmt.Fprintf(s, "node%p->node%p;\n", &(*actual), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*anterior), &(*actual))
		}
		grafico(anterior.Siguiente, s, anterior, departamento, indice)
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
	cmd, _ := exec.Command(path, "-Tpdf", "./"+nombreArchivo).Output()
	mode := int(0777)
	ioutil.WriteFile(nombreArchivo+".pdf", cmd, os.FileMode(mode))
}