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
	Departamento string
	Norte interface{}
	Sur interface{}
	Este interface{}
	Oeste interface{}
}

type NodoCabeceraHorizontal struct{
	Dia int
	Norte interface{}
	Sur interface{}
	Este interface{}
	Oeste interface{}
}

type NodoPedido struct{
	NombreTienda string
	Departamento string
	Calificacion int
	CodigoProducto int
	Norte interface{}
	Sur interface{}
	Este interface{}
	Oeste interface{}
	Fecha string
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

func (M *Matriz) BurbujaMes(listaMeses ListaMes) *ListaMes {
	var meses []NodoMes
	imp := listaMeses.Cabeza
	for imp != nil{
		nodo := NodoMes{Mes: imp.Mes, MatrizMes: imp.MatrizMes}
		meses = append(meses, nodo)
		imp = imp.Siguiente
	}
	listaMeses = *burbuja(meses)
	return &listaMeses
}

func burbuja(listaNodos []NodoMes) *ListaMes{
	linkGA := &ListaMes{}
	var i,j int
	var aux NodoMes
	for i = 0; i < len(listaNodos)-1; i++ {
		for j = 0; j < len(listaNodos)-i-1 ; j++ {
			siguiente := listaNodos[j+1]
			anterior := listaNodos[j]
			if siguiente.Mes < anterior.Mes{
				aux = listaNodos[j+1]
				listaNodos[j+1] = listaNodos[j]
				listaNodos[j] = aux
			}
		}
	}
	for k := 0; k < len(listaNodos); k++ {
		linkGA.Insertar(&listaNodos[k])
	}
	return linkGA
}

func (M *Matriz) BurbujaAnio(listaMeses ListaAnio) *ListaAnio {
	var anios []NodoAnio
	imp := listaMeses.Cabeza
	for imp != nil{
		nodo := NodoAnio{Anio: imp.Anio, ListaMatricesMes: imp.ListaMatricesMes}
		anios = append(anios, nodo)
		imp = imp.Siguiente
	}
	listaMeses = *burbujaA(anios)
	return &listaMeses
}

func burbujaA(listaNodos []NodoAnio) *ListaAnio{
	linkGA := &ListaAnio{}
	var i,j int
	var aux NodoAnio
	for i = 0; i < len(listaNodos)-1; i++ {
		for j = 0; j < len(listaNodos)-i-1 ; j++ {
			siguiente := listaNodos[j+1]
			anterior := listaNodos[j]
			if siguiente.Anio < anterior.Anio{
				aux = listaNodos[j+1]
				listaNodos[j+1] = listaNodos[j]
				listaNodos[j] = aux
			}
		}
	}
	for k := 0; k < len(listaNodos); k++ {
		linkGA.Insertar(&listaNodos[k])
	}
	return linkGA
}