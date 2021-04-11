package Usuarios

import "strconv"

type ArbolB struct{
	k int
	Raiz *Pagina
}

func (A *ArbolB) NuevoArbol (nivel int) *ArbolB{
	a:=ArbolB{nivel, nil}
	nodoRaiz := NuevoNodo(nivel)
	a.Raiz=nodoRaiz
	return &a
}

func (A *ArbolB) Insertar(nuevaLlave *Llave){
	if A.Raiz.Llaves[0] == nil{
		A.Raiz.Colocar(0, nuevaLlave)
	}else if A.Raiz.Llaves[0].Izq == nil{
		lugarins :=  -1
		nodoAux := A.Raiz
		lugarins = A.colocarNodo(nodoAux, nuevaLlave)
		if lugarins != -1 {
			if lugarins == nodoAux.Maximo-1 {
				medio := nodoAux.Maximo/2
				llaveCentro := nodoAux.Llaves[medio]
				derecha := NuevoNodo(A.k)
				izquierda := NuevoNodo(A.k)
				indiceIzq := 0
				indiceDer :=0
				for i := 0; i < nodoAux.Maximo; i++ {
					if nodoAux.Llaves[i].Usuario.DPI < llaveCentro.Usuario.DPI {
						izquierda.Colocar(indiceIzq, nodoAux.Llaves[i])
						indiceIzq++
						nodoAux.Colocar(i, nil)
					}else if nodoAux.Llaves[i].Usuario.DPI > llaveCentro.Usuario.DPI{
						derecha.Colocar(indiceDer, nodoAux.Llaves[i])
						indiceDer++
						nodoAux.Colocar(i, nil)
					}
				}
				nodoAux.Colocar(medio, nil)
				A.Raiz = nodoAux
				A.Raiz.Colocar(0, llaveCentro)
				izquierda.NodoPadre = A.Raiz
				derecha.NodoPadre = A.Raiz
				llaveCentro.Izq = izquierda
				llaveCentro.Der = derecha
			}
		}
	}else if A.Raiz.Llaves[0].Izq != nil{
		nodoAux := A.Raiz
		for nodoAux.Llaves[0].Izq != nil{
			contador := 0
			for i := 0; i < nodoAux.Maximo; i, contador = i+1, contador+1 {
				if nodoAux.Llaves[i] != nil {
					if nodoAux.Llaves[i].Usuario.DPI > nuevaLlave.Usuario.DPI {
						nodoAux = nodoAux.Llaves[i].Izq
						break
					}
				}else{
					nodoAux = nodoAux.Llaves[i-1].Der
					break
				}
			}
			if contador==nodoAux.Maximo{
				nodoAux = nodoAux.Llaves[contador-1].Der
			}
		}
		indiceCol := A.colocarNodo(nodoAux, nuevaLlave)
		if indiceCol == nodoAux.Maximo-1 {
			for nodoAux.NodoPadre != nil{
				medio := nodoAux.Maximo/2
				centro := nodoAux.Llaves[medio]
				izq := NuevoNodo(A.k)
				der := NuevoNodo(A.k)
				indiceIzq := 0
				indiceDer := 0
				for i := 0; i < nodoAux.Maximo; i++ {
					if nodoAux.Llaves[i].Usuario.DPI < centro.Usuario.DPI {
						izq.Colocar(indiceIzq, nodoAux.Llaves[i])
						indiceIzq++
						nodoAux.Colocar(i, nil)
					}else if nodoAux.Llaves[i].Usuario.DPI > centro.Usuario.DPI{
						der.Colocar(indiceDer, nodoAux.Llaves[i])
						indiceDer++
						nodoAux.Colocar(i, nil)
					}
				}
				nodoAux.Colocar(medio, nil)
				centro.Izq = izq
				centro.Der = der
				nodoAux = nodoAux.NodoPadre
				izq.NodoPadre = nodoAux
				der.NodoPadre = nodoAux
				for i := 0; i < izq.Maximo; i++ {
					if izq.Llaves[i] != nil {
						if izq.Llaves[i].Izq != nil {
							izq.Llaves[i].Izq.NodoPadre = izq
						}
						if izq.Llaves[i].Der != nil{
							izq.Llaves[i].Der.NodoPadre = izq
						}
					}
				}
				for i := 0; i < der.Maximo; i++ {
					if der.Llaves[i] != nil {
						if der.Llaves[i].Izq != nil{
							der.Llaves[i].Izq.NodoPadre = der
						}
						if der.Llaves[i].Der != nil {
							der.Llaves[i].Der.NodoPadre = der
						}
					}
				}
				colocado := A.colocarNodo(nodoAux, centro)
				if colocado == nodoAux.Maximo-1 {
					if nodoAux.NodoPadre == nil {
						centralRaiz := nodoAux.Maximo/2
						llaveCentral :=  nodoAux.Llaves[centralRaiz]
						izqRaiz := NuevoNodo(A.k)
						derRaiz := NuevoNodo(A.k)
						indiceIzqRaiz := 0
						indiceDerRaiz := 0
						for i := 0; i < nodoAux.Maximo; i++ {
							if nodoAux.Llaves[i].Usuario.DPI < llaveCentral.Usuario.DPI{
								izqRaiz.Colocar(indiceIzqRaiz, nodoAux.Llaves[i])
								indiceIzqRaiz++
								nodoAux.Colocar(i, nil)
							}else if nodoAux.Llaves[i].Usuario.DPI > llaveCentral.Usuario.DPI{
								derRaiz.Colocar(indiceDerRaiz, nodoAux.Llaves[i])
								indiceDerRaiz++
								nodoAux.Colocar(i, nil)
							}
						}
						nodoAux.Colocar(centralRaiz, nil)
						nodoAux.Colocar(0, llaveCentral)
						for i := 0; i < A.k; i++ {
							if izqRaiz.Llaves[i] != nil {
								izqRaiz.Llaves[i].Izq.NodoPadre = izqRaiz
								izqRaiz.Llaves[i].Der.NodoPadre = izqRaiz
							}
						}
						for i := 0; i < A.k; i++ {
							if derRaiz.Llaves[i] != nil {
								derRaiz.Llaves[i].Izq.NodoPadre = derRaiz
								derRaiz.Llaves[i].Der.NodoPadre = derRaiz
							}
						}
						llaveCentral.Izq = izqRaiz
						llaveCentral.Der = derRaiz
						izqRaiz.NodoPadre = nodoAux
						derRaiz.NodoPadre = nodoAux
						A.Raiz = nodoAux
					}
					continue
				}else{
					break
				}
			}
		}
	}
}

