package Usuarios

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var colores = []string{"brown", "yellow", "gray", "blue", "violet", "green", "orange"}

func (A *ArbolB) Grafico(tipo string){
	cadena := strings.Builder{}
	fmt.Fprintf(&cadena, "digraph G{ \n")
	fmt.Fprintf(&cadena, "node[shape=record] \n")
	m := make(map[string]*Pagina)
	if tipo == "Si" {
		graficoEC(A.Raiz, &cadena, m, nil, 0)
		fmt.Fprintf(&cadena, "}")
		guardarArchivo(cadena.String())
		Imagen("arbolTodoCifrado.pdf")
	}else if tipo == "No"{
		grafico(A.Raiz, &cadena, m, nil, 0)
		fmt.Fprintf(&cadena, "}")
		guardarArchivo(cadena.String())
		Imagen("arbol.pdf")
	}else if tipo == "Medio"{
		graficoEM(A.Raiz, &cadena, m, nil, 0)
		fmt.Fprintf(&cadena, "}")
		guardarArchivo(cadena.String())
		Imagen("arbolMedio.pdf")
	}
}

func grafico(actual *Pagina, cadena *strings.Builder, arreglo map[string]*Pagina, padre *Pagina, posicion int){
	if actual == nil{
		return
	}
	t := 0
	contiene := arreglo[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arreglo[fmt.Sprint(&(*actual))] = actual
		return
	}else{
		arreglo[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cadena, "node%p [style = filled color=\""+colores[rand.Intn(7)]+"\" label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			return
		}else{
			if enlace{
				if i != actual.Maximo-1 {
					fmt.Fprintf(cadena, "<f%d>|", t)
				}else{
					fmt.Fprintf(cadena, "<f%d>", t)
					break
				}
				enlace = false
				i--
				t++
			}else{
				fmt.Fprintf(cadena, "{DPI: %s|Nombre: %s|Tipo Usuario: %s|Correo: %s|Password: %s}|", actual.Llaves[i].Usuario.DPI, actual.Llaves[i].Usuario.Nombre, actual.Llaves[i].Usuario.Cuenta, actual.Llaves[i].Usuario.Correo, actual.Llaves[i].Usuario.Contra)
				t++
				enlace = true
				if i<actual.Maximo-1 {
					if actual.Llaves[i+1] == nil {
						fmt.Fprintf(cadena, "<f%d>", t)
						t++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cadena, "\"] \n")
	ti := 0
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			break
		}
		grafico(actual.Llaves[i].Izq, cadena, arreglo, actual, ti)
		ti++
		ti++
		grafico(actual.Llaves[i].Der, cadena, arreglo, actual, ti)
		ti++
		ti--
	}
	if padre != nil {
		fmt.Fprintf(cadena, "node%p:f%d->node%p\n", &(*padre), posicion, &(*actual))
	}
}

/*

	nuevo.DPI = fmt.Sprintf("%x", sha256.Sum256([]byte(datosUsuario.DPI)))
	nuevo.Nombre = fmt.Sprintf("%x", )
	nuevo.Correo = fmt.Sprintf("%x", sha256.Sum256([]byte(datosUsuario.Correo)))
	nuevo.Contra = fmt.Sprintf("%x", sha256.Sum256([]byte(datosUsuario.Contra)))
	nuevo.Cuenta = fmt.Sprintf("%x", sha256.Sum256([]byte(datosUsuario.Cuenta)))


*/



func graficoEC(actual *Pagina, cadena *strings.Builder, arreglo map[string]*Pagina, padre *Pagina, posicion int){
	if actual == nil{
		return
	}
	t := 0
	contiene := arreglo[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arreglo[fmt.Sprint(&(*actual))] = actual
		return
	}else{
		arreglo[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cadena, "node%p [style = filled color=\""+colores[rand.Intn(7)]+"\" label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			return
		}else{
			if enlace{
				if i != actual.Maximo-1 {
					fmt.Fprintf(cadena, "<f%d>|", t)
				}else{
					fmt.Fprintf(cadena, "<f%d>", t)
					break
				}
				enlace = false
				i--
				t++
			}else{
				fmt.Fprintf(cadena, "{DPI: %x|Nombre: %x|Tipo Usuario: %x|Correo: %x|Password: %x}|", sha256.Sum256([]byte(actual.Llaves[i].Usuario.DPI)), sha256.Sum256([]byte(actual.Llaves[i].Usuario.Nombre)), sha256.Sum256([]byte(actual.Llaves[i].Usuario.Cuenta)), sha256.Sum256([]byte(actual.Llaves[i].Usuario.Correo)), sha256.Sum256([]byte(actual.Llaves[i].Usuario.Contra)))
				t++
				enlace = true
				if i<actual.Maximo-1 {
					if actual.Llaves[i+1] == nil {
						fmt.Fprintf(cadena, "<f%d>", t)
						t++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cadena, "\"] \n")
	ti := 0
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			break
		}
		graficoEC(actual.Llaves[i].Izq, cadena, arreglo, actual, ti)
		ti++
		ti++
		graficoEC(actual.Llaves[i].Der, cadena, arreglo, actual, ti)
		ti++
		ti--
	}
	if padre != nil {
		fmt.Fprintf(cadena, "node%p:f%d->node%p\n", &(*padre), posicion, &(*actual))
	}
}

func graficoEM(actual *Pagina, cadena *strings.Builder, arreglo map[string]*Pagina, padre *Pagina, posicion int){
	if actual == nil{
		return
	}
	t := 0
	contiene := arreglo[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arreglo[fmt.Sprint(&(*actual))] = actual
		return
	}else{
		arreglo[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cadena, "node%p [style = filled color=\""+colores[rand.Intn(7)]+"\" label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			return
		}else{
			if enlace{
				if i != actual.Maximo-1 {
					fmt.Fprintf(cadena, "<f%d>|", t)
				}else{
					fmt.Fprintf(cadena, "<f%d>", t)
					break
				}
				enlace = false
				i--
				t++
			}else{
				fmt.Fprintf(cadena, "{DPI: %x|Nombre: %s|Tipo Usuario: %s|Correo: %x|Password: %x}|", sha256.Sum256([]byte(actual.Llaves[i].Usuario.DPI)), actual.Llaves[i].Usuario.Nombre, actual.Llaves[i].Usuario.Cuenta, sha256.Sum256([]byte(actual.Llaves[i].Usuario.Correo)), sha256.Sum256([]byte(actual.Llaves[i].Usuario.Contra)))
				t++
				enlace = true
				if i<actual.Maximo-1 {
					if actual.Llaves[i+1] == nil {
						fmt.Fprintf(cadena, "<f%d>", t)
						t++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cadena, "\"] \n")
	ti := 0
	for i := 0; i < actual.Maximo; i++ {
		if actual.Llaves[i] == nil {
			break
		}
		graficoEM(actual.Llaves[i].Izq, cadena, arreglo, actual, ti)
		ti++
		ti++
		graficoEM(actual.Llaves[i].Der, cadena, arreglo, actual, ti)
		ti++
		ti--
	}
	if padre != nil {
		fmt.Fprintf(cadena, "node%p:f%d->node%p\n", &(*padre), posicion, &(*actual))
	}
}


func Imagen ( nombre string){
	path, _:= exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", ".\\cliente\\src\\ArbolesUsuarios\\diagrama.dot").Output()
	mode := int(0777)
	ioutil.WriteFile(".\\cliente\\src\\ArbolesUsuarios\\"+nombre, cmd, os.FileMode(mode))
}


func guardarArchivo(cadena string) {
	f, err := os.Create(".\\cliente\\src\\ArbolesUsuarios\\diagrama.dot")
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
}