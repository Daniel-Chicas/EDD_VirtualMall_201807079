package Reportes

import (
	"../Listas"
	"fmt"
	"strconv"
)

type Lista struct {
	llegue string
}

func (L *Lista) Arreglo(vector []Listas.NodoArray) string{
	for i := 0; i < len(vector); i++ {
		if(i % 10 == 0){
			fmt.Println("Creo archivo\t"+vector[i].Indice+"-->"+vector[i].Departamento+"-->"+strconv.Itoa(vector[i].Calificacion))
		}else{
			fmt.Println("Añadir archivo\t"+vector[i].Indice+"-->"+vector[i].Departamento+"-->"+strconv.Itoa(vector[i].Calificacion))
		}
	}
	return "Homla"
}