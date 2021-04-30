package MatrizDispersa

import (
	"../Grafo"
	"../Usuarios"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
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


type GeneralInfo struct{
	CabeceraDepa []string `json:"CabeceraDepartamentos"`
	CabeceraDia []string `json:"CabeceraDias"`
	Datos []Dato `json:"Datos"`
}

type  Dato struct{
	Tienda string `json:"Tienda"`
	Depa string `json:"Departamento"`
	Cal int `json:"Calificacion"`
	Dpi string `json:"Cliente"`
	Nombre string `json:"Nombre"`
	Correo string `json:"Correo"`
	Prod string `json:"NombreProducto"`
	CodProd int `json:"CodigoProducto"`
	Cant int `json:"Cantidad"`
}

func (M *Matriz) Imprimir(dia string, Arbol *Usuarios.ArbolB) GeneralInfo{
	var cabDepa []string
	var cabDia []string
	var nodosDatos []Dato

	var aux1 interface{} = M.CabeceraV
	for aux1!=nil{
	    cabDepa = append(cabDepa, aux1.(*NodoCabeceraVertical).Departamento)
		aux1 = aux1.(*NodoCabeceraVertical).Sur
	}
	var temp2 interface{} = M.CabeceraH
	for temp2 != nil{
		cabDia = append(cabDia, strconv.Itoa(temp2.(*NodoCabeceraHorizontal).Dia))
		temp2 = temp2.(*NodoCabeceraHorizontal).Este
	}

	var aux interface{} = M.CabeceraV
	for aux != nil{
		var a Dato
			temp := aux.(*NodoCabeceraVertical).Este
			for temp != nil{
				temporal2 := strings.Split(temp.(*NodoPedido).Fecha, "-")
				if temporal2[0] == dia || temporal2[0] == "0"+dia{
					datos := Arbol.DatosUsuario(Arbol.Raiz, temp.(*NodoPedido).Cliente)
					if datos != nil {
						a = Dato{Tienda: temp.(*NodoPedido).NombreTienda, Depa: temp.(*NodoPedido).Departamento, Prod: temp.(*NodoPedido).NombreProducto, Cal: temp.(*NodoPedido).Calificacion, CodProd: temp.(*NodoPedido).CodigoProducto, Cant: temp.(*NodoPedido).Cantidad, Dpi: strconv.Itoa(temp.(*NodoPedido).Cliente), Nombre: datos.Nombre, Correo: datos.Correo}
					}else{
						a = Dato{Tienda: temp.(*NodoPedido).NombreTienda, Depa: temp.(*NodoPedido).Departamento, Prod: temp.(*NodoPedido).NombreProducto, Cal: temp.(*NodoPedido).Calificacion, CodProd: temp.(*NodoPedido).CodigoProducto, Cant: temp.(*NodoPedido).Cantidad, Dpi: "Anónimo", Nombre: "Anónimo", Correo: "Anónimo"}
					}
					nodosDatos = append(nodosDatos, a)
					}
				temp = temp.(*NodoPedido).Este
			}
		aux = aux.(*NodoCabeceraVertical).Sur
	}
	linkReg := GeneralInfo{CabeceraDepa: cabDepa, CabeceraDia: cabDia, Datos: nodosDatos}
	return linkReg
}

func (M *Matriz) Recorrido(dia string) *GrafoRecorrido.ListaRecorrido{
	RecorridoRegresa := &GrafoRecorrido.ListaRecorrido{}
	var aux interface{} = M.CabeceraH
	for aux != nil{
		temp := aux.(*NodoCabeceraHorizontal).Sur
		for temp != nil{
			temporal2 := strings.Split(temp.(*NodoPedido).Fecha, "-")
			if temporal2[0] == dia || temporal2[0] == "0"+dia{
				recorrido := temp.(*NodoPedido).Recorrido
				imp := recorrido.Cabeza
				if RecorridoRegresa.Cabeza == nil{
					RecorridoRegresa.InsertarRecCabeza(imp)
				}else{
					Voltear := &GrafoRecorrido.ListaRecorrido{}
					if imp != nil{
						nuevo := &GrafoRecorrido.NodoRecorrido{Viene: imp.Viene, Va: imp.Va, Costo: imp.Costo, Siguiente: nil, Anterior: nil}
						Voltear.InsertarRecCabeza(nuevo)
						break
					}
					impv := Voltear.Cabeza
					for impv != nil{
						nuevoVolteado := &GrafoRecorrido.NodoRecorrido{Viene: impv.Viene, Va: impv.Va, Costo: impv.Costo, Siguiente: nil, Anterior: nil}
						RecorridoRegresa.InsertarRecCabeza(nuevoVolteado)
						impv = impv.Siguiente
					}
				}
			}
			temp = temp.(*NodoPedido).Sur
		}
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
	return RecorridoRegresa
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
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
}

func (this *Matriz) NuevaMatriz(mes int, anio int) *Matriz{
	return &Matriz{mes, anio, nil, nil}
}

func (this *Matriz) NuevoNodoPedido (fecha string, tienda string, departamento string, calificacion int, cliente int, nombreProducto string, codigoProducto int, cantidad int, dia string, recorrido *GrafoRecorrido.ListaRecorrido) *NodoPedido{
	return &NodoPedido{Norte: nil, Oeste: nil, Sur: nil, Este: nil, Fecha: fecha, NombreTienda: tienda, Departamento: departamento, Calificacion: calificacion, Cliente: cliente, NombreProducto: nombreProducto, CodigoProducto: codigoProducto, Cantidad: cantidad, Dia: dia, Recorrido: recorrido}
}

func (M *Matriz) DibujarMatriz(){
	colores := [10]string{"gray", "red", "blue", "yellow", "green", "orange", "brown", "pink", "violet", "purple"}
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph Daniel"+strconv.Itoa(M.Anio)+""+strconv.Itoa(M.Mes)+"{\n")
	fmt.Fprintf(&cadena, "node[shape=box];\n")
	var mes string
	var meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	mes = meses[M.Mes-1]
	fmt.Fprintf(&cadena, "MT[label=\"%v\", style = filled, color="+colores[rand.Intn(9)]+", group = 1];\n", mes)
	fmt.Fprintf(&cadena, "e0[shape = point, width = 0];\n")
	fmt.Fprintf(&cadena, "e0[shape = point, width = 0];\n")
	//RECORRIENDO SOLO LA CABECERA VERTICAL
	var auxVertical interface{} = M.CabeceraV
	for auxVertical != nil{
		fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"Depa: %v\", group = 1];\n", &(auxVertical.(*NodoCabeceraVertical).Departamento), auxVertical.(*NodoCabeceraVertical).Departamento )
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

	//RECORRIENDO SOLO LA CABECERA HORIZONTAL
	var auxHorizontal interface{} = M.CabeceraH
	var contador = 2
	var s strings.Builder
	fmt.Fprintf(&s, "{rank = same;MT;")
	for auxHorizontal != nil{
		fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"Día: %v\", group = "+strconv.Itoa(contador)+"];\n", &(auxHorizontal.(*NodoCabeceraHorizontal).Dia), auxHorizontal.(*NodoCabeceraHorizontal).Dia )
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


	//RECORRIENDO CABECERA HORIZONTAL CON PEDIDOS
	var aux interface{} = M.CabeceraH
	var nodos []int
	contador = 2
	for aux != nil{
		temp := aux.(*NodoCabeceraHorizontal).Sur
		fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"%v\", group = "+strconv.Itoa(contador)+"];\n", &(temp.(*NodoPedido).CodigoProducto), "PEDIDOS" )
		fmt.Fprintf(&cadena, "node%p->node%p;\n", &(aux.(*NodoCabeceraHorizontal).Dia), &(temp.(*NodoPedido).CodigoProducto))
		var a = &(temp.(*NodoPedido).CodigoProducto)
		var casi = false
		for temp != nil{
			if casi == false {
				a = &(temp.(*NodoPedido).CodigoProducto)
				nodos = append(nodos, *a)
				casi = true
			}
			if temp.(*NodoPedido).Sur != nil {
				if temp.(*NodoPedido).Sur.(*NodoPedido).Departamento != temp.(*NodoPedido).Departamento{
					fmt.Fprintf(&cadena, "node%v[color="+colores[rand.Intn(9)]+", label=\"%v\", group = "+strconv.Itoa(contador)+"];\n", &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto), "PEDIDOS" )
					b := &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto)
					nodos = append(nodos, *b)
					fmt.Fprintf(&cadena, "node%p->node%p;\n", a, &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto))
					fmt.Fprintf(&cadena, "node%p->node%p;\n", &(temp.(*NodoPedido).Sur.(*NodoPedido).CodigoProducto), a)
				}
			}
			temp = temp.(*NodoPedido).Sur
		}
		contador++
		aux = aux.(*NodoCabeceraHorizontal).Este
	}

	//RECORRIENDO CABECERA VERTICAL CON NODOS PEDIDOS
	var auxV interface{} = M.CabeceraV
	var GeneralNodos []int
	for auxV != nil{
		var d strings.Builder
		fmt.Fprintf(&d, "{rank = same;node%v;", &(auxV.(*NodoCabeceraVertical).Departamento))
		temp := auxV.(*NodoCabeceraVertical).Este
			for temp != nil{
				busqueda := Posicion(nodos, temp.(*NodoPedido).CodigoProducto)
				if busqueda == true{
					busquedaGeneral := Posicion(GeneralNodos, temp.(*NodoPedido).CodigoProducto)
					if busquedaGeneral == false{
						fmt.Fprintf(&d, "node%v;",&(temp.(*NodoPedido).CodigoProducto))
						GeneralNodos = append(GeneralNodos, temp.(*NodoPedido).CodigoProducto)
						if temp.(*NodoPedido).Oeste == auxV.(*NodoCabeceraVertical) {
							fmt.Fprintf(&cadena, "node%p->node%p;\n", &(auxV.(*NodoCabeceraVertical).Departamento), &(temp.(*NodoPedido).CodigoProducto))
						}
					}
				}
				temp = temp.(*NodoPedido).Este
		}
		fmt.Fprintf(&d, "}")
		fmt.Fprintf(&cadena, d.String()+"\n")
		auxV = auxV.(*NodoCabeceraVertical).Sur
	}

	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), "Matriz")
}

func guardarArchivo(cadena string, nombreArchivo string) {
	f, err := os.Create("..\\cliente\\src\\ImagenMatriz\\"+nombreArchivo+".dot")
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
	cmd, _ := exec.Command(path, "-Tpng", "..\\cliente\\src\\ImagenMatriz\\"+nombreArchivo+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile("..\\cliente\\src\\ImagenMatriz\\"+nombreArchivo+".png", cmd, os.FileMode(mode))
}

func Posicion(arreglo []int, busqueda int) bool {
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == busqueda {
			return true
		}
	}
	return false
}