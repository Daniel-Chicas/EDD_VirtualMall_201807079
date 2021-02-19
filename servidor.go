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
var tiendaEl TiendaEspecifica.GeneralEliminar
var buscar TiendaEspecifica.Buscar
var eliminar TiendaEspecifica.Eliminar


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/cargartienda", cargar).Methods("Post") //ORDENAR TIENDAS POR ASCII
	router.HandleFunc("/getArreglo", arreglo).Methods("Get")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("Post")
	router.HandleFunc("/id/{numero}", busquedaposicion).Methods("Get")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("Delete")
	router.HandleFunc("/guardar", guardarTodo).Methods("Get")
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
			w.WriteHeader(http.StatusFound)
			json.Unmarshal(reqBody, &ms)
			json.NewEncoder(w).Encode(retorno)
		}else{
			mensaje := Mensaje{"Revise el archivo de entrada."}
			w.WriteHeader(http.StatusFound)
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

func busquedaposicion (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var busca []busquedaTienda
	posicion,_ := strconv.Atoi(vars["numero"])
	retorna := buscar.BusquedaPosicion(Vector, posicion)
	if retorna != "" {
		if retorna == "No hay tienda." || retorna == "Posición inválida."{
			mensaje := Mensaje{retorna}
			w.WriteHeader(http.StatusLocked)
			json.NewEncoder(w).Encode(mensaje)
		}else{
			tiendasGen := strings.Split(retorna, "%")
			for i := 1; i < len(tiendasGen); i++ {
				tiendaE := strings.Split(tiendasGen[i], "&")
				nombre := tiendaE[0]
				descripcion := tiendaE[1]
				contacto := tiendaE[2]
				cal,_ := strconv.Atoi(tiendaE[3])
				buscarT := busquedaTienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: cal}
				busca = append(busca, buscarT)
			}
			mensaje := general{busca}
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(mensaje)
		}
	}else{
		mensaje := Mensaje{"Debe ingresar un listado de tiendas."}
		w.WriteHeader(http.StatusLocked)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func eliminarTienda (w http.ResponseWriter, r * http.Request) {
	Indi := list.Indi()
	Departa := list.Departa()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error al insertar")
	}
	if Indi != nil && Departa != nil {
		json.Unmarshal(reqBody, &tiendaEl)
		eliminar.NombreB = tiendaEl.Nombre
		eliminar.Categ = tiendaEl.Categoria
		eliminar.Cal = tiendaEl.Calificacion
		Vector = eliminar.Eliminar(Vector, Indi, Departa)
		mensaje := Mensaje{Retorna: "La tienda ha sido eliminada con éxito."}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func guardarTodo (w http.ResponseWriter, r *http.Request){
	var tiendasRE []TiendaR
	var DepartamentosRE []DepartamentoR
	var datosRE []DatosR
	var generalRE GeneralR
	indi := list.Indi()
	departa := list.Departa()
	vector := Vector
	//tama := len(indi) * len(departa)
	//for l := 0; l < tama; l++ {
		for j := 0; j < len(indi); j++ {
			for k := 0; k < len(departa); k++ {
				for i := 0; i < len(vector); i++ {
					if vector[i].Indice == indi[j] {
						if vector[i].Departamento == departa[k] {
							imp := vector[i].ListGA.Cabeza
							for imp != nil {
								tiendasRE = append(tiendasRE, TiendaR{Nombre: imp.NombreTienda, Descripcion: imp.Descripcion, Contacto: imp.Contacto, Calificacion: imp.Calificacion})
								imp = imp.Siguiente
							}
						}
					}
				}
				DepartamentosRE = append(DepartamentosRE, DepartamentoR{NombreDepa: departa[k], Tiendas: tiendasRE})
				tiendasRE = nil
			}
			datosRE = append(datosRE, DatosR{Indice: indi[j], Departamentos: DepartamentosRE})
			DepartamentosRE = nil
		}
	generalRE = GeneralR{Inicio: datosRE}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(generalRE)
}

type GeneralR struct{
	Inicio []DatosR `json:"Datos"`
}

type DatosR struct{
	Indice string `json:"Indice"`
	Departamentos []DepartamentoR `json:"Departamentos"`
}

type DepartamentoR struct{
	NombreDepa string `json:"Nombre"`
	Tiendas []TiendaR `json:"Tiendas"`
}

type TiendaR struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
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

type general struct{
	Tiendas []busquedaTienda `json:"Tiendas"`
}

type busquedaTienda struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
}
