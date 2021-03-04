package Reportes

import (
	"../Listas"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Lista struct {
}

func (L *Lista) Arreglo(vector []Listas.NodoArray) string{
	var cuenta = 0
	respuesta := "No se ha cargado ningun listado de tiendas."
	if vector != nil {
		var cadena strings.Builder
		var contador int = 5
		fmt.Fprintf(&cadena, "digraph Daniel"+strconv.Itoa(cuenta)+"{\n")
		fmt.Fprintf(&cadena, "node[shape=record];\n")

		for i := 0; i < len(vector); i++ {

			if i<contador{
				posicion := vector[i].ListGA
				colores := [10]string{"gray", "red", "blue", "yellow", "green", "orange", "brown", "pink", "violet", "purple"}
				fmt.Fprintf(&cadena, "node"+strconv.Itoa(i)+"[color="+colores[rand.Intn(9)]+", fillcolor="+colores[rand.Intn(9)]+", label=\"{%v|%v}|{%v|%v}\"];\n", "Indice: "+vector[i].Indice, vector[i].Departamento, "Pos: "+strconv.Itoa(i) , "Calificación: "+strconv.Itoa(vector[i].Calificacion))
				if posicion.Cabeza != nil {
					grafico(&cadena, &posicion, i, colores)
				}
			}
			if i == contador-1 {
				fmt.Fprintf(&cadena, "}")
				guardarArchivo(cadena.String(), "Archivo_"+strconv.Itoa(cuenta)+".dot")
				cadena.Reset()
				cuenta++
				fmt.Fprintf(&cadena, "digraph Daniel"+strconv.Itoa(cuenta)+"{\n")
				fmt.Fprintf(&cadena, "node[shape=record];\n")
				contador = contador +5
			}
		}
		respuesta = "Los archivos han sido creados"
	}
	return respuesta
}

func grafico(s *strings.Builder, lista *Listas.ListaGA, i int, colores [10]string) {
	imp := lista.Cabeza
	for imp!=nil {
		sig := imp.Siguiente
		ant := imp.Anterior
		fmt.Fprintf(s, "node%p[color="+colores[rand.Intn(9)]+", label=\"{%v|%v}\"];\n", &(*imp), imp.NombreTienda, imp.Contacto)
		if imp == lista.Cabeza {
			fmt.Fprintf(s, "node%v->node%p;\n", strconv.Itoa(i), &(*imp))
			if imp.Siguiente!=nil {
				fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*sig))
			}
		}else if imp.Siguiente !=nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*ant))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*sig))
		}else if imp.Siguiente == nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*ant))
		}
		imp = imp.Siguiente
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