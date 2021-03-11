package main

import (
	"./Inventario"
	"./Listas"
	"./MatrizDispersa"
	"./Reportes"
	"./TiendaEspecifica"
	"./Compras"
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
var eliminar TiendaEspecifica.Buscar
var arbol Inventario.General
var nodoArbol Inventario.NodoArbol
var matriz MatrizDispersa.General
var metodosMatriz MatrizDispersa.Matriz
var listaAnioa MatrizDispersa.ListaAnio
var carritojson Compras.General

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/cargarArchivos", cargar).Methods("Post") //ORDENAR TIENDAS POR ASCII
	router.HandleFunc("/getArreglo", arreglo).Methods("Get")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("Post")
	router.HandleFunc("/id/{numero}", busquedaposicion).Methods("Get")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("Delete")
	router.HandleFunc("/guardar", guardarTodo).Methods("Get")
	router.HandleFunc("/carritoCompras", carrito).Methods("Post")
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
				tiendas.Logo = c.Logo
				tienda := Listas.Tiendas{NombreTienda: c.Nombre, Descripcion: c.Descripcion, Contacto: c.Contacto, Calificacion: c.Calificacion, Logo: c.Logo}
				depa := Listas.Departamentos{NombreDepartamento: departamentos.NombreDepartamento, Tienda: tienda}
				nuevo := Listas.Nodo{Indice: a.Indice, Departamento: depa }
				fmt.Fprint(w, list.Insertar(&nuevo))
			}
		}
	}

	Vector = list.CrearMatriz()
	Indi := list.Indi()
	Departa := list.Departa()
	if len(Vector) != 0 {
		json.Unmarshal(reqBody, &arbol)
		for i := 0; i < len(arbol.Inventarios); i++ {
			NombreTienda := arbol.Inventarios[i].NombreTienda
			Departamento := arbol.Inventarios[i].Departamento
			Calificacion := arbol.Inventarios[i].Calificacion
			Tercero := posicionTercero(NombreTienda, Departamento,Calificacion, Indi, Departa)
			imp := Vector[Tercero].ListGA.Cabeza
			for imp != nil {
				if imp.NombreTienda == NombreTienda{
					arbolPosicion := imp.Inventario.NuevoArbol()
					Productos := arbol.Inventarios[i].Productos
					for j := 0; j < len(Productos); j++ {
						a := Productos[j]
						nodoArbol.NombreProducto = a.NombreProducto
						nodoArbol.Codigo = a.Codigo
						nodoArbol.Descripcion = a.Descripcion
						nodoArbol.Precio = a.PrecioP
						nodoArbol.Cantidad = a.Cantidad
						nodoArbol.Imagen = a.Imagen
						arbolPosicion.Insertar(nodoArbol.NombreProducto, nodoArbol.Codigo, nodoArbol.Descripcion, nodoArbol.Precio, nodoArbol.Cantidad, nodoArbol.Imagen)
					}
					imp.Inventario = *arbolPosicion
				}
				imp = imp.Siguiente
			}
		}

		json.Unmarshal(reqBody, &matriz)
		for i := 0; i < len(matriz.Pedidos); i++ {
			fecha := matriz.Pedidos[i].Fecha
			mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
			anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
			existeAnio := EncontrarAnio(&listaAnioa, anio)
			if existeAnio == false {
				var listaMes MatrizDispersa.ListaMes
				nodoAnio := MatrizDispersa.NodoAnio{Anio: anio, ListaMatricesMes: &listaMes}
				listaAnioa.Insertar(&nodoAnio)

			}
			existeMes := EncontrarMes(&listaAnioa, anio, mes)
			if existeMes == false {
				imp := listaAnioa.Cabeza
				for imp != nil {
					if imp.Anio == anio {
						nodoMes := MatrizDispersa.NodoMes{Mes: mes, MatrizMes: &MatrizDispersa.Matriz{Mes: mes, Anio: anio}}
						imp.ListaMatricesMes.Insertar(&nodoMes)
					}
					imp = imp.Siguiente
				}
			}
			NombreTienda := matriz.Pedidos[i].NombreTienda
			Departamento := matriz.Pedidos[i].Departamento
			Calificacion := matriz.Pedidos[i].Calificacion
			Productos := matriz.Pedidos[i].Productos
			Tercero := posicionTercero(NombreTienda, Departamento,Calificacion, Indi, Departa)
			for j := 0; j < len(Productos); j++ {
				imp := Vector[Tercero].ListGA.Cabeza
				for imp != nil {
					if imp.NombreTienda == NombreTienda && imp.Calificacion == Calificacion {
						a := inOrden(imp.Inventario.Raiz, Productos[j].Codigo)
						if a == true {
							impa := listaAnioa.Cabeza
							for impa != nil {
								if impa.Anio == anio {
									impr := impa.ListaMatricesMes.Cabeza
									for impr != nil {
										if impr.Mes == mes {
											nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, NombreTienda, Departamento,Calificacion,Productos[j].Codigo)
											impr.MatrizMes.Insertar(nodoPedido)
										}
										impr = impr.Siguiente
									}
								}
								impa = impa.Siguiente
							}
						}
					}
					imp = imp.Siguiente
				}
			}
		}
		if listaAnioa.Cabeza != nil {
			listaAnioa = *metodosMatriz.BurbujaAnio(listaAnioa)
			impa := listaAnioa.Cabeza
			for impa != nil{
				impa.ListaMatricesMes = metodosMatriz.BurbujaMes(*impa.ListaMatricesMes)
				/*
				impm := impa.ListaMatricesMes.Cabeza
				for impm != nil{
					impm.MatrizMes.Imprimir()
					impm.MatrizMes.Imprimir2()
					fmt.Println()
					fmt.Println("---------------------------------------------------------------------------------------------------------")
					fmt.Println("---------------------------------------------------------------------------------------------------------")
					fmt.Println()

					impm = impm.Siguiente
				}*/
				impa = impa.Siguiente
			}
		}
	}
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
			retorno := TiendaEs{mens[0], mens[1], mens[2], cali, mens[4]}
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
				Logo := tiendaE[4]
				cal,_ := strconv.Atoi(tiendaE[3])
				buscarT := busquedaTienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: cal, Logo: Logo}
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
		eliminar.Depa = tiendaEl.Categoria
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
								tiendasRE = append(tiendasRE, TiendaR{Nombre: imp.NombreTienda, Descripcion: imp.Descripcion, Contacto: imp.Contacto, Calificacion: imp.Calificacion, Logo: imp.Logo})
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