func (A *ArbolB) colocarNodo(nodo *Pagina, nuevaLlave *Llave) int{
	indice := -1
	for i := 0; i < nodo.Maximo; i++ {
		if nodo.Llaves[i] == nil {
			agregado := false
			for j := i-1; j >= 0; j-- {
				if nodo.Llaves[j].Usuario.DPI > nuevaLlave.Usuario.DPI {
					nodo.Colocar(j+1, nodo.Llaves[j])
				}else{
					nodo.Colocar(j+1, nuevaLlave)
					nodo.Llaves[j].Der =nuevaLlave.Izq
					if j+2 < A.k && nodo.Llaves[j+2] !=nil {
						nodo.Llaves[j+2].Izq = nuevaLlave.Der
					}
					agregado = true
					break
				}
			}
			if agregado == false {
				nodo.Colocar(0, nuevaLlave)
				nodo.Llaves[1].Izq = nuevaLlave.Der
			}
			indice = i
			break
		}
	}
	return indice
}


func (A *ArbolB) DatosUsuario (pagina *Pagina, dpi int) *NodoUsuario{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
						existe = true
						return pagina.Llaves[i].Usuario
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := A.DatosUsuario(pagina.Llaves[i].Izq, dpi)
						if a != nil {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := A.DatosUsuario(pagina.Llaves[i].Der, dpi)
						if a != nil {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := A.DatosUsuario(pagina.Llaves[i].Der, dpi)
						if a != nil {
							return a
						}
					}
				}
			}
		}else{
			return nil
		}
	}
	return nil
}