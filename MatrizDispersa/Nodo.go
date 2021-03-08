package MatrizDispersa

type General struct {
	Pedidos []Pedidos `json:"Pedidos"`
}

type Pedidos struct {
	Fecha string `json:"Fecha"`
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []NodoProductoViene `json:"Productos"`
}

type NodoProductoViene struct {
	Codigo int `json:"Codigo"`
}

type NodoCabeceraVertical struct{
	Este interface{}
	Norte interface{}
	Sur interface{}
	Oeste interface{}
	Departamento string
}

type NodoCabeceraHorizontal struct{
	Este interface{}
	Norte interface{}
	Sur interface{}
	Oeste interface{}
	Dia int
}

type NodoPedido struct{
	Este interface{}
	Norte interface{}
	Sur interface{}
	Oeste interface{}
	Fecha string
	NombreTienda string
	Departamento string
	Calificacion int
	CodigoProducto int
}

type NodoMes struct{
	Mes int
	MatrizMes *Matriz
	Siguiente *NodoMes
	Anterior *NodoMes
}

type ListaMes struct {
	Cabeza *NodoMes
	Cola *NodoMes
}

func (L *ListaMes) Insertar(nuevo *NodoMes){
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
}

type NodoAnio struct{
	Anio int
	ListaMatricesMes *ListaMes
	Siguiente *NodoAnio
	Anterior *NodoAnio
}

type ListaAnio struct {
	Cabeza *NodoAnio
	Cola *NodoAnio
}

func (L *ListaAnio) Insertar(nuevo *NodoAnio){
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
}

type NodoEntrada struct {
	Fecha string
	Tienda string
	Departamento string
	Calificacion int
	ProductoCodigo int
}
