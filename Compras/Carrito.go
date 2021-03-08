package Compras

type General struct {
	Pedidos []Pedidos `json:"Pedidos"`
}

type Pedidos struct {
	Fecha string `json:"Fecha"`
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []NodoProducto `json:"Productos"`
}

type NodoProducto struct {
	Codigo int `json:"Codigo"`
}

type NodoArbol struct{
	Codigo int
}