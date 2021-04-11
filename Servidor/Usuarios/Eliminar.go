package Usuarios

import "strconv"

func (A *ArbolB) ExisteBEliminar (pagina *Pagina, dpi int, contrasenia string) bool{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) && pagina.Llaves[i].Usuario.Contra == contrasenia{
					pagina = A.usuarioEliminado(pagina, dpi)
					arreglar := A.recorrerArbol(A.Raiz)
					if arreglar != nil {
						A.usuarioEliminado(arreglar, 0)
					}
					if arreglar == nil && A.Raiz.Llaves[0] == nil {
						A.Raiz = nil
					}
					existe = true
					return true
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := A.ExisteBEliminar(pagina.Llaves[i].Izq, dpi, contrasenia)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := A.ExisteBEliminar(pagina.Llaves[i].Der, dpi, contrasenia)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := A.ExisteBEliminar(pagina.Llaves[i].Der, dpi, contrasenia)
						if a == true {
							return true
						}
					}
				}
			}
		}else{
			return false
		}
	}
	return false
}

func (A *ArbolB) usuarioEliminado(pagina *Pagina, dpi int) *Pagina{
	//Para hojas
	eliminado := false
	hoja := false
	for i := 0; i < len(pagina.Llaves); i++ {
		if pagina.Llaves[i] != nil {
			if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
				if pagina.Llaves[i].Izq == nil && pagina.Llaves[i].Der == nil {
					hoja = true
					pagina.Llaves[i] = nil
					eliminado = true
				}
			}
			if eliminado == true {
				pagina.Llaves[i]=pagina.Llaves[i+1]
			}
		}
	}
	if dpi == 0 {
		hoja = true
	}
	cambiado := false
	if pagina.Llaves[1] == nil && hoja == true{
		padre := pagina.NodoPadre
		if padre != nil {
			for i := 0; i < len(padre.Llaves); i++ {
				if padre.Llaves[i] != nil {
					if padre.Llaves[i].Izq == pagina {
						if padre.Llaves[i].Der.Llaves[2] != nil {
							rebalanceoDer(padre.Llaves[i].Der, pagina.NodoPadre, pagina)
							cambiado = true
							break
						}
						if padre.Llaves[i].Izq.Llaves[2] != nil {
							rebalanceoDer(padre.Llaves[i].Izq, pagina.NodoPadre, pagina)
							cambiado = true
							break
						}
					}
					if padre.Llaves[i].Der == pagina {
						if padre.Llaves[i].Izq.Llaves[2] != nil {
							rebalanceoIzq(padre.Llaves[i].Izq, pagina.NodoPadre, pagina)
							cambiado = true
							break
						}
						if padre.Llaves[i].Der.Llaves[2] != nil {
							rebalanceoDer(padre.Llaves[i].Der, pagina.NodoPadre, pagina)
							cambiado = true
							break
						}
					}
				}
			}
			if cambiado == false {
				for i := 0; i < len(padre.Llaves); i++ {
					if padre.Llaves[i] != nil {
						if padre.Llaves[i].Izq == pagina {
							A.juntarTodosDer(padre.Llaves[i].Der, pagina.NodoPadre, pagina)
							break
						}
						if padre.Llaves[i].Der == pagina {
							A.juntarTodosIzq(padre.Llaves[i].Izq, pagina.NodoPadre, pagina)
							break
						}
					}
				}
			}
		}
	}


	// No Hoja
	eliminado = false
	for i := 0; i < len(pagina.Llaves); i++ {
		if pagina.Llaves[i] != nil {
			if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
				if pagina.Llaves[i].Izq != nil || pagina.Llaves[i].Der != nil {
					var menorMayor Llave
					for j := 0; j < len(pagina.Llaves[i].Izq.Llaves); j++ {
						if pagina.Llaves[i].Izq.Llaves[j+1] == nil {
							menorMayor = *pagina.Llaves[i].Izq.Llaves[j]
							if menorMayor.Izq != nil || menorMayor.Der != nil{
								menorMayor = *obtenerMenorMayor(menorMayor.Der)
							}
							break
						}
					}
					for k := 0; k < len(pagina.Llaves[i].Izq.Llaves); k++ {
						if pagina.Llaves[i].Izq.Llaves[k] != nil {
							if *pagina.Llaves[i].Izq.Llaves[k] == menorMayor {
								pagina.Llaves[i].Izq.Llaves[k] = nil
								eliminado = true
							}
							if eliminado == true {
								pagina.Llaves[i].Izq.Llaves[k]= pagina.Llaves[i].Izq.Llaves[k+1]
							}
						}
					}
					menorMayor.Izq = pagina.Llaves[i].Izq
					menorMayor.Der = pagina.Llaves[i].Der
					pagina.Llaves[i] = &menorMayor
					eliminado = true
					if pagina.Llaves[i].Izq.Llaves[1] == nil {
						A.usuarioEliminado(pagina.Llaves[i].Izq, 0)
					}
				}
			}
		}
	}
	return pagina
}

