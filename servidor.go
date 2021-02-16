package main

import (
	"./Listas"
	"./Reportes"
	"./TiendaEspecifica"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var ms Listas.General
var nodo Listas.Nodo
var departamentos Listas.Departamentos
var tiendas Listas.Tienda
var list Listas.Lista
var reportes Reportes.Lista
var Vector []Listas.NodoArray
var tiendaEsp TiendaEspecifica.General
var buscar TiendaEspecifica.Buscar


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/cargartienda", cargar).Methods("Post")
	router.HandleFunc("/getArreglo", arreglo).Methods("Get")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("Post")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func inicio(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "INICIO PROYECTO ESTRUCTURAS DE DATOS")
}

func cargar(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.Inicio); i++ {
		a := ms.Inicio[i]
		nodo.Indice = a.Indice
		for j := 0; j < len(a.Departamentos); j++ {
			b := a.Departamentos[j]
			departamentos.NombreDepartamento = a.Departamentos[j].Nombre
			for k := 0; k < len(b.Tiendas); k++ {
				c := b.Tiendas[k]
				tiendas.Nombre = c.Nombre
				tiendas.Descripcion = c.Descripcion
				tiendas.Contacto = c.Contacto
				tiendas.Calificacion = c.Calificacion
				tienda := Listas.Tiendas{NombreTienda: c.Nombre, Descripcion: c.Descripcion, Contacto: c.Contacto, Calificacion: c.Calificacion}
				depa := Listas.Departamentos{NombreDepartamento: departamentos.NombreDepartamento, Tienda: tienda}
				nuevo := Listas.Nodo{Indice: a.Indice, Departamento: depa }
				fmt.Fprint(w, list.Insertar(&nuevo))
			}
		}
	}
	Vector = list.CrearMatriz()
	w.Header().Set("Content-type", "application/json")
	if list.Cabeza == nil{
		mensaje := Mensaje{"NO SE HA PODIDO CARGAR EL ARCHIVO"}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}else{
		mensaje := Mensaje{"EL ARCHIVO HA SIDO GUARDADO EXISTOSAMENTE"}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}
} //ORDENAR TIENDAS POR ASCII

func arreglo(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		mensaje := Mensaje{reportes.Arreglo(Vector)}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func tiendaEspecifica (w http.ResponseWriter, r *http.Request){
	Indi := list.Indi()
	Departa := list.Departa()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	if Indi != nil && Departa != nil{
		json.Unmarshal(reqBody, &tiendaEsp)
		buscar.Depa = tiendaEsp.Departamento
		buscar.NombreB = tiendaEsp.Nombre
		buscar.Cal = tiendaEsp.Calificacion
		mensaje := buscar.Buscar(Vector, Indi, Departa)
		if mensaje != "" {
			mens := strings.Split(mensaje, "&")
			cali, _ := strconv.Atoi(mens[3])
			retorno := TiendaEs{mens[0], mens[1], mens[2], cali}
			w.WriteHeader(http.StatusCreated)
			json.Unmarshal(reqBody, &ms)
			json.NewEncoder(w).Encode(retorno)
		}else{
			mensaje := Mensaje{"Revise el archivo de entrada."}
			w.WriteHeader(http.StatusCreated)
			json.Unmarshal(reqBody, &ms)
			json.NewEncoder(w).Encode(mensaje)
		}
	}else{
		mensaje := Mensaje{"Debe cargar un archivo."}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}
}

type Mensaje struct {
	Retorna string `json:"Alerta"`
}

type TiendaEs struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
}