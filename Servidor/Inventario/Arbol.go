package Inventario

import (
	"../Comentarios"
)

type General struct {
	Inventarios []Inventario `json:"Inventarios"`
}

type Inventario struct {
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []NodoProducto `json:"Productos"`
}

type NodoProducto struct {
	NombreProducto string `json:"Nombre"`
	Codigo int `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	PrecioP float64 `json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string `json:"Imagen"`
	Almacenamiento string `json:"Almacenamiento"`
}

type NodoArbol struct{
	NombreProducto string
	Codigo int
	Descripcion string
	Precio float64
	Cantidad int
	Imagen string
	Factor int
	Almacenamiento string
	Comentarios Comentarios.TablaHash
	Izq *NodoArbol
	Der *NodoArbol
}


type Arbol struct{
	Raiz *NodoArbol
}

func (this *Arbol) NuevoArbol() *Arbol{
	return &Arbol{nil}
}

func NuevoNodo (nombre string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string) *NodoArbol{
	tabla := Comentarios.NuevaTabla(7, 50, 20)
	return &NodoArbol{nombre, codigo, descripcion, precio, cantidad, imagen, 0, almacenamiento, *tabla,nil, nil}
}

func RotacionII(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n.Izq = n1.Der
	n1.Der = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func RotacionDD(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n.Der = n1.Izq
	n1.Izq = n
	if n1.Factor == 1{
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func RotacionDI(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n2 := n1.Izq
	n.Der = n2.Izq
	n2.Izq = n
	n1.Izq = n2.Der
	n2.Der = n1
	if n2.Factor == 1 {
		n.Factor = -1
	}else{
		n.Factor = 0
	}
	if n2.Factor == -1{
		n1.Factor = 1
	}else{
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func RotacionID(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n2 := n1.Der
	n.Izq = n2.Der
	n2.Der = n
	n1.Der = n2.Izq
	n2.Izq = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	}else{
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	}else{
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func insertar(ra *NodoArbol, nombre string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string, ya *bool) *NodoArbol{
	var n1 *NodoArbol
	if ra == nil {
		ra = NuevoNodo(nombre, codigo, descripcion, precio, cantidad, imagen, almacenamiento)
		*ya = true
	}else if codigo<ra.Codigo{
		izq := insertar(ra.Izq, nombre, codigo, descripcion, precio, cantidad, imagen, almacenamiento, ya)
		ra.Izq = izq
		if *ya == true {
			switch ra.Factor {
			case 1:
				ra.Factor = 0
				*ya = false
				break
			case 0:
				ra.Factor = -1
				break
			case -1:
				n1 = ra.Izq
				if n1.Factor==-1 {
					ra = RotacionII(ra, n1)
				}else{
					ra = RotacionID(ra, n1)
				}
				*ya = false
			}
		}
	}else if codigo > ra.Codigo{
		der := insertar(ra.Der, nombre, codigo, descripcion, precio, cantidad, imagen, almacenamiento, ya)
		ra.Der = der
		if *ya{
			switch ra.Factor {
			case 1:
				n1 = ra.Der
				if n1.Factor == 1 {
					ra = RotacionDD(ra, n1)
				}else{
					ra = RotacionDI(ra, n1)
				}

				*ya = false
				break
			case 0:
				ra.Factor = 1
				break
			case -1:
				ra.Factor = 0
				*ya = false
			}
		}
	}else if codigo == ra.Codigo{
		ra.Cantidad = ra.Cantidad+cantidad
	}
	return ra
}

func (L *Arbol) Insertar(nombre string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string){
	b := false
	a := &b
	L.Raiz = insertar(L.Raiz, nombre, codigo, descripcion, precio, cantidad, imagen, almacenamiento, a)
}