func rebalanceoIzq(presta *Pagina, padre *Pagina, pide *Pagina){
	var herIzq Llave
	eliminado := false
	for i := 0; i < len(presta.Llaves); i++ {
		if presta.Llaves[i+1] == nil {
			herIzq = *presta.Llaves[i]
			break
		}
	}
	for i := 0; i < len(presta.Llaves); i++ {
		if presta.Llaves[i] != nil {
			if *presta.Llaves[i] == herIzq {
				//if presta.Llaves[i].Izq == nil && presta.Llaves[i].Der == nil {
				presta.Llaves[i] = nil
				eliminado = true
				//}
			}
			if eliminado == true {
				presta.Llaves[i]= presta.Llaves[i+1]
			}
		}
	}
	var nodoSeparador *Llave
	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil {
			if padre.Llaves[i].Izq == presta && padre.Llaves[i].Der == pide {

				nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: nil, Der: nil}
				if herIzq.Der != nil {
					nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: herIzq.Der, Der: nil}
					nodoSeparador.Izq.NodoPadre = pide
				}
				herIzq.Izq = padre.Llaves[i].Izq
				herIzq.Der = padre.Llaves[i].Der
				padre.Llaves[i] = &herIzq
			}
		}
	}

	for i := 0; i < len(pide.Llaves); i++ {
		if i == 0{
			pide.Colocar(1, pide.Llaves[0])
			pide.Colocar(0, nodoSeparador)
			pide.Llaves[0].Der = pide.Llaves[1].Izq
		}
	}
}

func rebalanceoDer(presta *Pagina, padre *Pagina, pide *Pagina){
	var herIzq Llave
	eliminado := false
	herIzq = *presta.Llaves[0]
	hoja := false
	noHoja := false
	var izqNoHoja *Pagina
	var derNoHoja *Pagina

	for i := 0; i < len(presta.Llaves); i++ {
		if presta.Llaves[i] != nil {
			if *presta.Llaves[i] == herIzq {
				if presta.Llaves[i].Izq == nil && presta.Llaves[i].Der == nil {
					presta.Llaves[i] = nil
					hoja = true
					eliminado = true
				}else{
					noHoja = true
					izqNoHoja = presta.Llaves[i].Izq
					derNoHoja = presta.Llaves[i].Der
					presta.Llaves[i] = nil
					hoja = true
					eliminado = true
				}
			}
			if eliminado == true {
				presta.Llaves[i]= presta.Llaves[i+1]
			}
		}
	}
	var nodoSeparador *Llave
	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil && hoja == true{
			if padre.Llaves[i].Der == presta && padre.Llaves[i].Izq == pide {
				nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: nil, Der: nil}
				if noHoja == true {
					if izqNoHoja != nil {
						if padre.Llaves[i].Usuario.DPI < izqNoHoja.Llaves[0].Usuario.DPI {
							nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: nil , Der: izqNoHoja}
							nodoSeparador.Der.NodoPadre = pide
						}
					}
					if derNoHoja != nil{
						if padre.Llaves[i].Usuario.DPI > derNoHoja.Llaves[0].Usuario.DPI {
							nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: derNoHoja, Der: nil }
							nodoSeparador.Izq.NodoPadre = pide
						}
					}
				}
				herIzq.Izq = padre.Llaves[i].Izq
				herIzq.Der = padre.Llaves[i].Der
				padre.Llaves[i] = &herIzq
			}
		}
	}

	for i := 0; i < len(pide.Llaves); i++ {
		if i == 0 && hoja == true{
			pide.Colocar(0, pide.Llaves[0])
			pide.Colocar(1, nodoSeparador)
			pide.Llaves[1].Izq = pide.Llaves[0].Der
		}
	}
}

