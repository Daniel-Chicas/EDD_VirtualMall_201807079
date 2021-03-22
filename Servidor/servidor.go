package main

import (
	"./Compras"
	"./Inventario"
	"./Listas"
	"./MatrizDispersa"
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
var eliminar TiendaEspecifica.Buscar
var arbol Inventario.General
var nodoArbol Inventario.NodoArbol
var matriz MatrizDispersa.General
var metodosMatriz MatrizDispersa.Matriz
var listaAnioa MatrizDispersa.ListaAnio
var carritojson Compras.General
var arregloProductos []Inventario.NodoArbol


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", inicio).Methods("Get")
	router.HandleFunc("/cargarArchivos", cargar).Methods("Post")
	router.HandleFunc("/getArreglo", arreglo).Methods("Get")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("Post")
	router.HandleFunc("/id/{numero}", busquedaposicion).Methods("Get")
	router.HandleFunc("/Tienda/{infoTienda}", busquedaProductosTienda).Methods("Get")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("Delete")
	router.HandleFunc("/guardar", guardarTodo).Methods("Get")
	router.HandleFunc("/carritoCompras", carrito).Methods("Post")
	router.HandleFunc("/Pedido/{infoPedido}", pedidoMes).Methods("Get")
	router.HandleFunc("/DatosMatriz", datosMatriz).Methods("Get")
	router.HandleFunc("/ImagenMatriz/{datos}", imagenMatriz).Methods("Get")
	router.HandleFunc("/Arbol/{datos}", arbolTienda).Methods("Get")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func inicio(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "INICIO PROYECTO ESTRUCTURAS DE DATOS")
}

