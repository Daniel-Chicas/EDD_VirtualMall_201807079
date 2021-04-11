package Usuarios

type Pagina struct{
	Maximo int
	NodoPadre *Pagina
	Llaves []*Llave
}

func NuevoNodo(max int) *Pagina {
	llaves :=make([]*Llave, max)
	nodo := Pagina{max, nil, llaves}
	return &nodo
}

func (N *Pagina) Colocar(i int, llave *Llave){
	N.Llaves[i] = llave
}