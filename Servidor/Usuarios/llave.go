package Usuarios

import "strconv"

type Llave struct{
	Usuario *NodoUsuario
	Izq *Pagina
	Der *Pagina
}

type NodoUsuario struct{
	DPI string
	Nombre string
	Correo string
	Contra string
	Cuenta string
}

type General struct {
	Usuarios []Usuario `json:"Usuarios"`
}

type Usuario struct {
	DPI int `json:"Dpi"`
	Nombre string `json:"Nombre"`
	Correo string `json:"Correo"`
	Contra string `json:"Password"`
	Cuenta string `json:"Cuenta"`
}

type Inicio struct {
	Nombre int `json:"Nombre"`
	Contra string `json:"Password"`
}


func NuevaLlave(dpi int, nombre string, correo string, contrasenia string, cuenta string) *Llave{
	dpiUsuario := strconv.Itoa(dpi)
	usuario := NodoUsuario{DPI: dpiUsuario, Nombre: nombre, Correo: correo, Contra: contrasenia, Cuenta: cuenta}
	k := Llave{&usuario, nil, nil}
	return &k
}