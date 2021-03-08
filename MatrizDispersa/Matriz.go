package MatrizDispersa

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Matriz struct{
	Mes int
	Anio int
	CabeceraH *NodoCabeceraHorizontal
	CabeceraV *NodoCabeceraVertical
}

func (MV *Matriz) obtenerVertical(departamento string) interface{}{
	if MV.CabeceraV == nil{
		return nil
	}
	var aux interface{} = MV.CabeceraV
	for aux!=nil{
		if aux.(*NodoCabeceraVertical).Departamento == departamento {
			return aux
		}
		aux = aux.(*NodoCabeceraVertical).Sur
	}
	return nil
}

func (MH *Matriz) obtenerHorizontal (dia int) interface{}{
	if MH.CabeceraH == nil{
		return nil
	}
	var aux interface{} = MH.CabeceraH
	for aux!=nil{
		if aux.(*NodoCabeceraHorizontal).Dia == dia {
			return aux
		}
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
	return nil
}

func (CV *Matriz) crearVertical (departamento string) *NodoCabeceraVertical{
	if CV.CabeceraV == nil {
		nueva:=&NodoCabeceraVertical{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Departamento: departamento}
		CV.CabeceraV = nueva
		return nueva
	}
	var aux interface{} = CV.CabeceraV
	if departamento[0]<= aux.(*NodoCabeceraVertical).Departamento[0] {
		nueva:=&NodoCabeceraVertical{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Departamento: departamento}
		nueva.Sur = CV.CabeceraV
		CV.CabeceraV.Norte = nueva
		CV.CabeceraV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).Sur!=nil{
		if departamento[0]>aux.(*NodoCabeceraVertical).Departamento[0] && departamento[0]<aux.(*NodoCabeceraVertical).Sur.(*NodoCabeceraVertical).Departamento[0]{
			nueva:=&NodoCabeceraVertical{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Departamento: departamento}
			temporal := aux.(*NodoCabeceraVertical).Sur
			temporal.(*NodoCabeceraVertical).Norte = nueva
			nueva.Sur=temporal
			aux.(*NodoCabeceraVertical).Sur=nueva
			nueva.Norte=aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).Sur
	}
	nueva:=&NodoCabeceraVertical{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Departamento: departamento}
	aux.(*NodoCabeceraVertical).Sur = nueva
	nueva.Norte = aux
	return nueva
}

func (CH *Matriz) crearHorizontal (dia int) *NodoCabeceraHorizontal{
	if CH.CabeceraH == nil {
		nueva:=&NodoCabeceraHorizontal{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Dia: dia}
		CH.CabeceraH = nueva
		return nueva
	}
	var aux interface{} = CH.CabeceraH
	if dia<= aux.(*NodoCabeceraHorizontal).Dia {
		nueva:=&NodoCabeceraHorizontal{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Dia: dia}
		nueva.Este = CH.CabeceraH
		CH.CabeceraH.Oeste = nueva
		CH.CabeceraH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).Este!=nil{
		if dia>aux.(*NodoCabeceraHorizontal).Dia && dia <= aux.(*NodoCabeceraHorizontal).Este.(*NodoCabeceraHorizontal).Dia{
			nueva:=&NodoCabeceraHorizontal{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Dia: dia}
			temporal := aux.(*NodoCabeceraHorizontal).Este
			temporal.(*NodoCabeceraHorizontal).Oeste = nueva
			nueva.Este=temporal
			aux.(*NodoCabeceraHorizontal).Este=nueva
			nueva.Oeste=aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
	nueva:=&NodoCabeceraHorizontal{Este: nil, Oeste: nil, Sur: nil, Norte: nil, Dia: dia}
	aux.(*NodoCabeceraHorizontal).Este = nueva
	nueva.Oeste = aux
	return nueva
}

func (M *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, departamento string) interface{} {
	if cabecera == nil{
		return cabecera
	}
	aux := cabecera.Sur
	if aux == nil {
		return cabecera
	}
	if departamento[0]<=aux.(*NodoPedido).Departamento[0]{
		return cabecera
	}
	for aux.(*NodoPedido).Sur != nil {
		if departamento[0]>aux.(*NodoPedido).Departamento[0] && departamento[0]<= aux.(*NodoPedido).Sur.(*NodoPedido).Departamento[0] {
			return aux
		}
		aux = aux.(*NodoPedido).Sur
	}
	if departamento[0]<=aux.(*NodoPedido).Departamento[0] {
		return aux.(*NodoPedido).Norte
	}
	return aux
}

func (M *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dia int) interface{} {
	if cabecera == nil{
		return cabecera
	}
	aux := cabecera.Este
	var diaFecha int
	if aux == nil {
		return cabecera
	}else{
		diaFecha,_ = strconv.Atoi(strings.Split(aux.(*NodoPedido).Fecha, "-")[0])
	}
	if dia <= diaFecha{
		return cabecera
	}
	for aux.(*NodoPedido).Este != nil {
		diaComp,_ := strconv.Atoi(strings.Split(aux.(*NodoPedido).Fecha, "-")[0])
		diaCompSig,_ := strconv.Atoi(strings.Split(aux.(*NodoPedido).Este.(*NodoPedido).Fecha, "-")[0])
		if dia>diaComp && dia<= diaCompSig {
			return aux
		}
		aux = aux.(*NodoPedido).Este
	}
	if dia<=diaFecha {
		return aux.(*NodoPedido).Oeste
	}
	return aux
}

func (M *Matriz) Insertar(nuevo *NodoPedido) *Matriz{
	vertical:= M.obtenerVertical(nuevo.Departamento)
	fecha := strings.Split(nuevo.Fecha, "-")
	dia,_ := strconv.Atoi(fecha[0])
	horizontal:=M.obtenerHorizontal(dia)
	if vertical==nil {
		vertical=M.crearVertical(nuevo.Departamento)
	}
	if horizontal == nil {
		horizontal=M.crearHorizontal(dia)
	}
	izq := M.obtenerUltimoH(vertical.(*NodoCabeceraVertical), dia)
	sup := M.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nuevo.Departamento)
	if reflect.TypeOf(izq).String() == "*MatrizDispersa.NodoPedido" {
		if izq.(*NodoPedido).Este == nil {
			izq.(*NodoPedido).Este=nuevo
			nuevo.Oeste = izq
		}else{
			temp:=izq.(*NodoPedido).Este
			izq.(*NodoPedido).Este = nuevo
			nuevo.Oeste = izq
			temp.(*NodoPedido).Oeste = nuevo
			nuevo.Este = temp
		}
	}else{
		if izq.(*NodoCabeceraVertical).Este == nil {
			izq.(*NodoCabeceraVertical).Este = 	nuevo
			nuevo.Oeste = izq
		}else{
			temp:= izq.(*NodoCabeceraVertical).Este
			izq.(*NodoCabeceraVertical).Este = nuevo
			nuevo.Oeste = izq
			temp.(*NodoPedido).Oeste = nuevo
			nuevo.Este = temp
		}
	}

	if reflect.TypeOf(sup).String() == "*MatrizDispersa.NodoPedido" {
		if sup.(*NodoPedido).Sur == nil {
			sup.(*NodoPedido).Sur=nuevo
			nuevo.Norte = sup
		}else{
			temp:=sup.(*NodoPedido).Sur
			sup.(*NodoPedido).Sur = nuevo
			nuevo.Norte = sup
			temp.(*NodoPedido).Norte = nuevo
			nuevo.Sur = temp
		}
	}else{
		if sup.(*NodoCabeceraHorizontal).Sur == nil {
			sup.(*NodoCabeceraHorizontal).Sur = nuevo
			nuevo.Norte = sup
		}else{
			temp:= sup.(*NodoCabeceraHorizontal).Sur
			sup.(*NodoCabeceraHorizontal).Sur = nuevo
			nuevo.Norte = sup
			temp.(*NodoPedido).Norte = nuevo
			nuevo.Sur = temp
		}
	}
	return &Matriz{}
}

func (M *Matriz) Imprimir(){
	var aux interface{} = M.CabeceraV
	for aux != nil{
		fmt.Print(aux.(*NodoCabeceraVertical).Departamento, "---------------->")
		temp := aux.(*NodoCabeceraVertical).Este
		for temp != nil{
			fmt.Printf("%v, %v, %v, %v----->", temp.(*NodoPedido).Departamento, temp.(*NodoPedido).NombreTienda, temp.(*NodoPedido).Fecha, temp.(*NodoPedido).CodigoProducto)
			temp = temp.(*NodoPedido).Este
		}
		fmt.Println("\n")
		aux = aux.(*NodoCabeceraVertical).Sur
	}
}

func (M *Matriz) Imprimir2(){
	var aux interface{} = M.CabeceraH
	for aux != nil{
		fmt.Print(aux.(*NodoCabeceraHorizontal).Dia, "---------------->")
		temp := aux.(*NodoCabeceraHorizontal).Sur
		for temp != nil{
			fmt.Printf("%v, %v, %v, %v----->", temp.(*NodoPedido).Departamento, temp.(*NodoPedido).NombreTienda, temp.(*NodoPedido).Fecha, temp.(*NodoPedido).CodigoProducto)
			temp = temp.(*NodoPedido).Sur
		}
		fmt.Println("\n")
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
}

func (this *Matriz) NuevaMatriz(mes int, anio int) *Matriz{
	return &Matriz{mes, anio, nil, nil}
}

func (this *Matriz) NuevoNodoPedido (fecha string, tienda string, departamento string, calificacion int, codigoProducto int) *NodoPedido{
	return &NodoPedido{Norte: nil, Oeste: nil, Sur: nil, Este: nil, Fecha: fecha, NombreTienda: tienda, Departamento: departamento, Calificacion: calificacion, CodigoProducto: codigoProducto}
}