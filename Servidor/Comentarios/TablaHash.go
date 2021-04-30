package Comentarios

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type NodoCom struct {
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Producto int `json:"Producto"`
	Comentarios []Respuestas `json:"Comentarios"`
}

type Respuestas struct{
	Dpi int `json:"DPI"`
	Fecha string `json:"Fecha"`
	Comentario string `json:"Comentario"`
}

type NodoComentarios struct {
	DpiPadre int
	Comentario string
	FechaComentario string
	Respuestas *TablaHash
}

type TablaHash struct{
	Tam int
	Carga int
	PorcentajeIn int
	PorcentajeCre int
	Arreglo []*NodoComentarios
}

func NuevaTabla(tam int, porcentaje int, porcentajeCre int) *TablaHash{
	arreglo := make([]*NodoComentarios, tam)
	return &TablaHash{tam, 0, porcentaje, porcentajeCre, arreglo}
}

func (T *TablaHash)Insertar(nuevo int, valor string, fecha string){
	tabla := NuevaTabla(7, 50, 20)
	nuevoNodo := &NodoComentarios{nuevo, valor, fecha, tabla}
	posicion := T.Posicion(nuevo)
	T.Arreglo[posicion] = nuevoNodo
	T.Carga++
	if ((T.Carga*100)/T.Tam) > T.PorcentajeIn {
		nuevoTam := T.Tam
		primo := false
		contador := nuevoTam
		for primo == false {
			for j := 1; j < contador; j++ {
				cuenta := 0
				for k := 1; k <= j; k++ {
					if j % k == 0{
						cuenta++
					}
				}
				if cuenta == 2 && j > nuevoTam{
					nuevoTam = j
					primo = true
					break
				}
			}
			contador ++
		}
		nuevoArreglo := make([]*NodoComentarios, nuevoTam)
		antiguo := T.Arreglo
		T.Arreglo = nuevoArreglo
		T.Tam = nuevoTam
		aux := 0
		for i := 0; i < len(antiguo); i++ {
			if antiguo[i] != nil {
				aux =T.Posicion(antiguo[i].DpiPadre)
				nuevoArreglo[aux] = antiguo[i]
			}
		}
	}
}

func (T *TablaHash) Posicion(Clave int) int{
	i,p := 0,0
	p1 := float64(Clave) * 0.2520
	parteEntera := int(p1)
	d := (float64(Clave) * 0.2520) - float64(parteEntera)
	p = int(float64(T.Tam) * d)

	for T.Arreglo[p] != nil {
		i++
		p = int(float64(T.Tam) * d)
		p = p+(i*i)
		for p >= T.Tam {
			p = p-T.Tam
		}
		if i == 5000 {
			break
		}
	}
	return p
}

func (T *TablaHash) Buscar (Clave int, valor string, fecha string) int{
	i,p := 0,0
	p1 := float64(Clave) * 0.2520
	parteEntera := int(p1)
	d := (float64(Clave) * 0.2520) - float64(parteEntera)
	p = int(float64(T.Tam) * d)
	if T.Arreglo[p].DpiPadre == Clave && T.Arreglo[p].Comentario == valor && T.Arreglo[p].FechaComentario == fecha{
		return p
	}
	for T.Arreglo[p] != nil {
		if T.Arreglo[p].DpiPadre == Clave && T.Arreglo[p].Comentario == valor && T.Arreglo[p].FechaComentario == fecha{
			return p
		}
		i++
		p = int(float64(T.Tam) * d)
		p = p+(i*i)
		for p >= T.Tam {
			p = p-T.Tam
		}
		if i == 5000 {
			break
		}
	}
	return p
}

func (T *TablaHash) Imprimir(){
	data:= make([][]string, T.Tam)
	for i := 0; i < len(T.Arreglo); i++ {
		tmp := make([]string, 4)
		aux := T.Arreglo[i]
		if aux != nil {
			tmp[0] = strconv.Itoa(i)
			tmp[1] = strconv.Itoa(aux.DpiPadre)
			tmp[2] = aux.Comentario
			tmp[3] = aux.FechaComentario
		}else{
			tmp[0] = strconv.Itoa(i)
			tmp[1] = "---"
			tmp[2] = "---"
			tmp[3] = "---"
		}
		data[i] = tmp
	}
	tabla := tablewriter.NewWriter(os.Stdout)
	tabla.SetHeader([]string{"Posicion","DPI COMENTARIO", "COMENTARIO", "FECHA"})
	tabla.SetFooter([]string{"Tamaño:  "+ strconv.Itoa(T.Tam), "Carga: "+ strconv.Itoa(T.Carga), "", ""})
	tabla.AppendBulk(data)
	tabla.Render()
}
