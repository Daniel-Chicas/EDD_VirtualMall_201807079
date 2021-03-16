package MatrizDispersa

import (
	"fmt"
	"math/rand"
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
			fmt.Printf("%v, %v,%v, %v, %v----->", temp.(*NodoPedido).Departamento, temp.(*NodoPedido).NombreTienda, temp.(*NodoPedido).Fecha, temp.(*NodoPedido).NombreProducto, temp.(*NodoPedido).CodigoProducto)
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
			fmt.Printf("%v, %v, %v, %v, %v----->", temp.(*NodoPedido).Departamento, temp.(*NodoPedido).NombreTienda, temp.(*NodoPedido).Fecha, temp.(*NodoPedido).NombreProducto, temp.(*NodoPedido).CodigoProducto)
			temp = temp.(*NodoPedido).Sur
		}
		fmt.Println("")
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
}

func (this *Matriz) NuevaMatriz(mes int, anio int) *Matriz{
	return &Matriz{mes, anio, nil, nil}
}

func (this *Matriz) NuevoNodoPedido (fecha string, tienda string, departamento string, calificacion int, nombreProducto string, codigoProducto int) *NodoPedido{
	return &NodoPedido{Norte: nil, Oeste: nil, Sur: nil, Este: nil, Fecha: fecha, NombreTienda: tienda, Departamento: departamento, Calificacion: calificacion, NombreProducto: nombreProducto, CodigoProducto: codigoProducto}
}

func (M *Matriz) DibujarMatriz(){
	colores := [10]string{"gray", "red", "blue", "yellow", "green", "orange", "brown", "pink", "violet", "purple"}
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph Daniel"+strconv.Itoa(M.Anio)+""+strconv.Itoa(M.Mes)+"{\n")
	fmt.Fprintf(&cadena, "node[shape=box];\n")
	fmt.Fprintf(&cadena, "MT[label=\"Matriz\", style = filled, color="+colores[rand.Intn(9)]+", group = 1];\n")
	fmt.Fprintf(&cadena, "e0[shape = point, width = 0];\n")
	fmt.Fprintf(&cadena, "e0[shape = point, width = 0];\n")
	var auxVertical interface{} = M.CabeceraV
	for auxVertical != nil{
		fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"%v\", group = 1];\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), auxVertical.(*NodoCabeceraVertical).Departamento )
		if auxVertical.(*NodoCabeceraVertical).Norte == nil {
			fmt.Fprintf(&cadena, "MT->node%p;\n", &(auxVertical.(*NodoCabeceraVertical).Departamento))
			if auxVertical.(*NodoCabeceraVertical).Sur != nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), &(auxVertical.(*NodoCabeceraVertical).Sur.(*NodoCabeceraVertical).Departamento))
			}
		}else if auxVertical.(*NodoCabeceraVertical).Sur != nil{
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), &(auxVertical.(*NodoCabeceraVertical).Norte.(*NodoCabeceraVertical).Departamento))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), &(auxVertical.(*NodoCabeceraVertical).Sur.(*NodoCabeceraVertical).Departamento))
		}else if auxVertical.(*NodoCabeceraVertical).Sur == nil{
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), &(auxVertical.(*NodoCabeceraVertical).Norte.(*NodoCabeceraVertical).Departamento))
		}
		auxVertical = auxVertical.(*NodoCabeceraVertical).Sur
	}
	var auxHorizontal interface{} = M.CabeceraH
	var contador = 2
	var s strings.Builder
	fmt.Fprintf(&s, "{rank = same;MT;")
	for auxHorizontal != nil{
		fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"%v\", group = "+strconv.Itoa(contador)+"];\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), auxHorizontal.(*NodoCabeceraHorizontal).Dia )
		if auxHorizontal.(*NodoCabeceraHorizontal).Oeste == nil {
			fmt.Fprintf(&cadena, "MT->node%p;\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia))
			if auxHorizontal.(*NodoCabeceraHorizontal).Este != nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), &(auxHorizontal.(*NodoCabeceraHorizontal).Este.(*NodoCabeceraHorizontal).Dia))
			}
		}else if auxHorizontal.(*NodoCabeceraHorizontal).Este != nil{
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), &(auxHorizontal.(*NodoCabeceraHorizontal).Oeste.(*NodoCabeceraHorizontal).Dia))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), &(auxHorizontal.(*NodoCabeceraHorizontal).Este.(*NodoCabeceraHorizontal).Dia))
		}else if auxHorizontal.(*NodoCabeceraHorizontal).Este == nil{
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), &(auxHorizontal.(*NodoCabeceraHorizontal).Oeste.(*NodoCabeceraHorizontal).Dia))
		}
		fmt.Fprintf(&s, "node%v;",&(auxHorizontal.(*NodoCabeceraHorizontal).Dia))
		auxHorizontal = auxHorizontal.(*NodoCabeceraHorizontal).Este
		contador++
	}
	fmt.Fprintf(&s, "}")
	fmt.Fprintf(&cadena, s.String()+"\n")

	var auxV interface{} = M.CabeceraV
	for auxV != nil{
		var d strings.Builder
		temp := auxV.(*NodoCabeceraVertical).Este
		fmt.Fprintf(&d, "{rank = same;node%v;", &(auxV.(*NodoCabeceraVertical).Departamento))
		for temp != nil{
			if temp.(*NodoPedido).Oeste == auxV.(*NodoCabeceraVertical) {
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxV.(*NodoCabeceraVertical).Departamento), &(temp.(*NodoPedido).CodigoProducto))
				//if temp.(*NodoPedido).Este != nil {
				//	fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Este.(*NodoPedido).CodigoProducto))
				//}
			}
			/*else if temp.(*NodoPedido).Este != nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Oeste.(*NodoPedido).CodigoProducto))
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Este.(*NodoPedido).CodigoProducto))
			}else if temp.(*NodoPedido).Este == nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Oeste.(*NodoPedido).CodigoProducto))
			}
			 */
			fmt.Fprintf(&d, "node%v;",&(temp.(*NodoPedido).CodigoProducto))
			temp = temp.(*NodoPedido).Este
		}
		fmt.Fprintf(&d, "}")
		fmt.Fprintf(&cadena, d.String()+"\n")
		auxV = auxV.(*NodoCabeceraVertical).Sur
	}

	var aux interface{} = M.CabeceraH
	contador = 2
	for aux != nil{
		temp := aux.(*NodoCabeceraHorizontal).Sur
		for temp != nil{
			fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"%v\", group ="+strconv.Itoa(contador)+"];\n", &(temp.(*NodoPedido).CodigoProducto), temp.(*NodoPedido).NombreProducto )
			if temp.(*NodoPedido).Norte == aux.(*NodoCabeceraHorizontal) {
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(aux.(*NodoCabeceraHorizontal).Dia), &(temp.(*NodoPedido).CodigoProducto))
				if temp.(*NodoPedido).Sur != nil {
					fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto))
				}
			}else if temp.(*NodoPedido).Sur != nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Norte.(*NodoPedido).CodigoProducto))
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto))
			}else if temp.(*NodoPedido).Sur == nil{
				fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).CodigoProducto), &(temp.(*NodoPedido).Norte.(*NodoPedido).CodigoProducto))
			}
			temp = temp.(*NodoPedido).Sur
		}
		contador++
		aux = aux.(*NodoCabeceraHorizontal).Este
	}



	fmt.Fprintf(&cadena, "}\n")
	fmt.Println(cadena.String())
}