func carrito (w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	Indi := list.Indi()
	Departa := list.Departa()
	json.Unmarshal(reqBody, &carritojson)
	for i := 0; i < len(carritojson.Pedidos); i++ {
		a := carritojson.Pedidos[i]
		fecha := a.Fecha
		mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
		anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
		Tienda := a.NombreTienda
		Departamento := a.Departamento
		Calificacion := a.Calificacion
		Estado := a.Estado
		Productos := a.CodigoProductos
		Tercero := posicionTercero(Tienda, Departamento, Calificacion, Indi, Departa)
		for j := 0; j < len(Productos); j++ {
			imp := Vector[Tercero].ListGA.Cabeza
			CodigoProducto := Productos[j].Codigo
			Cantidad := Productos[j].Cantidad
			for imp != nil{
				if imp.NombreTienda == Tienda && imp.Calificacion == Calificacion {
					if Estado == "Vendido" {

						existeAnio := EncontrarAnio(&listaAnioa, anio)
						if existeAnio == false {
							var listaMes MatrizDispersa.ListaMes
							nodoAnio := MatrizDispersa.NodoAnio{Anio: anio, ListaMatricesMes: &listaMes}
							listaAnioa.Insertar(&nodoAnio)

						}
						existeMes := EncontrarMes(&listaAnioa, anio, mes)
						if existeMes == false {
							imp := listaAnioa.Cabeza
							for imp != nil {
								if imp.Anio == anio {
									nodoMes := MatrizDispersa.NodoMes{Mes: mes, MatrizMes: &MatrizDispersa.Matriz{Mes: mes, Anio: anio}}
									imp.ListaMatricesMes.Insertar(&nodoMes)
								}
								imp = imp.Siguiente
							}
						}

						arbolTienda := imp.Inventario.Raiz
						Compras.DescontarProducto(arbolTienda, CodigoProducto, Cantidad)
						existeProducto := inOrden(imp.Inventario.Raiz, Productos[j].Codigo)
						if existeProducto == true {
							impa := listaAnioa.Cabeza
							for impa != nil {
								if impa.Anio == anio {
									impr := impa.ListaMatricesMes.Cabeza
									for impr != nil {
										if impr.Mes == mes {
											nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, Tienda, Departamento,Calificacion,Productos[j].Codigo)
											impr.MatrizMes.Insertar(nodoPedido)
										}
										impr = impr.Siguiente
									}
								}
								impa = impa.Siguiente
							}
						}

					}
				}
				imp = imp.Siguiente
			}
		}
	}
	if listaAnioa.Cabeza != nil {
		listaAnioa = *metodosMatriz.BurbujaAnio(listaAnioa)
		impa := listaAnioa.Cabeza
		for impa != nil{
			impa.ListaMatricesMes = metodosMatriz.BurbujaMes(*impa.ListaMatricesMes)
			impm := impa.ListaMatricesMes.Cabeza
			for impm != nil{
				/*
				impm.MatrizMes.Imprimir()
				impm.MatrizMes.Imprimir2()
				fmt.Println()
				fmt.Println("---------------------------------------------------------------------------------------------------------")
				fmt.Println("---------------------------------------------------------------------------------------------------------")
				fmt.Println()

				 */
				impm = impm.Siguiente
			}
			impa = impa.Siguiente
		}
	}
}

