package Compras

import (
	"../Inventario"
)

type General struct {
	Pedidos []Pedidos `json:"Compras"`
}

type Pedidos struct {
	Fecha string `json:"Fecha"`
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Cliente int `json:"Cliente"`
	CodigoProductos []NodoProducto `json:"Productos"`
}

type NodoProducto struct{
	Codigo int `json:"Codigo"`
	Cantidad int `json:"Cantidad"`
}

func DescontarProducto(ra *Inventario.NodoArbol, codigo int, cantidad int) *Inventario.NodoArbol{
	if codigo<ra.Codigo{
		izq := DescontarProducto(ra.Izq,codigo, cantidad)
		ra.Izq = izq
	}else if codigo > ra.Codigo{
		der := DescontarProducto(ra.Der, codigo, cantidad)
		ra.Der = der
	}else if codigo == ra.Codigo{
		if ra.Cantidad > 0 {
			ra.Cantidad = ra.Cantidad-cantidad
		}
	}
	return ra
}