func (A *ArbolB) juntarTodosDer(migra *Pagina, padre *Pagina, elimina *Pagina){
	ag:=false
	var nodoSeparador *Llave
	ordenar := false
	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil {
			if padre.Llaves[i].Der == migra && padre.Llaves[i].Izq == elimina {
				nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: nil, Der: nil}
				padre.Llaves[i] = nil
				ordenar = true
			}
			if ordenar == true {
				padre.Llaves[i] = padre.Llaves[i+1]
			}

		}
	}

	for i := len(migra.Llaves)-1; i >= 0; i-- {
		for migra.Llaves[0] != nil{
			if i == 0 {
				migra.Llaves[0] = nil
				break
			}
			migra.Llaves[i] = migra.Llaves[i-1]
			i--
		}
		if migra.Llaves[i] == nil && ag == false{
			migra.Llaves[i] = nodoSeparador
			ag = true
			i = len(migra.Llaves)-1
		}
		if ag == true {
			for migra.Llaves[0] != nil{
				if i == 0 {
					migra.Llaves[0] = nil
					break
				}
				migra.Llaves[i] = migra.Llaves[i-1]
				i--
			}
			nodoNuevo := &Llave{Usuario: elimina.Llaves[0].Usuario, Izq: nil, Der: nil}
			if elimina.Llaves[0].Izq != nil {
				nodoNuevo = &Llave{Usuario: elimina.Llaves[i].Usuario, Izq: elimina.Llaves[i].Izq , Der: elimina.Llaves[i].Der}
				nodoNuevo.Izq.NodoPadre = migra
				nodoNuevo.Der.NodoPadre = migra
			}
			migra.Llaves[i] = nodoNuevo
			break
		}
	}

	migra.NodoPadre = padre

	for i := 0; i < len(migra.Llaves); i++ {
		if migra.Llaves[i] != nil {
			if migra.Llaves[i].Izq == nil && migra.Llaves[i].Der == nil && i > 0 && i <= 3 && migra.Llaves[i+1] != nil {
				migra.Llaves[i].Izq = migra.Llaves[i-1].Der
				migra.Llaves[i].Der = migra.Llaves[i+1].Izq
			}
			if migra.Llaves[i+1] == nil && migra.Llaves[i].Der != nil{
				migra.Llaves[i].Der.NodoPadre = migra
				migra.Llaves[i].Izq.NodoPadre = migra
			}
		}
	}

	papa := false
	for i := 0; i < len(migra.NodoPadre.Llaves); i++ {
		if migra.NodoPadre.Llaves[i] != nil {
			papa = true
		}
	}

	if papa == false {
		A.Raiz = migra
	}

	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil {
			if padre.Llaves[i].Usuario.DPI < migra.Llaves[i].Usuario.DPI {
				if padre.Llaves[i+1] != nil {
					if padre.Llaves[i].Usuario.DPI < migra.Llaves[i].Usuario.DPI && padre.Llaves[i+1].Usuario.DPI > migra.Llaves[i].Usuario.DPI{
						padre.Llaves[i].Der = nil
						padre.Llaves[i+1].Izq = nil
						padre.Llaves[i].Der = migra
					}
				}else{
					padre.Llaves[i].Der = migra
				}
			}
			if padre.Llaves[i].Usuario.DPI > migra.Llaves[i].Usuario.DPI && i == 0 {
				padre.Llaves[i].Izq = migra
			}
		}
	}

	if padre.Llaves[1] == nil && padre.NodoPadre != nil{
		cambiado := false
		for i := 0; i < len(padre.NodoPadre.Llaves); i++ {
			if padre.NodoPadre.Llaves[i] != nil {

				if padre.NodoPadre.Llaves[i].Izq == padre {
					if padre.NodoPadre.Llaves[i].Der.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Der, padre.NodoPadre, padre)
						cambiado = true
						break
					}
					if padre.NodoPadre.Llaves[i].Izq.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Izq, padre.NodoPadre, padre)
						cambiado = true
						break
					}
				}
				if padre.NodoPadre.Llaves[i].Der == padre {
					if padre.NodoPadre.Llaves[i].Izq.Llaves[2] != nil {
						rebalanceoIzq(padre.NodoPadre.Llaves[i].Izq, padre.NodoPadre, padre)
						cambiado = true
						break
					}
					if padre.NodoPadre.Llaves[i].Der.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Der, padre.NodoPadre, padre)
						cambiado = true
						break
					}
				}
			}
		}
		if cambiado == false {
			for k := 0; k < len(padre.NodoPadre.Llaves); k++ {
				if padre.NodoPadre.Llaves[k] != nil {
					if padre.NodoPadre.Llaves[k].Izq == padre {
						A.juntarTodosDer(padre.NodoPadre.Llaves[k].Der, padre.NodoPadre, padre)
						break
					}
					if padre.NodoPadre.Llaves[k].Der == padre {
						A.juntarTodosIzq(padre.NodoPadre.Llaves[k].Izq, padre.NodoPadre, padre)
						break
					}
				}
			}
		}
	}
}