//----------------------------------------------------------------------------------------------------------------------

func posicionTercero(Nombre string, Depa string, Calificacion int, Indices []string, Departamentos []string) int{
	indice := strings.Split(Nombre, "")
	posFila := Posicion(Indices, indice[0])
	posColumna := Posicion(Departamentos, Depa)
	Primero := posFila-0
	Segundo := Primero * len(Departamentos) + posColumna
	Tercero := Segundo*5+(Calificacion-1)
	return Tercero
}

func Posicion(arreglo []string, busqueda string) int {
	for indice, valor := range arreglo {
		if valor == busqueda {
			return indice
		}
	}
	return -1
}

func inOrden(raiz *Inventario.NodoArbol, codigo int) bool{
	if raiz!=nil {
		if raiz.Codigo == codigo {
			return true
		}
		a := inOrden(raiz.Izq, codigo)
		if a == true {
			return true
		}
		b := inOrden(raiz.Der, codigo)
		if b == true {
			return true
		}
	}
	return false
}

func EncontrarAnio(lista *MatrizDispersa.ListaAnio, anio int) bool{
	impa := lista.Cabeza
	for impa != nil{
		if impa.Anio == anio {
			return true
		}
		impa = impa.Siguiente
	}
	return false
}

func EncontrarMes(lista *MatrizDispersa.ListaAnio, anio int, mes int) bool{
	impa := lista.Cabeza
	for impa != nil {
		if impa.Anio == anio {
			impr := impa.ListaMatricesMes.Cabeza
			for impr != nil {
				if impr.Mes == mes {
					return true
				}
				impr = impr.Siguiente
			}
		}
		impa = impa.Siguiente
	}
	return false
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
	Logo string `json:"Logo"`
}

type Mensaje struct {
	Retorna string `json:"Alerta"`
}

type TiendaEs struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
	Logo string `json:"Logo"`
}

type general struct{
	Tiendas []busquedaTienda `json:"Tiendas"`
}

type busquedaTienda struct{
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
	Logo string `json:"Logo"`
}
