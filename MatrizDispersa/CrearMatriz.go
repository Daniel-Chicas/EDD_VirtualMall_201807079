package MatrizDispersa

import (
	"fmt"
	"strconv"
	"strings"
)

func (N *NodoEntrada) LlenarMatriz(ListaEntrada []NodoEntrada){
	MatrizIns := Matriz{}
	nuevaMatriz := MatrizIns.NuevaMatriz(0,0)
	var mesanio []string
	for i := 0; i < len(ListaEntrada); i++ {
		fecha := ListaEntrada[i].Fecha
		fechaActual := strings.Split(fecha, "-")
		mes, _ := strconv.Atoi(fechaActual[1])
		anio, _ := strconv.Atoi(fechaActual[2])
		tienda := ListaEntrada[i].Tienda
		Departamento := ListaEntrada[i].Departamento
		Cal := ListaEntrada[i].Calificacion
		Codigo := ListaEntrada[i].ProductoCodigo
		existencia := Existe(mesanio, fechaActual[1]+""+fechaActual[2])
		if existencia == true {
			for existencia == true {
				fecha = ListaEntrada[i].Fecha
				fechaActual = strings.Split(fecha, "-")
				mes, _ = strconv.Atoi(fechaActual[1])
				anio, _ = strconv.Atoi(fechaActual[2])
				tienda = ListaEntrada[i].Tienda
				Departamento = ListaEntrada[i].Departamento
				Cal = ListaEntrada[i].Calificacion
				Codigo = ListaEntrada[i].ProductoCodigo
				existencia = Existe(mesanio, fechaActual[1]+""+fechaActual[2])
				if existencia == true {
					nodoMatriz := MatrizIns.NuevoNodoPedido(fecha, tienda,Departamento,Cal, Codigo)
					nuevaMatriz.Insertar(nodoMatriz)
					i++
					if i == len(ListaEntrada) {
						nuevaMatriz.Imprimir()
						nuevaMatriz.Imprimir2()
						fmt.Println("-----------------------------------------METER-----------------------------------------")
						break
					}
				}
			}
		}
		if existencia == false{
			if nuevaMatriz.Mes != 0 {
				nuevaMatriz.Imprimir()
				nuevaMatriz.Imprimir2()
				fmt.Println("-----------------------------------------METER-----------------------------------------")
			}
			nuevaMatriz = MatrizIns.NuevaMatriz(mes, anio)
			nodoMatriz := MatrizIns.NuevoNodoPedido(fecha, tienda,Departamento,Cal, Codigo)
			nuevaMatriz.Insertar(nodoMatriz)
			mesanio = append(mesanio, fechaActual[1]+""+fechaActual[2])
			if i == len(ListaEntrada)-1 {
				nuevaMatriz.Imprimir()
				nuevaMatriz.Imprimir2()
				fmt.Println("-----------------------------------------METER-----------------------------------------")
			}
		}
		if i != 0 {
			fmt.Println("*******************************************************lista anual*******************************************************")
		}
	}
}

func Existe(arreglo []string, mesanio string) bool{
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == mesanio{
			return true
		}
	}
	return false
}