func cargar(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	for reqBody[0] != 123{
		reqBody = remove(reqBody, 0)
	}
	for reqBody[len(reqBody)-1] != 125{
		reqBody = remove(reqBody, len(reqBody)-1)
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
						nodoArbol.Precio =  a.PrecioP
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
			dia,_ := strconv.Atoi(strings.Split(fecha, "-")[0])
			mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
			anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
			NombreTienda := matriz.Pedidos[i].NombreTienda
			Departamento := matriz.Pedidos[i].Departamento
			Calificacion := matriz.Pedidos[i].Calificacion
			Productos := matriz.Pedidos[i].Productos
			Tercero := posicionTercero(NombreTienda, Departamento,Calificacion, Indi, Departa)
			for j := 0; j < len(Productos); j++ {
				if Tercero<0{
					break
				}
				imp := Vector[Tercero].ListGA.Cabeza
				for imp != nil {
					if imp.NombreTienda == NombreTienda && imp.Calificacion == Calificacion {
						a := inOrden(imp.Inventario.Raiz, Productos[j].Codigo)
						if a == true {

							existeAnio := EncontrarAnio(&listaAnioa, anio)
							if existeAnio == false {
								var listaMes MatrizDispersa.ListaMes
								nodoAnio := MatrizDispersa.NodoAnio{Anio: anio, ListaMatricesMes: &listaMes}
								listaAnioa.Insertar(&nodoAnio)

							}
							existeMes := EncontrarMes(&listaAnioa, anio, mes)
							if existeMes == false {
								impm := listaAnioa.Cabeza
								for impm != nil {
									if impm.Anio == anio {
										nodoMes := MatrizDispersa.NodoMes{Mes: mes, MatrizMes: &MatrizDispersa.Matriz{Mes: mes, Anio: anio}}
										impm.ListaMatricesMes.Insertar(&nodoMes)
									}
									impm = impm.Siguiente
								}
							}

							impa := listaAnioa.Cabeza
							for impa != nil {
								if impa.Anio == anio {
									impr := impa.ListaMatricesMes.Cabeza
									for impr != nil {
										if impr.Mes == mes {
											nombreProducto := inOrdenNombre(imp.Inventario.Raiz, Productos[j].Codigo)
											nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, NombreTienda, Departamento,Calificacion, nombreProducto,Productos[j].Codigo, 0, strconv.Itoa(dia))
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
				impa = impa.Siguiente
			}
		}
	}
	w.Header().Set("Content-type", "application/json")
	if list.Cabeza == nil{
		mensaje := Mensaje{"NO SE HA PODIDO CARGAR EL ARCHIVO"}
		w.WriteHeader(http.StatusFailedDependency)
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
	indexHandler(w, r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		mensaje := Mensaje{reportes.Arreglo(Vector)}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func tiendaEspecifica (w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)

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
	indexHandler(w, r)
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

func busquedaProductosTienda (w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	vars := mux.Vars(r)
	indi := list.Indi()
	departa := list.Departa()
	var nodosReg []NodoProductoReg
	var InventarioTienda []InventarioReg
	datos := strings.Split(vars["infoTienda"], "&")
	posicion,_ := strconv.Atoi(datos[2])
	Tercero := posicionTercero(datos[1], datos[0], posicion, indi, departa)
		imp := Vector[Tercero].ListGA.Cabeza
		for imp != nil {
			if imp.NombreTienda == datos[1]{
				a := imp.Inventario
				arregloProductos = nil
				inOrdenNombreRegresa(a.Raiz)
				for i := 0; i < len(arregloProductos); i++ {
					n := arregloProductos[i]
					nodosReg = append(nodosReg, NodoProductoReg{NombreProducto: n.NombreProducto, Codigo: n.Codigo, Descripcion: n.Descripcion, PrecioP: n.Precio, Cantidad: n.Cantidad, Imagen: n.Imagen})
				}
				nodoTienda := InventarioReg{NombreTienda: imp.NombreTienda, Departamento: Vector[Tercero].Departamento, Calificacion: imp.Calificacion, Productos: nodosReg}
				InventarioTienda = append(InventarioTienda, nodoTienda)
			}
			imp = imp.Siguiente
		}
		//generalRE := GeneralReg{InventarioTienda}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(InventarioTienda)
}

func eliminarTienda (w http.ResponseWriter, r * http.Request) {
	indexHandler(w, r)
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
	indexHandler(w, r)
	var tiendasRE []TiendaR
	var DepartamentosRE []DepartamentoR
	var datosRE []DatosR
	var generalRE GeneralR
	indi := list.Indi()
	departa := list.Departa()
	vector := Vector
		for j := 0; j < len(indi); j++ {
			for k := 0; k < len(departa); k++ {
				for i := 0; i < len(vector); i++ {
					if vector[i].Indice == indi[j] {
						if vector[i].Departamento == departa[k] {
							imp := vector[i].ListGA.Cabeza
							for imp != nil {
								tiendasRE = append(tiendasRE, TiendaR{Nombre: imp.NombreTienda, Descripcion: imp.Descripcion, Contacto: imp.Contacto, Calificacion: imp.Calificacion, Logo: imp.Logo, PosicionVector: i})
								imp = imp.Siguiente
							}
						}
					}
				}
				if tiendasRE != nil {
					DepartamentosRE = append(DepartamentosRE, DepartamentoR{NombreDepa: departa[k], Tiendas: tiendasRE})
				}
				tiendasRE = nil
			}
			if DepartamentosRE != nil{
				datosRE = append(datosRE, DatosR{Indice: indi[j], Departamentos: DepartamentosRE})
			}
			DepartamentosRE = nil
		}
	generalRE = GeneralR{Inicio: datosRE}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(generalRE)

}

func carrito (w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
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
		dia,_ := strconv.Atoi(strings.Split(fecha, "-")[0])
		mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
		anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
		Tienda := a.NombreTienda
		Departamento := a.Departamento
		Calificacion := a.Calificacion
		Productos := a.CodigoProductos
		Tercero := posicionTercero(Tienda, Departamento, Calificacion, Indi, Departa)
		for j := 0; j < len(Productos); j++ {
			imp := Vector[Tercero].ListGA.Cabeza
			CodigoProducto := Productos[j].Codigo
			Cantidad := Productos[j].Cantidad
			for imp != nil{
				if imp.NombreTienda == Tienda && imp.Calificacion == Calificacion {
						existeAnio := EncontrarAnio(&listaAnioa, anio)
						if existeAnio == false {
							var listaMes MatrizDispersa.ListaMes
							nodoAnio := MatrizDispersa.NodoAnio{Anio: anio, ListaMatricesMes: &listaMes}
							listaAnioa.Insertar(&nodoAnio)

						}
						existeMes := EncontrarMes(&listaAnioa, anio, mes)
						if existeMes == false {
							impm := listaAnioa.Cabeza
							for impm != nil {
								if impm.Anio == anio {
									nodoMes := MatrizDispersa.NodoMes{Mes: mes, MatrizMes: &MatrizDispersa.Matriz{Mes: mes, Anio: anio}}
									impm.ListaMatricesMes.Insertar(&nodoMes)
								}
								impm = impm.Siguiente
							}
						}

						arbolTienda := imp.Inventario.Raiz
						arbolito := Compras.DescontarProducto(arbolTienda, CodigoProducto, Cantidad)
						imp.Inventario.Raiz = arbolito
						existeProducto := inOrden(imp.Inventario.Raiz, Productos[j].Codigo)
						if existeProducto == true {
							impa := listaAnioa.Cabeza
							for impa != nil {
								if impa.Anio == anio {
									impr := impa.ListaMatricesMes.Cabeza
									for impr != nil {
										if impr.Mes == mes {
											nombreProducto := inOrdenNombre(imp.Inventario.Raiz, Productos[j].Codigo)
											nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, Tienda, Departamento,Calificacion, nombreProducto,Productos[j].Codigo, Cantidad, strconv.Itoa(dia))
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
			impm := impa.ListaMatricesMes.Cabeza
			for impm != nil{
				//fmt.Println()
				//fmt.Println("---------------------------------------------------------------------------------------------------------")
				//fmt.Println("---------------------------------------------------------------------------------------------------------")
				//fmt.Println()
				//impm.MatrizMes.DibujarMatriz()
				impm = impm.Siguiente
			}
			impa = impa.Siguiente
		}
		mensaje := Mensaje{Retorna: "¡PRODUCTOS VENDIDOS!"}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func pedidoMes(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	var regresa MatrizDispersa.GeneralInfo
	vars := mux.Vars(r)
	datos := strings.Split(vars["infoPedido"], "&")
	anio,_ := strconv.Atoi(datos[0])
	mes,_ := strconv.Atoi(datos[1])
	dia := datos[2]
	imp := listaAnioa.Cabeza
	for imp != nil{
		if imp.Anio == anio {
			impm := imp.ListaMatricesMes.Cabeza
			for impm != nil{
				if impm.Mes == mes {
					regresa = impm.MatrizMes.Imprimir(dia)
				}
				impm = impm.Siguiente
			}
		}
		imp = imp.Siguiente
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regresa)
}

func datosMatriz(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	var anios []AnioReg
	if listaAnioa.Cabeza != nil {
		listaAnioa = *metodosMatriz.BurbujaAnio(listaAnioa)
		impa := listaAnioa.Cabeza
		for impa != nil{
			impa.ListaMatricesMes = metodosMatriz.BurbujaMes(*impa.ListaMatricesMes)
			impm := impa.ListaMatricesMes.Cabeza
			var a []Mes
			for impm != nil{
				x := Mes{MesA: impm.Mes}
				a = append(a, x)
				impm = impm.Siguiente
			}
			anio := AnioReg{Anio: impa.Anio, Meses: a}
			anios = append(anios, anio)
			impa = impa.Siguiente
		}
	}
	generalAnios :=GeneralReg{anios}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(generalAnios)
}

func imagenMatriz(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	vars := mux.Vars(r)
	datos := strings.Split(vars["datos"], "&")
	anio,_ := strconv.Atoi(datos[0])
	mes,_ := strconv.Atoi(datos[1])
	imp := listaAnioa.Cabeza
	var existe = false
	for imp != nil{
		if imp.Anio == anio {
			impm := imp.ListaMatricesMes.Cabeza
			for impm != nil{
				if impm.Mes == mes{
					impm.MatrizMes.DibujarMatriz()
					existe = true
				}
				impm = impm.Siguiente
			}
		}
		imp = imp.Siguiente
	}
	if existe == true {
		mensaje := Mensaje{Retorna: "EDD_VirtualMall_201807079/Matriz.png"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func arbolTienda(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	vars := mux.Vars(r)
	indi := list.Indi()
	departa := list.Departa()
	datos := strings.Split(vars["datos"], "&")
	posicion,_ := strconv.Atoi(datos[2])
	Tercero := posicionTercero(datos[1], datos[0], posicion, indi, departa)
	imp := Vector[Tercero].ListGA.Cabeza
	for imp != nil {
		if imp.NombreTienda == datos[1]{
			a := imp.Inventario
			a.Generar()
		}
		imp = imp.Siguiente
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("")
}

//----------------------------------------------------------------------------------------------------------------------

func remove(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

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

func inOrdenNombre(raiz *Inventario.NodoArbol, codigo int) string{
	if raiz!=nil {
		if raiz.Codigo == codigo {
			return raiz.NombreProducto
		}
		a := inOrdenNombre(raiz.Izq, codigo)
		if a != "" {
			return a
		}
		b := inOrdenNombre(raiz.Der, codigo)
		if b != "" {
			return b
		}
	}
	return ""
}

func inOrdenNombreRegresa(raiz *Inventario.NodoArbol){
	if raiz!=nil {
		inOrdenNombreRegresa(raiz.Izq)
		nodoIng := Inventario.NodoArbol{NombreProducto: raiz.NombreProducto, Codigo: raiz.Codigo, Factor: raiz.Factor, Cantidad: raiz.Cantidad, Descripcion: raiz.Descripcion, Imagen: raiz.Imagen, Precio: raiz.Precio}
		arregloProductos = append(arregloProductos, nodoIng)
		inOrdenNombreRegresa(raiz.Der)
	}
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
	PosicionVector int
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

type InventarioReg struct {
	NombreTienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []NodoProductoReg `json:"Productos"`
}

type NodoProductoReg struct {
	NombreProducto string `json:"Nombre"`
	Codigo int `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	PrecioP int `json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string `json:"Imagen"`
}

type GeneralReg struct{
	General []AnioReg
}

type AnioReg struct{
	Anio int `json:"Anio"`
	Meses []Mes `json:"Meses"`
}

type Mes struct{
	MesA int `json:"MesA"`
}