func (A *ArbolB)juntarTodosIzq(migra *Pagina, padre *Pagina, elimina *Pagina){
	var nodoSeparador *Llave
	ordenar := false
	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil {
			if padre.Llaves[i].Izq == migra && padre.Llaves[i].Der == elimina {
				nodoSeparador = &Llave{Usuario: padre.Llaves[i].Usuario, Izq: nil, Der: nil}
				padre.Llaves[i] = nil
				ordenar = true
			}
			if ordenar == true {
				padre.Llaves[i] = padre.Llaves[i+1]
			}

		}
	}
	ag:=false
	for i := 0; i < len(migra.Llaves); i++ {
		if migra.Llaves[i] == nil && ag == false{
			migra.Llaves[i] = nodoSeparador
			ag = true
		}else if ag == true {
			migra.Llaves[i] = elimina.Llaves[0]
			break
		}
	}
	migra.NodoPadre = padre

	for i := 0; i < len(migra.Llaves); i++ {
		if migra.Llaves[i] != nil {
			if migra.Llaves[i].Izq == nil && migra.Llaves[i].Der == nil && i > 0 && i <= 3 && migra.Llaves[i+1] != nil {
				migra.Llaves[i].Izq = migra.Llaves[i-1].Der
				migra.Llaves[i].Der = migra.Llaves[i+1].Izq
			}
			if migra.Llaves[i+1] == nil && migra.Llaves[i].Der != nil{
				migra.Llaves[i].Der.NodoPadre = migra
				migra.Llaves[i].Izq.NodoPadre = migra
			}
		}
	}

	papa := false
	for i := 0; i < len(migra.NodoPadre.Llaves); i++ {
		if migra.NodoPadre.Llaves[i] != nil {
			papa = true
		}
	}

	if papa == false {
		A.Raiz = migra
	}

	for i := 0; i < len(padre.Llaves); i++ {
		if padre.Llaves[i] != nil {
			if padre.Llaves[i].Usuario.DPI < migra.Llaves[i].Usuario.DPI {
				if padre.Llaves[i+1] != nil {
					if padre.Llaves[i].Usuario.DPI < migra.Llaves[i].Usuario.DPI && padre.Llaves[i+1].Usuario.DPI > migra.Llaves[i].Usuario.DPI{
						padre.Llaves[i].Der = nil
						padre.Llaves[i+1].Izq = nil
						padre.Llaves[i].Der = migra
					}
				}else{
					padre.Llaves[i].Der = migra
				}
			}
			if padre.Llaves[i].Usuario.DPI > migra.Llaves[i].Usuario.DPI && i == 0 {
				padre.Llaves[i].Izq = migra
			}
		}
	}
	if padre.Llaves[1] == nil && padre.NodoPadre != nil{
		cambiado := false
		for i := 0; i < len(padre.NodoPadre.Llaves); i++ {
			if padre.NodoPadre.Llaves[i] != nil {
				if padre.NodoPadre.Llaves[i].Izq == padre {
					if padre.NodoPadre.Llaves[i].Der.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Der, padre.NodoPadre, padre)
						cambiado = true
						break
					}
					if padre.NodoPadre.Llaves[i].Izq.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Izq, padre.NodoPadre, padre)
						cambiado = true
						break
					}
				}
				if padre.NodoPadre.Llaves[i].Der == padre {
					if padre.NodoPadre.Llaves[i].Izq.Llaves[2] != nil {
						rebalanceoIzq(padre.NodoPadre.Llaves[i].Izq, padre.NodoPadre, padre)
						cambiado = true
						break
					}
					if padre.NodoPadre.Llaves[i].Der.Llaves[2] != nil {
						rebalanceoDer(padre.NodoPadre.Llaves[i].Der, padre.NodoPadre, padre)
						cambiado = true
						break
					}
				}
			}
		}

		if cambiado == false {
			for k := 0; k < len(padre.NodoPadre.Llaves); k++ {
				if padre.NodoPadre.Llaves[k] != nil {
					if padre.NodoPadre.Llaves[k].Izq == padre {
						A.juntarTodosDer(padre.NodoPadre.Llaves[k].Der, padre.NodoPadre, padre)
						break
					}
					if padre.NodoPadre.Llaves[k].Der == padre {
						A.juntarTodosIzq(padre.NodoPadre.Llaves[k].Izq, padre.NodoPadre, padre)
						break
					}
				}
			}
		}
	}
}

