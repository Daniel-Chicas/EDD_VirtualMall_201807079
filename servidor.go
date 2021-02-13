package main

import (
	"./Listas"
	"./Reportes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)
var ms Listas.General
var nodo Listas.Nodo
var departamentos Listas.Departamentos
var tiendas Listas.Tienda
var list Listas.Lista
var reportes Reportes.Lista
var Vector []Listas.NodoArray

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/cargartienda", cargar).Methods("Post")
	router.HandleFunc("/getArreglo", arreglo).Methods("Get")
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
}

func arreglo(w http.ResponseWriter, r *http.Request){
	archivo := reportes.Arreglo(Vector)
	fmt.Println(archivo)
}

type Mensaje struct {
	Retorna string `json:"Regresa"`
}