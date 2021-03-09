package MatrizDispersa

import (
	"strconv"
	"strings"
)

func (N *NodoEntrada) LlenarMatriz(ListaEntrada []NodoEntrada) *ListaAnio{
	MatrizIns := Matriz{}
	nuevaMatriz := MatrizIns.NuevaMatriz(0,0)
	linkMes := &ListaMes{}
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
						nodomes := NodoMes{Mes: nuevaMatriz.Mes, MatrizMes: nuevaMatriz}
						linkMes.Insertar(&nodomes)
						break
					}
				}
			}
		}
		if existencia == false{
			if nuevaMatriz.Mes != 0 {
				nodomes := NodoMes{Mes: nuevaMatriz.Mes, MatrizMes: nuevaMatriz}
				linkMes.Insertar(&nodomes)
			}
			nuevaMatriz = MatrizIns.NuevaMatriz(mes, anio)
			nodoMatriz := MatrizIns.NuevoNodoPedido(fecha, tienda,Departamento,Cal, Codigo)
			nuevaMatriz.Insertar(nodoMatriz)
			mesanio = append(mesanio, fechaActual[1]+""+fechaActual[2])
			if i == len(ListaEntrada)-1 {
				nodomes := NodoMes{Mes: nuevaMatriz.Mes, MatrizMes: nuevaMatriz}
				linkMes.Insertar(&nodomes)
			}
		}
	}
	var anios []int
	var mes []int
	if linkMes.Cabeza != nil {
		imp := linkMes.Cabeza
		for imp != nil {
			existeAnio := ExisteAnio(anios, imp.MatrizMes.Anio)
			existeMes := ExisteAnio(mes, imp.MatrizMes.Mes)
			if existeAnio == false {
				anios = append(anios, imp.MatrizMes.Anio)
			}
			if existeMes == false {
				mes = append(mes, imp.MatrizMes.Mes)
			}
			imp = imp.Siguiente
		}
	}
	anios = *burbuja(anios)
	mes = *burbuja(mes)
	listaAnios := ListaAnio{}
	for i := 0; i < len(anios); i++ {
		linkIng := &ListaMes{}
		for j := 0; j < len(mes); j++ {
			imp := linkMes.Cabeza
			for imp != nil{
				if imp.MatrizMes.Mes == mes[j] && imp.MatrizMes.Anio == anios[i]{
					nodomes := NodoMes{Mes: imp.MatrizMes.Mes, MatrizMes: imp.MatrizMes}
					linkIng.Insertar(&nodomes)
				}
				imp = imp.Siguiente
			}
		}
		nodoanio := NodoAnio{Anio: anios[i], ListaMatricesMes: linkIng}
		listaAnios.Insertar(&nodoanio)
	}
	return &listaAnios
}

func Existe(arreglo []string, mesanio string) bool{
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == mesanio{
			return true
		}
	}
	return false
}

func ExisteAnio(arreglo []int, mesanio int) bool{
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == mesanio{
			return true
		}
	}
	return false
}

func burbuja(listaNodos []int) *[]int{
	var i,j int
	var aux int
	for i = 0; i < len(listaNodos)-1; i++ {
		for j = 0; j < len(listaNodos)-i-1 ; j++ {
			siguiente := listaNodos[j+1]
			anterior := listaNodos[j]
			if siguiente < anterior{
				aux = listaNodos[j+1]
				listaNodos[j+1] = listaNodos[j]
				listaNodos[j] = aux
			}
		}
	}
	return &listaNodos
}