func obtenerMenorMayor (pagina *Pagina) *Llave {
	var regresa *Llave
	for i := 0; i < len(pagina.Llaves); i++ {
		if pagina.Llaves[i] != nil {
			if pagina.Llaves[i+1] == nil {
				regresa = pagina.Llaves[i]
			}
		}
	}
	eliminado := false
	for i := 0; i < len(pagina.Llaves); i++ {
		if pagina.Llaves[i] != nil {
			if pagina.Llaves[i] == regresa {
				pagina.Llaves[i] = nil
				eliminado = true
			}
			if eliminado == true {
				pagina.Llaves[i] = pagina.Llaves[i+1]
			}
		}
	}
	if regresa.Der != nil{
		regresa = obtenerMenorMayor(regresa.Der)
	}
	return regresa
}

func (A *ArbolB) recorrerArbol(pagina *Pagina) *Pagina{
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil && pagina.Llaves[i+1] != nil{
				pagina.Llaves[i+1].Izq = pagina.Llaves[i].Der
			}
		}
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.NodoPadre != nil {
					if pagina.Llaves[1] == nil {
						return pagina
					}
				}
			}
		}
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				regresar := A.recorrerArbol(pagina.Llaves[i].Izq)
				if regresar != nil {
					return regresar
				}
				regresar = A.recorrerArbol(pagina.Llaves[i].Der)
				if regresar != nil {
					return regresar
				}
			}
		}
	}
	return nil
}
