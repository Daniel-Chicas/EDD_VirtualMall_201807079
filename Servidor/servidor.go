package main

import (
	"./ArbolMerkle"
	"./Comentarios"
	"./Compras"
	"./Grafo"
	"./Inventario"
	"./Listas"
	"./MatrizDispersa"
	"./Reportes"
	"./TiendaEspecifica"
	"./Usuarios"
	"container/list"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)
var CopiasGuardadas ListaGuardar

var ms Listas.General
var nodo Listas.Nodo
var departamentos Listas.Departamentos
var tiendas Listas.Tienda
var List Listas.Lista
var reportes Reportes.Lista
var Vector []Listas.NodoArray
var tiendaEsp TiendaEspecifica.General
var tiendaEl TiendaEspecifica.GeneralEliminar
var buscar TiendaEspecifica.Buscar
var GrafoRe GrafoRecorrido.Archivo
var eliminar TiendaEspecifica.Buscar
var arbol Inventario.General
var nodoArbol Inventario.NodoArbol
var matriz MatrizDispersa.General
var metodosMatriz MatrizDispersa.Matriz
var listaAnioa MatrizDispersa.ListaAnio
var carritojson Compras.General
var arregloProductos []Inventario.NodoArbol
var arbolUsuarios Usuarios.ArbolB
var Usuario = arbolUsuarios.NuevoArbol(5)
var UsuariosEntrada Usuarios.General
var Inicio Usuarios.Inicio
var Comen Comentarios.NodoCom
var archivosCreados = false

var CopiaTiendas = ArbolMerkle.NuevoArbol()
var CopiaProductos = ArbolMerkle.NuevoArbolProducto()
var CopiaPedidos = ArbolMerkle.NuevoArbolPedidos()
var CopiaUsuario = ArbolMerkle.NuevoArbolUsuarios()
var CopiaComentariosTienda = ArbolMerkle.NuevoArbolComentarios()
var CopiaComentariosProducto = ArbolMerkle.NuevoArbolComentariosProducto()


var usuarioLinea int
var nueva = GrafoRecorrido.NuevaListaAdyacencia()
var inicioReco string
var finReco string
var LlaveEncriptar = ""
var tiempo time.Duration

func main(){
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "%x", sha256.Sum256([]byte("1234")))
	LlaveEncriptar = cadena.String()
	Usuario.Insertar(Usuarios.NuevaLlave(1234567890101, "EDD2021", " auxiliar@edd.com", cadena.String(), "Administrador"))
	x := (time.Minute * time.Duration(5)) / time.Duration(1)
	tiempo = x
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

}

func serve(ctx context.Context) (err error) {

	router := mux.NewRouter()

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("Server Started")
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
	router.HandleFunc("/Usuarios", usuarios).Methods("Post")
	router.HandleFunc("/EliminarUsuario", eliminarUsuario).Methods("Post")
	router.HandleFunc("/IniciarSesion", IniciarSesion).Methods("Post")
	router.HandleFunc("/UsuarioLinea", UsuarioLinea).Methods("Post")
	router.HandleFunc("/DatosLinea", DatosLinea).Methods("Get")
	router.HandleFunc("/ArbolesB", GraficosArboles).Methods("Post")
	router.HandleFunc("/CambiarContra", CambiarContra).Methods("Post")
	router.HandleFunc("/Comentarios/{DatosTienda}", AgregarComentarios).Methods("Post")
	router.HandleFunc("/ComentariosT/{DatosTienda}", ObtenerComentarios).Methods("Get")
	router.HandleFunc("/ComentariosP/{DatosTienda}", ObtenerComentariosP).Methods("Get")
	router.HandleFunc("/HacerArboles", HacerArboles).Methods("Get")
	router.HandleFunc("/VerificarArboles", VerificarArboles).Methods("Get")
	//log.Fatal(http.ListenAndServe(":3000", router))
	go bloques()


	<-ctx.Done()

	FinAplicacion()
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}

	return
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

	if ms.Inicio == nil {
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
					fmt.Fprint(w, List.Insertar(&nuevo))
					var Hash strings.Builder
					fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(nuevo.Indice+nuevo.Departamento.NombreDepartamento+nuevo.Departamento.Tienda.NombreTienda+nuevo.Departamento.Tienda.Descripcion+nuevo.Departamento.Tienda.Contacto+strconv.Itoa(nuevo.Departamento.Tienda.Calificacion)+nuevo.Departamento.Tienda.Logo)))
					CopiaTiendas.Insertar(Hash.String(), "Crear", nuevo.Indice,nuevo.Departamento.NombreDepartamento, c.Nombre, c.Descripcion, c.Contacto,c.Calificacion,c.Logo)
				}
			}
		}
	}

	Vector = List.CrearMatriz()
	if archivosCreados == false {
		reportes.Arreglo(Vector)
		archivosCreados = true
	}
	Indi := List.Indi()
	Departa := List.Departa()

	var generalReg []Usuarios.General
	var usuariosEx []Usuarios.Usuario
	if UsuariosEntrada.Usuarios == nil{
		json.Unmarshal(reqBody, &UsuariosEntrada)
		for i := 0; i < len(UsuariosEntrada.Usuarios); i++ {
			a := UsuariosEntrada.Usuarios[i]
			existe := existeB(Usuario.Raiz, a.DPI)
			if existe == false {
				var cadena strings.Builder
				fmt.Fprintf(&cadena, "%x", sha256.Sum256([]byte(a.Contra)))
				Usuario.Insertar(Usuarios.NuevaLlave(a.DPI, a.Nombre, a.Correo, cadena.String(), a.Cuenta))
				var Hash strings.Builder
				fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(strconv.Itoa(a.DPI)+a.Nombre+a.Correo+a.Contra)))
				CopiaUsuario.Insertar(Hash.String(), "Crear", a.DPI, a.Nombre, a.Correo, cadena.String(), a.Cuenta)
			}else{
				usuariosEx = append(usuariosEx, a)
			}
		}
		usuarioGen := Usuarios.General{Usuarios: usuariosEx}
		generalReg = append(generalReg, usuarioGen)
	}

	if len(Vector) != 0 {
		if arbol.Inventarios == nil{
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
							nodoArbol.Almacenamiento = a.Almacenamiento
							arbolPosicion.Insertar(nodoArbol.NombreProducto, nodoArbol.Codigo, nodoArbol.Descripcion, nodoArbol.Precio, nodoArbol.Cantidad, nodoArbol.Imagen, nodoArbol.Almacenamiento)
							var Hash strings.Builder
							var x strings.Builder
							fmt.Fprintf(&x, "%x", a.PrecioP)
							fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(NombreTienda+Departamento+strconv.Itoa(Calificacion)+a.NombreProducto+strconv.Itoa(a.Codigo)+a.Descripcion+x.String()+strconv.Itoa(a.Cantidad)+a.Imagen+a.Almacenamiento)))
							CopiaProductos.Insertar(Hash.String(), "Crear", NombreTienda, Departamento, Calificacion, a.NombreProducto, a.Codigo, a.Descripcion, a.PrecioP, a.Cantidad, a.Imagen, a.Almacenamiento)
						}
						imp.Inventario = *arbolPosicion
					}
					imp = imp.Siguiente
				}
			}
		}
		if matriz.Pedidos == nil{
			json.Unmarshal(reqBody, &matriz)
			for i := 0; i < len(matriz.Pedidos); i++ {
				fecha := matriz.Pedidos[i].Fecha
				dia,_ := strconv.Atoi(strings.Split(fecha, "-")[0])
				mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
				anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
				NombreTienda := matriz.Pedidos[i].NombreTienda
				Departamento := matriz.Pedidos[i].Departamento
				Calificacion := matriz.Pedidos[i].Calificacion
				Cliente := matriz.Pedidos[i].Cliente
				Productos := matriz.Pedidos[i].Productos
				Tercero := posicionTercero(NombreTienda, Departamento,Calificacion, Indi, Departa)
				dpi := validarDPI(Usuario.Raiz, Cliente)

				var recorridoFinal GrafoRecorrido.ListaRecorrido
				var recorridos []*GrafoRecorrido.ListaRecorrido
				if dpi == true{
					if len(Productos) > 1 {
						for j := 0; j < len(Productos); j++ {
							if Tercero<0{
								break
							}
							imp := Vector[Tercero].ListGA.Cabeza
							for imp != nil {
								if imp.NombreTienda == NombreTienda && imp.Calificacion == Calificacion {
									a := inOrden(imp.Inventario.Raiz, Productos[j].Codigo)
									if a == true {
										finalRecorrido := inOrdenAlmacenamiento(imp.Inventario.Raiz, Productos[j].Codigo)
										recorrido := nueva.Dijkstra(inicioReco, finalRecorrido,GrafoRe.General)
										recorridos = append(recorridos, recorrido )
									}
								}
								imp = imp.Siguiente
							}
						}
						var temp GrafoRecorrido.NodoRecorrido
						var final GrafoRecorrido.NodoRecorrido
						inicioR := inicioReco
						for len(recorridos) != 0{

							for k := 0; k < len(recorridos); k++ {
								imp := recorridos[k].Cabeza
								recorridos[k] = nueva.Dijkstra(inicioR, recorridos[k].Cola.Va,GrafoRe.General)
								var cuenta float64 = 0
								for imp != nil{
									cuenta = cuenta+imp.Costo
									temp = GrafoRecorrido.NodoRecorrido{Viene: inicioR, Va: imp.Va, Costo: cuenta, Siguiente: nil, Anterior: nil}
									imp = imp.Siguiente
								}
								if k == 0 {
									final = temp
								}else{
									if temp.Costo < final.Costo {
										final = temp
									}
								}
							}

							recorridoFin := nueva.Dijkstra(inicioR, final.Va, GrafoRe.General)


							for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
								recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
							}
							for l := 0; l < len(recorridos); l++ {
								if recorridos[l].Cabeza.Viene == inicioR && recorridos[l].Cola.Va == final.Va {
									recorridos = removerRecorrido(recorridos, l)
									l = 0
								}
							}

							inicioR = final.Va

							if len(recorridos) == 0 {
								recorridoFin = nueva.Dijkstra(inicioR, finReco, GrafoRe.General)
								for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
									recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
								}
								recorridoFin = nueva.Dijkstra(finReco, inicioReco, GrafoRe.General)
								for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
									recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
								}
							}
						}
					}
				}



				if dpi == true {
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
													arbolito := Compras.DescontarProducto(imp.Inventario.Raiz, Productos[j].Codigo, 1)
													imp.Inventario.Raiz = arbolito
													producto := inOrdenDatos(imp.Inventario.Raiz, Productos[j].Codigo)
													if len(Productos) == 1 {
														//mas = false
														finalRecorrido := inOrdenAlmacenamiento(imp.Inventario.Raiz, Productos[j].Codigo)
														recorrido := nueva.Dijkstra(inicioReco, finalRecorrido,GrafoRe.General)
														recorridoFin := nueva.Dijkstra(finalRecorrido, finReco, GrafoRe.General)
														for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
															recorrido.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
														}
														nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, NombreTienda, Departamento,Calificacion, Cliente, nombreProducto,Productos[j].Codigo, 1, strconv.Itoa(dia), recorrido)
														impr.MatrizMes.Insertar(nodoPedido)
														var Hash strings.Builder
														fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(fecha+strconv.Itoa(Cliente)+NombreTienda+Departamento+strconv.Itoa(Calificacion)+strconv.Itoa(producto.Codigo))))
														var recorridoPeido []GrafoRecorrido.NodoRecorrido
														imprp := recorrido.Cabeza
														for imprp != nil{
															nuevo := GrafoRecorrido.NodoRecorrido{Viene: imprp.Viene, Va: imprp.Va, Costo: imprp.Costo, Siguiente: nil, Anterior: nil}
															recorridoPeido = append(recorridoPeido, nuevo)
															imprp = imprp.Siguiente
														}
														CopiaPedidos.Insertar(Hash.String(), "Crear", fecha, NombreTienda, Departamento, Calificacion, Cliente, producto.Codigo, 1, recorridoPeido)
													}else{
														nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, NombreTienda, Departamento,Calificacion, Cliente, nombreProducto,Productos[j].Codigo, 1, strconv.Itoa(dia), &recorridoFinal)
														impr.MatrizMes.Insertar(nodoPedido)
														var Hash strings.Builder
														fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(fecha+strconv.Itoa(Cliente)+NombreTienda+Departamento+strconv.Itoa(Calificacion)+strconv.Itoa(producto.Codigo))))
														var recorridoPeido []GrafoRecorrido.NodoRecorrido
														imprp := recorridoFinal.Cabeza
														for imprp != nil{
															nuevo := GrafoRecorrido.NodoRecorrido{Viene: imprp.Viene, Va: imprp.Va, Costo: imprp.Costo, Siguiente: nil, Anterior: nil}
															recorridoPeido = append(recorridoPeido, nuevo)
															imprp = imprp.Siguiente
														}
														CopiaPedidos.Insertar(Hash.String(), "Crear", fecha, NombreTienda, Departamento, Calificacion, Cliente, producto.Codigo,1, recorridoPeido)
													}
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
	}

	if GrafoRe.General == nil {
		json.Unmarshal(reqBody, &GrafoRe)
		if inicioReco != GrafoRe.PosicionInicialRobot {
			b := GrafoRe.PosicionInicialRobot
			c := GrafoRe.Entrega
			inicioReco = b
			finReco = c
			nueva.Insertar(b, 0)
			nueva.Insertar(c, 0)
			for i := 0; i < len(GrafoRe.General); i++ {
				a := GrafoRe.General[i]
				nueva.Insertar(a.Nombre, 0)
			}
			for i := 0; i < len(GrafoRe.General); i++ {
				a := GrafoRe.General[i]
				for j := 0; j < len(a.Enlaces); j++ {
					nueva.Enlazar(a.Nombre,a.Enlaces[j].Nombre)
				}
			}
			nueva.Dibujar(b,c,GrafoRe.General, &GrafoRecorrido.ListaRecorrido{})
		}
	}

	w.Header().Set("Content-type", "application/json")
	if List.Cabeza == nil{
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

	Indi := List.Indi()
	Departa := List.Departa()
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
	indi := List.Indi()
	departa := List.Departa()
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
					nodosReg = append(nodosReg, NodoProductoReg{NombreProducto: n.NombreProducto, Codigo: n.Codigo, Descripcion: n.Descripcion, PrecioP: n.Precio, Cantidad: n.Cantidad, Imagen: n.Imagen, Almacenamiento: n.Almacenamiento})
				}
				nodoTienda := InventarioReg{NombreTienda: imp.NombreTienda, Departamento: Vector[Tercero].Departamento, Calificacion: imp.Calificacion, Productos: nodosReg}
				InventarioTienda = append(InventarioTienda, nodoTienda)
			}
			imp = imp.Siguiente
		}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(InventarioTienda)
}

func eliminarTienda (w http.ResponseWriter, r * http.Request) {
	indexHandler(w, r)
	Indi := List.Indi()
	Departa := List.Departa()
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
	indi := List.Indi()
	departa := List.Departa()
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
	Indi := List.Indi()
	Departa := List.Departa()
	Tercero := 0
	Tienda := ""
	Calificacion := 0
	json.Unmarshal(reqBody, &carritojson)
	var NProductos []MatrizDispersa.NodoProductoViene
	var recorridoFinal GrafoRecorrido.ListaRecorrido
	var recorridos []*GrafoRecorrido.ListaRecorrido
	if len(carritojson.Pedidos) != 1 {
		for i := 0; i < len(carritojson.Pedidos); i++ {
			a := carritojson.Pedidos[i]
			Tienda = a.NombreTienda
			Departamento := a.Departamento
			Calificacion = a.Calificacion
			Productos := a.CodigoProductos
			Tercero = posicionTercero(Tienda, Departamento, Calificacion, Indi, Departa)
			for j := 0; j < len(Productos); j++ {
				nuevo := MatrizDispersa.NodoProductoViene{Codigo: Productos[j].Codigo}
				NProductos = append(NProductos, nuevo)
			}
		}

		if len(NProductos) != 0 {
			for j := 0; j < len(NProductos); j++ {
				if Tercero<0{
					break
				}
				imp := Vector[Tercero].ListGA.Cabeza
				for imp != nil {
					if imp.NombreTienda == Tienda && imp.Calificacion == Calificacion {
						a := inOrden(imp.Inventario.Raiz, NProductos[j].Codigo)
						if a == true {
							finalRecorrido := inOrdenAlmacenamiento(imp.Inventario.Raiz, NProductos[j].Codigo)
							recorrido := nueva.Dijkstra(inicioReco, finalRecorrido,GrafoRe.General)
							recorridos = append(recorridos, recorrido )
						}
					}
					imp = imp.Siguiente
				}
			}

			var temp GrafoRecorrido.NodoRecorrido
			var final GrafoRecorrido.NodoRecorrido
			inicioR := inicioReco
			for len(recorridos) != 0{

				for k := 0; k < len(recorridos); k++ {
					imp := recorridos[k].Cabeza
					recorridos[k] = nueva.Dijkstra(inicioR, recorridos[k].Cola.Va,GrafoRe.General)
					var cuenta float64 = 0
					for imp != nil{
						cuenta = cuenta+imp.Costo
						temp = GrafoRecorrido.NodoRecorrido{Viene: inicioR, Va: imp.Va, Costo: cuenta, Siguiente: nil, Anterior: nil}
						imp = imp.Siguiente
					}
					if k == 0 {
						final = temp
					}else{
						if temp.Costo < final.Costo {
							final = temp
						}
					}
				}

				recorridoFin := nueva.Dijkstra(inicioR, final.Va, GrafoRe.General)

				for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
					recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
				}
				for l := 0; l < len(recorridos); l++ {
					if recorridos[l].Cabeza.Viene == inicioR && recorridos[l].Cola.Va == final.Va {
						recorridos = removerRecorrido(recorridos, l)
						l = 0
					}
				}

				inicioR = final.Va

				if len(recorridos) == 0 {
					recorridoFin = nueva.Dijkstra(inicioR, finReco, GrafoRe.General)
					for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
						recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
					}
					recorridoFin = nueva.Dijkstra(finReco, inicioReco, GrafoRe.General)
					for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
						recorridoFinal.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
					}
				}

			}
		}
	}

	for i := 0; i < len(carritojson.Pedidos); i++ {
		a := carritojson.Pedidos[i]
		fecha := a.Fecha
		dia,_ := strconv.Atoi(strings.Split(fecha, "-")[0])
		mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
		anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
		Tienda = a.NombreTienda
		Departamento := a.Departamento
		Calificacion = a.Calificacion
		Cliente := a.Cliente
		Productos := a.CodigoProductos
		Tercero = posicionTercero(Tienda, Departamento, Calificacion, Indi, Departa)
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

						arbolTiend := imp.Inventario.Raiz
						arbolito := Compras.DescontarProducto(arbolTiend, CodigoProducto, Cantidad)
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
											producto := inOrdenDatos(imp.Inventario.Raiz, Productos[j].Codigo)
											if len(carritojson.Pedidos) == 1 {
												finalRecorrido := inOrdenAlmacenamiento(imp.Inventario.Raiz, Productos[j].Codigo)
												recorrido := nueva.Dijkstra(inicioReco, finalRecorrido,GrafoRe.General)
												recorridoFin := nueva.Dijkstra(finalRecorrido, finReco, GrafoRe.General)
												for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
													recorrido.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
												}
												recorridoFin = nueva.Dijkstra(finReco, inicioReco, GrafoRe.General)
												for e := recorridoFin.Cabeza; e != nil ; e = e.Siguiente {
													recorrido.InsertarRec(&GrafoRecorrido.NodoRecorrido{Viene: e.Viene, Va: e.Va, Costo: e.Costo, Siguiente: nil, Anterior: nil})
												}
												nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, Tienda, Departamento,Calificacion, Cliente, nombreProducto,Productos[j].Codigo, Productos[0].Cantidad, strconv.Itoa(dia), recorrido)
												impr.MatrizMes.Insertar(nodoPedido)

												//Para Pedidos
												var Hash strings.Builder
												fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(fecha+strconv.Itoa(Cliente)+Tienda+Departamento+strconv.Itoa(Calificacion)+strconv.Itoa(producto.Codigo))))
												var recorridoPeido []GrafoRecorrido.NodoRecorrido
												imprp := recorrido.Cabeza
												for imprp != nil{
													nuevo := GrafoRecorrido.NodoRecorrido{Viene: imprp.Viene, Va: imprp.Va, Costo: imprp.Costo, Siguiente: nil, Anterior: nil}
													recorridoPeido = append(recorridoPeido, nuevo)
													imprp = imprp.Siguiente
												}
												CopiaPedidos.Insertar(Hash.String(), "Crear", fecha, Tienda, Departamento, Calificacion, Cliente, producto.Codigo, Productos[0].Cantidad, recorridoPeido)

											}else{
												var recorrido GrafoRecorrido.ListaRecorrido
												nodoPedido := metodosMatriz.NuevoNodoPedido(fecha, Tienda, Departamento,Calificacion, Cliente, nombreProducto,Productos[j].Codigo, Productos[0].Cantidad, strconv.Itoa(dia), &recorrido)
												//Para Pedidos
												var Hash strings.Builder
												fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(fecha+strconv.Itoa(Cliente)+Tienda+Departamento+strconv.Itoa(Calificacion)+strconv.Itoa(producto.Codigo))))
												var recorridoPeido []GrafoRecorrido.NodoRecorrido
												imprp := recorridoFinal.Cabeza
												for imprp != nil{
													nuevo := GrafoRecorrido.NodoRecorrido{Viene: imprp.Viene, Va: imprp.Va, Costo: imprp.Costo, Siguiente: nil, Anterior: nil}
													recorridoPeido = append(recorridoPeido, nuevo)
													imprp = imprp.Siguiente
												}
												CopiaPedidos.Insertar(Hash.String(), "Crear", fecha, Tienda, Departamento, Calificacion, Cliente, producto.Codigo, Productos[0].Cantidad, recorridoPeido)
												if i == len(carritojson.Pedidos)-1 {
													nodoPedido = metodosMatriz.NuevoNodoPedido(fecha, Tienda, Departamento,Calificacion, Cliente, nombreProducto,Productos[j].Codigo, Productos[0].Cantidad, strconv.Itoa(dia), &recorridoFinal)
													impr.MatrizMes.Insertar(nodoPedido)
												}else{
													impr.MatrizMes.Insertar(nodoPedido)
												}
											}
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
					regresa = impm.MatrizMes.Imprimir(dia, Usuario)
					recorrido := impm.MatrizMes.Recorrido(dia)
					nueva.Dibujar(inicioReco, finReco, GrafoRe.General, recorrido)
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
	indi := List.Indi()
	departa := List.Departa()
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

func usuarios(w http.ResponseWriter, r *http.Request){
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
	var generalReg []Usuarios.General
	var usuariosEx []Usuarios.Usuario
	json.Unmarshal(reqBody, &UsuariosEntrada)
	for i := 0; i < len(UsuariosEntrada.Usuarios); i++ {
		a := UsuariosEntrada.Usuarios[i]
		existe := existeB(Usuario.Raiz, a.DPI)
		if existe == false {
			Usuario.Insertar(Usuarios.NuevaLlave(a.DPI, a.Nombre, a.Correo, a.Contra, a.Cuenta))
			var Hash strings.Builder
			fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(strconv.Itoa(a.DPI)+a.Nombre+a.Correo+a.Contra)))
			CopiaUsuario.Insertar(Hash.String(), "Crear", a.DPI, a.Nombre, a.Correo, a.Contra, a.Cuenta)
		}else{
			usuariosEx = append(usuariosEx, a)
		}
	}
	usuarioGen := Usuarios.General{Usuarios: usuariosEx}
	generalReg = append(generalReg, usuarioGen)
	if generalReg[0].Usuarios == nil {
		mensaje := Mensaje{"EL ARCHIVO HA SIDO GUARDADO EXISTOSAMENTE"}
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(mensaje)
	}else{
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(generalReg)
	}
}

func eliminarUsuario(w http.ResponseWriter, r *http.Request){
	var regresa []string
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

	json.Unmarshal(reqBody, &Inicio)
	dpi := validarDPI(Usuario.Raiz, Inicio.Nombre)
	contra := validarContra(Usuario.Raiz, Inicio.Nombre, Inicio.Contra)
	datosUsuario := DatosUsuario(Usuario.Raiz, Inicio.Nombre)
	if dpi == true {
		regresa = append(regresa, "si" )
	}else{
		regresa = append(regresa, "no" )
	}
	if contra == true {
		regresa = append(regresa, "si" )
	}else{
		regresa = append(regresa, "no" )
	}
	if regresa[0] == "si" && regresa[1] == "si" {
		existe := Usuario.ExisteBEliminar(Usuario.Raiz, Inicio.Nombre, Inicio.Contra)
		var Hash strings.Builder
		fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(datosUsuario.DPI+datosUsuario.Nombre+datosUsuario.Correo+datosUsuario.Contra)))
		dpI,_ := strconv.Atoi(datosUsuario.DPI)
		CopiaUsuario.Insertar(Hash.String(), "Eliminar", dpI, datosUsuario.Nombre, datosUsuario.Correo, datosUsuario.Contra, datosUsuario.Cuenta)
		if existe == true {
			regresa = append(regresa, "EL USUARIO HA SIDO ELIMINADO")
		}else{
			regresa = append(regresa, "REVISE LOS DATOS DE LA CUENTA, ESTOS SON INCORRECTOS O EL USUARIO NO EXISTE")
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ms)
	json.NewEncoder(w).Encode(regresa)
}

func IniciarSesion (w http.ResponseWriter, r *http.Request){
	var regresa []string
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
	json.Unmarshal(reqBody, &Inicio)
	dpi := validarDPI(Usuario.Raiz, Inicio.Nombre)
	contra := validarContra(Usuario.Raiz, Inicio.Nombre, Inicio.Contra)

	if dpi == true {
		regresa = append(regresa, "si" )
	}else{
		regresa = append(regresa, "no" )
	}
	if contra == true {
		regresa = append(regresa, "si" )
	}else{
		regresa = append(regresa, "no" )
	}
	if regresa[0] == "si" && regresa[1] == "si" {
		tipo := tipoUsuario(Usuario.Raiz, Inicio.Nombre, Inicio.Contra)
		regresa = append(regresa, tipo)
		usuarioLinea = Inicio.Nombre
	}

	var salir = CerrarSesion{}
	json.Unmarshal(reqBody, &salir)
	if salir.Cerrar == "si"{
		usuarioLinea = 0
	}

	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ms)
	json.NewEncoder(w).Encode(regresa)
}

func UsuarioLinea (w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprint(w, "Error al insertar")
	}
	if usuarioLinea != 0 {
		tipo := regresaUsuario(Usuario.Raiz, usuarioLinea)
		if tipo == "" {
			tipo = "no"
			w.WriteHeader(http.StatusCreated)
			json.Unmarshal(reqBody, &ms)
			json.NewEncoder(w).Encode(tipo)
		}else{
			w.WriteHeader(http.StatusCreated)
			json.Unmarshal(reqBody, &ms)
			json.NewEncoder(w).Encode(tipo)
		}
	}else{
		tipo := "no"
		w.WriteHeader(http.StatusCreated)
		json.Unmarshal(reqBody, &ms)
		json.NewEncoder(w).Encode(tipo)
	}
}

func DatosLinea(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuarioLinea)
}

func GraficosArboles(w http.ResponseWriter, r *http.Request){
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
	var encriptar Encriptar
	json.Unmarshal(reqBody, &encriptar)
	m := Mensaje{}
	if encriptar.LlaveNueva == LlaveEncriptar {
		Usuario.Grafico("No")
		Usuario.Grafico("Si")
		Usuario.Grafico("Medio")
		m = Mensaje{Retorna: "Los archivos han sido creados"}
	}else{
		m = Mensaje{Retorna: "No"}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func CambiarContra(w http.ResponseWriter, r *http.Request){
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
	var encriptar Encriptar
	json.Unmarshal(reqBody, &encriptar)
	m := Mensaje{}
	if encriptar.LlaveAntigua == LlaveEncriptar {
		LlaveEncriptar = encriptar.LlaveNueva
		m = Mensaje{"Si"}
	}else if encriptar.LlaveAntigua != ""{
		m = Mensaje{"No"}
	}
	if encriptar.Tiempo != "0" && encriptar.LlaveAntigua == "" && encriptar.LlaveNueva == ""{
		if s, err := strconv.ParseFloat(encriptar.Tiempo, 64); err == nil {
			var a,b float64 = 1,1
			var aux float64 = 1
			for aux != s {
				aux = a/b
				if aux < s {
					a++
				}else if aux > s {
					a--
					b++
				}
			}
			tiempo = (time.Minute * time.Duration(a)) / time.Duration(b)
			m = Mensaje{"Acepta"}
		}else{
			fmt.Println(err)
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func AgregarComentarios(w http.ResponseWriter, r *http.Request){
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
	m := Mensaje{"No se ha ingresado"}
	vars := mux.Vars(r)
	datos := strings.Split(vars["DatosTienda"], "&")
	Indi := List.Indi()
	Departa := List.Departa()
	json.Unmarshal(reqBody, &Comen)
	Comen.Departamento = datos[0]
	Comen.NombreTienda = datos[1]
	p,_ := strconv.Atoi(datos[2])
	Comen.Calificacion = p
	Tercero := posicionTercero(Comen.NombreTienda, Comen.Departamento, Comen.Calificacion, Indi, Departa)
	Tiendas := Vector[Tercero].ListGA.Cabeza
	if Comen.Producto == -1 {
		for Tiendas != nil {
			if Tiendas.NombreTienda == Comen.NombreTienda && Tiendas.Calificacion == Comen.Calificacion{
				if len(Comen.Comentarios) == 1 {
					Tiendas.Comentarios.Insertar(Comen.Comentarios[0].Dpi, Comen.Comentarios[0].Comentario, Comen.Comentarios[0].Fecha)
					var Hash strings.Builder
					fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(Comen.NombreTienda+Comen.Departamento+strconv.Itoa(Comen.Calificacion)+strconv.Itoa(Comen.Comentarios[0].Dpi)+Comen.Comentarios[0].Comentario+Comen.Comentarios[0].Fecha)))
					CopiaComentariosTienda.Insertar(Hash.String(), "Crear", Comen.NombreTienda,Comen.Departamento,Comen.Calificacion,"",Comen.Comentarios[0].Dpi,Comen.Comentarios[0].Fecha,Comen.Comentarios[0].Comentario)
					m = Mensaje{"Ingresado"}
				}else{
					posicion := 0
					var tabla *Comentarios.TablaHash
					var respondiendo string
					for i := 0; i < len(Comen.Comentarios); i++ {
						if i == 0 {
							posicion = Tiendas.Comentarios.Buscar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							tabla = Tiendas.Comentarios.Arreglo[posicion].Respuestas
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha+"&"
						}else if i < len(Comen.Comentarios)-1{
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha+"&"
							posicion = tabla.Buscar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							temp := tabla.Arreglo[posicion].Respuestas
							tabla = temp
						}else if i == len(Comen.Comentarios)-1{
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha
							tabla.Insertar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							var Hash strings.Builder
							fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(Comen.NombreTienda+Comen.Departamento+strconv.Itoa(Comen.Calificacion)+strconv.Itoa(Comen.Comentarios[i].Dpi)+Comen.Comentarios[i].Comentario+Comen.Comentarios[i].Fecha)))
							CopiaComentariosTienda.Insertar(Hash.String(), "Crear", Comen.NombreTienda,Comen.Departamento,Comen.Calificacion,respondiendo,Comen.Comentarios[i].Dpi,Comen.Comentarios[i].Fecha,Comen.Comentarios[i].Comentario)

							m = Mensaje{"Ingresado"}
						}
					}
				}
			}
			Tiendas = Tiendas.Siguiente
		}
	}else{
		for Tiendas != nil {
			if Tiendas.NombreTienda == Comen.NombreTienda && Tiendas.Calificacion == Comen.Calificacion{
				Arbol := AgregarComentarioProducto(Tiendas.Inventario.Raiz, Comen.Producto)
				if len(Comen.Comentarios) == 1 {
					Arbol.Comentarios.Insertar(Comen.Comentarios[0].Dpi, Comen.Comentarios[0].Comentario, Comen.Comentarios[0].Fecha)
					var Hash strings.Builder
					fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(Comen.NombreTienda+Comen.Departamento+strconv.Itoa(Comen.Calificacion)+strconv.Itoa(Comen.Comentarios[0].Dpi)+Comen.Comentarios[0].Comentario+Comen.Comentarios[0].Fecha)))
					CopiaComentariosProducto.Insertar(Hash.String(), "Crear", Comen.NombreTienda, Comen.Departamento, Comen.Calificacion, Comen.Producto,"", Comen.Comentarios[0].Dpi,Comen.Comentarios[0].Fecha,Comen.Comentarios[0].Comentario)

					m = Mensaje{"Ingresado"}
				}else{
					posicion := 0
					var tabla *Comentarios.TablaHash
					var respondiendo string
					for i := 0; i < len(Comen.Comentarios); i++ {
						if i == 0 {
							posicion = Arbol.Comentarios.Buscar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							tabla = Arbol.Comentarios.Arreglo[posicion].Respuestas
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha+"&"
						}else if i < len(Comen.Comentarios)-1{
							posicion = tabla.Buscar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha+"&"
							temp := tabla.Arreglo[posicion].Respuestas
							tabla = temp
						}else if i == len(Comen.Comentarios)-1{
							respondiendo = respondiendo+strconv.Itoa(Comen.Comentarios[i].Dpi)+"$"+Comen.Comentarios[i].Comentario+"$"+Comen.Comentarios[i].Fecha
							tabla.Insertar(Comen.Comentarios[i].Dpi, Comen.Comentarios[i].Comentario, Comen.Comentarios[i].Fecha)
							var Hash strings.Builder
							fmt.Fprintf(&Hash, "%x", sha256.Sum256([]byte(Comen.NombreTienda+Comen.Departamento+strconv.Itoa(Comen.Calificacion)+strconv.Itoa(Comen.Comentarios[i].Dpi)+Comen.Comentarios[i].Comentario+Comen.Comentarios[i].Fecha)))
							CopiaComentariosProducto.Insertar(Hash.String(), "Crear", Comen.NombreTienda,Comen.Departamento,Comen.Calificacion, Comen.Producto,respondiendo,Comen.Comentarios[i].Dpi,Comen.Comentarios[i].Fecha,Comen.Comentarios[i].Comentario)
							m = Mensaje{"Ingresado"}
						}
					}
				}
			}
			Tiendas = Tiendas.Siguiente
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

var comentarios []ComentariosReg

func ObtenerComentarios(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	var com ComentariosReg
	comentarios = nil
	vars := mux.Vars(r)
	datos := strings.Split(vars["DatosTienda"], "&")
	Indi := List.Indi()
	Departa := List.Departa()
	Comen.Departamento = datos[0]
	Comen.NombreTienda = datos[1]
	p,_ := strconv.Atoi(datos[2])
	Comen.Calificacion = p
	Tercero := posicionTercero(Comen.NombreTienda, Comen.Departamento, Comen.Calificacion, Indi, Departa)
	Tiendas := Vector[Tercero].ListGA.Cabeza
	for Tiendas != nil {
		if Tiendas.NombreTienda == Comen.NombreTienda && Tiendas.Calificacion == Comen.Calificacion{
			for i := 0; i < len(Tiendas.Comentarios.Arreglo); i++ {
				if Tiendas.Comentarios.Arreglo[i] != nil {
					resp := Respuestas(Tiendas.Comentarios.Arreglo[i].Respuestas)
					com = ComentariosReg{DPI: Tiendas.Comentarios.Arreglo[i].DpiPadre, Fecha: Tiendas.Comentarios.Arreglo[i].FechaComentario, Comentario: Tiendas.Comentarios.Arreglo[i].Comentario, Respuestas: *resp}
					comentarios = append(comentarios, com)
				}
			}
		}
		Tiendas = Tiendas.Siguiente
	}
	regresa := General{comentarios}
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(regresa)
	if err != nil {
		fmt.Println(err)
	}
}

func ObtenerComentariosP(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	var com ComentariosReg
	comentarios = nil
	vars := mux.Vars(r)
	datos := strings.Split(vars["DatosTienda"], "&")
	Indi := List.Indi()
	Departa := List.Departa()
	Comen.Departamento = datos[0]
	Comen.NombreTienda = datos[1]
	p,_ := strconv.Atoi(datos[2])
	Comen.Producto,_ = strconv.Atoi(datos[3])
	Comen.Calificacion = p
	Tercero := posicionTercero(Comen.NombreTienda, Comen.Departamento, Comen.Calificacion, Indi, Departa)
	Tiendas := Vector[Tercero].ListGA.Cabeza
	for Tiendas != nil {
		if Tiendas.NombreTienda == Comen.NombreTienda && Tiendas.Calificacion == Comen.Calificacion{
			Arbol := AgregarComentarioProducto(Tiendas.Inventario.Raiz, Comen.Producto)
			for i := 0; i < len(Arbol.Comentarios.Arreglo); i++ {
				if Arbol.Comentarios.Arreglo[i] != nil {
					resp := Respuestas(Arbol.Comentarios.Arreglo[i].Respuestas)
					com = ComentariosReg{DPI: Arbol.Comentarios.Arreglo[i].DpiPadre, Fecha: Arbol.Comentarios.Arreglo[i].FechaComentario, Comentario: Arbol.Comentarios.Arreglo[i].Comentario, Respuestas: *resp}
					comentarios = append(comentarios, com)
				}
			}
		}
		Tiendas = Tiendas.Siguiente
	}
	regresa := General{comentarios}
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(regresa)
	if err != nil {
		fmt.Println(err)
	}
}

func HacerArboles(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	if CopiasGuardadas.Cola != nil {
		CopiaTiendas.Generar()
		CopiaProductos.Generar()
		CopiaPedidos.Generar()
		CopiaUsuario.Generar()
		CopiaComentariosTienda.Generar()
		CopiaComentariosProducto.Generar()
		mensaje := Mensaje{"si"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mensaje)
	}else{
		mensaje := Mensaje{"NO SE HA CREADO NINGUNA COPIA"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mensaje)
	}
}

func VerificarArboles(w http.ResponseWriter, r *http.Request){
	indexHandler(w, r)
	if CopiasGuardadas.Cola != nil {
		listaT := list.New()
		ListaP := list.New()
		ListaPP := list.New()
		ListaU := list.New()
		ListaCT := list.New()
		ListaCP := list.New()

		listaT = CopiaTiendas.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz, listaT)
		ListaP = CopiaProductos.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz, ListaP)
		ListaPP = CopiaPedidos.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz, ListaPP)
		ListaU = CopiaUsuario.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz, ListaU)
		ListaCT = CopiaComentariosTienda.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz, ListaCT)
		ListaCP = CopiaComentariosProducto.Arreglar(CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz, ListaCP)

		CopiaTiendas.ConstruirArbol(listaT)
		CopiaProductos.ConstruirArbol(ListaP)
		CopiaPedidos.ConstruirArbol(ListaPP)
		CopiaUsuario.ConstruirArbol(ListaU)
		CopiaComentariosTienda.ConstruirArbol(ListaCT)
		CopiaComentariosProducto.ConstruirArbol(ListaCP)

		CopiaTiendas.Generar()
		CopiaProductos.Generar()
		CopiaPedidos.Generar()
		CopiaUsuario.Generar()
		CopiaComentariosTienda.Generar()
		CopiaComentariosProducto.Generar()

		mensaje := Mensaje{"si"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mensaje)
	}else{
		mensaje := Mensaje{"NO SE HA CREADO NINGUNA COPIA"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mensaje)
	}
}

var guarda = false

func bloques () {
	if _, err := os.Stat(".\\bloques"); os.IsNotExist(err) {
		err = os.Mkdir(".\\bloques", 0755)
		if err != nil {
			panic(err)
		}
	}

	CargaArchivosInicio()
	CopiaUsuario.Generar()
	t := time.Now()
	fecha := fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	cuenta := 0
	for i := CopiasGuardadas.Cabeza; i != nil ; i = i.Siguiente {
		cuenta++
	}

	for{
		t = time.Now()
		fecha = fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
		var ht,hp,hpp,hu,hct,hcp = "", "", "", "", "", ""
		if CopiaTiendas.Raiz == nil {ht = ""}else{ht = CopiaTiendas.Raiz.Hash}
		if CopiaProductos.Raiz == nil{hp = ""}else{hp = CopiaProductos.Raiz.Hash}
		if CopiaPedidos.Raiz == nil{hpp = ""}else {hpp = CopiaPedidos.Raiz.Hash}
		if CopiaUsuario.Raiz == nil{hu = ""}else {hu = CopiaUsuario.Raiz.Hash}
		if CopiaComentariosTienda.Raiz == nil{hct = ""}else {hct = CopiaComentariosTienda.Raiz.Hash}
		if CopiaComentariosProducto.Raiz == nil{hcp = ""}else{hcp = CopiaComentariosProducto.Raiz.Hash}

		if CopiasGuardadas.Cabeza == nil  && guarda == true{
			Data := ht + hp + hpp + hu + hct + hcp
			if Data == "" {
				cuenta--
				fmt.Println("Sin Cambios: " + fecha)
			}else {
				t = time.Now()
				fecha = fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
				Hash := strings.Split(PruebaDeTrabajo(cuenta, fecha, Data, "-1"), "&")
				nonce,_ := strconv.Atoi(Hash[1])
				generalCopias := Guardar{HashBloque: Hash[0],Fecha: fecha, PreviousHash: "-1", Indice: cuenta, Data: Data, Nonce: nonce, ClaveArbolB: LlaveEncriptar, CopiaTienda: CopiaTiendas, CopiaProducto: CopiaProductos, CopiaPedidos: CopiaPedidos, CopiaUsuarios: CopiaUsuario, CopiaComentariosTiendas: CopiaComentariosTienda, CopiaComentariosProductos: CopiaComentariosProducto, Grafo: &GrafoRe}
				data, err := json.MarshalIndent(generalCopias, "", "  ")
				if err != nil {
					fmt.Println(err)
				}
				nombreArchivo := "bloques\\"+strconv.Itoa(cuenta)+".json"
				for archivoExiste(nombreArchivo) != false{
					cuenta++
					nombreArchivo = "bloques\\"+strconv.Itoa(cuenta)+".json"
				}
				erro := ioutil.WriteFile(nombreArchivo, data, 0644)
				if erro != nil {
					fmt.Println(erro)
				}
				CopiasGuardadas.Insertar(&NodoGuardar{generalCopias, nil,nil})
				CopiaTiendas = ArbolMerkle.NuevoArbol()
				CopiaProductos = ArbolMerkle.NuevoArbolProducto()
				CopiaPedidos = ArbolMerkle.NuevoArbolPedidos()
				CopiaUsuario = ArbolMerkle.NuevoArbolUsuarios()
				CopiaComentariosTienda = ArbolMerkle.NuevoArbolComentarios()
				CopiaComentariosProducto = ArbolMerkle.NuevoArbolComentariosProducto()
				copia := *CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
				CopiaTiendas.Raiz = &copia
				copiaP := *CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
				CopiaProductos.Raiz = &copiaP
				copiaPp := *CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
				CopiaPedidos.Raiz = &copiaPp
				copiaU := *CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
				CopiaUsuario.Raiz = &copiaU
				copiaCt := *CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
				CopiaComentariosTienda.Raiz = &copiaCt
				copiaCp := *CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
				CopiaComentariosProducto.Raiz = &copiaCp
				fmt.Println("Archivo Creado: "+fecha)
			}
		}else if guarda == true{
			t = time.Now()
			fecha = fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
			if CopiaTiendas.Raiz != nil {
				if CopiaTiendas.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz.Hash {
					ht = ""
				}
			}else{
				CopiaTiendas.Raiz = CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
			}
			if CopiaProductos.Raiz != nil {
				if CopiaProductos.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz.Hash {
					hp = ""
				}
			}else{
				CopiaProductos.Raiz = CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
			}
			if CopiaUsuario.Raiz != nil {
				if CopiaUsuario.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz.Hash {
					hu = ""
				}
			}else{
				CopiaUsuario.Raiz = CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
			}
			if CopiaPedidos.Raiz != nil {
				if CopiaPedidos.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz.Hash {
					hpp = ""
				}
			}else{
				CopiaPedidos.Raiz = CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
			}
			if CopiaComentariosTienda.Raiz != nil {
				if CopiaComentariosTienda.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz.Hash {
					hct = ""
				}
			}else{
				CopiaComentariosTienda.Raiz = CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
			}
			if CopiaComentariosProducto.Raiz != nil {
				if CopiaComentariosProducto.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz.Hash {
					hcp = ""
				}
			}else{
				CopiaComentariosProducto.Raiz = CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
			}

			Data := ht + hp + hpp + hu + hct + hcp
			if Data == "" {
				cuenta--
				fmt.Println("Sin Cambios: " + fecha)
			}else {
				Hash := strings.Split(PruebaDeTrabajo(cuenta, fecha, Data, CopiasGuardadas.Cola.Bloque.HashBloque), "&")
				nonce, _ := strconv.Atoi(Hash[1])
				generalCopias := Guardar{HashBloque: Hash[0], Fecha: fecha, PreviousHash: CopiasGuardadas.Cola.Bloque.HashBloque, Indice: cuenta, Data: Data, Nonce: nonce, CopiaTienda: CopiaTiendas, CopiaProducto: CopiaProductos, CopiaPedidos: CopiaPedidos, CopiaUsuarios: CopiaUsuario, CopiaComentariosTiendas: CopiaComentariosTienda, CopiaComentariosProductos: CopiaComentariosProducto, Grafo: &GrafoRe}
				data, err := json.MarshalIndent(generalCopias, "", "  ")
				if err != nil {
					fmt.Println(err)
				}
				nombreArchivo := "bloques\\" + strconv.Itoa(cuenta) + ".json"
				for archivoExiste(nombreArchivo) != false {
					cuenta++
					nombreArchivo = "bloques\\" + strconv.Itoa(cuenta) + ".json"
				}
				erro := ioutil.WriteFile(nombreArchivo, data, 0644)
				if erro != nil {
					fmt.Println(erro)
				}
				CopiasGuardadas.Insertar(&NodoGuardar{generalCopias, nil, nil})
				CopiaTiendas = ArbolMerkle.NuevoArbol()
				CopiaProductos = ArbolMerkle.NuevoArbolProducto()
				CopiaPedidos = ArbolMerkle.NuevoArbolPedidos()
				CopiaUsuario = ArbolMerkle.NuevoArbolUsuarios()
				CopiaComentariosTienda = ArbolMerkle.NuevoArbolComentarios()
				CopiaComentariosProducto = ArbolMerkle.NuevoArbolComentariosProducto()
				copia := *CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
				CopiaTiendas.Raiz = &copia
				copiaP := *CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
				CopiaProductos.Raiz = &copiaP
				copiaPp := *CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
				CopiaPedidos.Raiz = &copiaPp
				copiaU := *CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
				CopiaUsuario.Raiz = &copiaU
				copiaCt := *CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
				CopiaComentariosTienda.Raiz = &copiaCt
				copiaCp := *CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
				CopiaComentariosProducto.Raiz = &copiaCp
				fmt.Println("Archivo Creado: "+fecha)
			}
		}
		time.Sleep(tiempo)
		guarda = true
		cuenta++
	}
}

func archivoExiste(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func PruebaDeTrabajo(indice int, Fecha string, Data string, PreviousHash string) string{
	nonce := 0
	for{
		var cadena strings.Builder
		fmt.Fprintf(&cadena, "%x", sha256.Sum256([]byte(strconv.Itoa(indice)+Fecha+Data+PreviousHash+strconv.Itoa(nonce))))
		if strings.HasPrefix(cadena.String(), "0000") {
			return cadena.String()+"&"+strconv.Itoa(nonce)
		}
		nonce++
	}
}

func CargaArchivosInicio(){
	contadorarchivos := 1

	t := time.Now()
	fecha := fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("Aplicación inicia: "+fecha)

	for{
		archivo,err := ioutil.ReadFile("bloques\\"+strconv.Itoa(contadorarchivos+1)+".json")
		if err != nil {
			archivo,err = ioutil.ReadFile("bloques\\"+strconv.Itoa(contadorarchivos)+".json")
			if err != nil {
				fmt.Println(err)
				break
			}
			generalRegresa := Guardar{}
			err = json.Unmarshal(archivo, &generalRegresa)
			datosTienda(generalRegresa.CopiaTienda.Raiz)
			Vector = List.CrearMatriz()
			//reportes.Arreglo(Vector)
			LlaveEncriptar = generalRegresa.ClaveArbolB
			GrafoRe = *generalRegresa.Grafo
			if inicioReco != GrafoRe.PosicionInicialRobot {
				b := GrafoRe.PosicionInicialRobot
				c := GrafoRe.Entrega
				inicioReco = b
				finReco = c
				nueva.Insertar(b, 0)
				nueva.Insertar(c, 0)
				for i := 0; i < len(GrafoRe.General); i++ {
					a := GrafoRe.General[i]
					nueva.Insertar(a.Nombre, 0)
				}
				for i := 0; i < len(GrafoRe.General); i++ {
					a := GrafoRe.General[i]
					for j := 0; j < len(a.Enlaces); j++ {
						nueva.Enlazar(a.Nombre,a.Enlaces[j].Nombre)
					}
				}
				nueva.Dibujar(b,c,GrafoRe.General, &GrafoRecorrido.ListaRecorrido{})
			}
			datosProductos(generalRegresa.CopiaProducto.Raiz)
			datosUsuarios(generalRegresa.CopiaUsuarios.Raiz)
			datosPedidos(generalRegresa.CopiaPedidos.Raiz)
			comentariosTiendas(generalRegresa.CopiaComentariosTiendas.Raiz)
			comentariosProductos(generalRegresa.CopiaComentariosProductos.Raiz)
			break
		}

		contadorarchivos++
	}
	contadorarchivos = 1
	for{
		archivo,err := ioutil.ReadFile("bloques\\"+strconv.Itoa(contadorarchivos)+".json")
		if err != nil {
			break
		}
		generalRegresa := Guardar{}
		err = json.Unmarshal(archivo, &generalRegresa)
		CopiasGuardadas.Insertar(&NodoGuardar{generalRegresa, nil,nil})

		CopiaTiendas = ArbolMerkle.NuevoArbol()
		CopiaProductos = ArbolMerkle.NuevoArbolProducto()
		CopiaPedidos = ArbolMerkle.NuevoArbolPedidos()
		CopiaUsuario = ArbolMerkle.NuevoArbolUsuarios()
		CopiaComentariosTienda = ArbolMerkle.NuevoArbolComentarios()
		CopiaComentariosProducto = ArbolMerkle.NuevoArbolComentariosProducto()

		copia := CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
		CopiaTiendas.Raiz = copia
		copiaP := CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
		CopiaProductos.Raiz = copiaP
		copiaPp := CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
		CopiaPedidos.Raiz = copiaPp
		copiaU := CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
		CopiaUsuario.Raiz = copiaU
		copiaCt := CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
		CopiaComentariosTienda.Raiz = copiaCt
		copiaCp := CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
		CopiaComentariosProducto.Raiz = copiaCp
		contadorarchivos++
	}
}

func FinAplicacion(){
	cuenta := 0
	for i := CopiasGuardadas.Cabeza; i != nil ; i = i.Siguiente {
		cuenta++
	}
	t := time.Now()
	fecha := fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	var ht,hp,hpp,hu,hct,hcp = "", "", "", "", "", ""
	if CopiaTiendas.Raiz == nil {ht = ""}else{ht = CopiaTiendas.Raiz.Hash}
	if CopiaProductos.Raiz == nil{hp = ""}else{hp = CopiaProductos.Raiz.Hash}
	if CopiaPedidos.Raiz == nil{hpp = ""}else {hpp = CopiaPedidos.Raiz.Hash}
	if CopiaUsuario.Raiz == nil{hu = ""}else {hu = CopiaUsuario.Raiz.Hash}
	if CopiaComentariosTienda.Raiz == nil{hct = ""}else {hct = CopiaComentariosTienda.Raiz.Hash}
	if CopiaComentariosProducto.Raiz == nil{hcp = ""}else{hcp = CopiaComentariosProducto.Raiz.Hash}
		if CopiaTiendas.Raiz != nil {
			if CopiaTiendas.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz.Hash {
				ht = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaTiendas.Raiz = CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
			}
		}
		if CopiaProductos.Raiz != nil {
			if CopiaProductos.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz.Hash {
				hp = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaProductos.Raiz = CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
			}
		}
		if CopiaUsuario.Raiz != nil {
			if CopiaUsuario.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz.Hash {
				hu = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaUsuario.Raiz = CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
			}
		}
		if CopiaPedidos.Raiz != nil {
			if CopiaPedidos.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz.Hash {
				hpp = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaPedidos.Raiz = CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
			}
		}
		if CopiaComentariosTienda.Raiz != nil {
			if CopiaComentariosTienda.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz.Hash {
				hct = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaComentariosTienda.Raiz = CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
			}
		}
		if CopiaComentariosProducto.Raiz != nil {
			if CopiaComentariosProducto.Raiz.Hash == CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz.Hash {
				hcp = ""
			}
		}else{
			if CopiasGuardadas.Cola != nil{
				CopiaComentariosProducto.Raiz = CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
			}
		}
		Data := ht + hp + hpp + hu + hct + hcp
		if Data == "" {
			cuenta--
			fmt.Println("Sin Cambios: " + fecha)
		}else {
			t = time.Now()
			fecha = fmt.Sprintf("%2d-%02d-%2d::%02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
			Hash := strings.Split(PruebaDeTrabajo(cuenta, fecha, Data, "-1"), "&")
			nonce,_ := strconv.Atoi(Hash[1])
			generalCopias := Guardar{HashBloque: Hash[0],Fecha: fecha, PreviousHash: "-1", Indice: cuenta, Data: Data, Nonce: nonce, ClaveArbolB: LlaveEncriptar, CopiaTienda: CopiaTiendas, CopiaProducto: CopiaProductos, CopiaPedidos: CopiaPedidos, CopiaUsuarios: CopiaUsuario, CopiaComentariosTiendas: CopiaComentariosTienda, CopiaComentariosProductos: CopiaComentariosProducto, Grafo: &GrafoRe}
			data, err := json.MarshalIndent(generalCopias, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			nombreArchivo := "bloques\\"+strconv.Itoa(cuenta)+".json"
			for archivoExiste(nombreArchivo) != false{
				cuenta++
				nombreArchivo = "bloques\\"+strconv.Itoa(cuenta)+".json"
			}
			erro := ioutil.WriteFile(nombreArchivo, data, 0644)
			if erro != nil {
				fmt.Println(erro)
			}
			CopiasGuardadas.Insertar(&NodoGuardar{generalCopias, nil,nil})
			CopiaTiendas = ArbolMerkle.NuevoArbol()
			CopiaProductos = ArbolMerkle.NuevoArbolProducto()
			CopiaPedidos = ArbolMerkle.NuevoArbolPedidos()
			CopiaUsuario = ArbolMerkle.NuevoArbolUsuarios()
			CopiaComentariosTienda = ArbolMerkle.NuevoArbolComentarios()
			CopiaComentariosProducto = ArbolMerkle.NuevoArbolComentariosProducto()
			copia := *CopiasGuardadas.Cola.Bloque.CopiaTienda.Raiz
			CopiaTiendas.Raiz = &copia
			copiaP := *CopiasGuardadas.Cola.Bloque.CopiaProducto.Raiz
			CopiaProductos.Raiz = &copiaP
			copiaPp := *CopiasGuardadas.Cola.Bloque.CopiaPedidos.Raiz
			CopiaPedidos.Raiz = &copiaPp
			copiaU := *CopiasGuardadas.Cola.Bloque.CopiaUsuarios.Raiz
			CopiaUsuario.Raiz = &copiaU
			copiaCt := *CopiasGuardadas.Cola.Bloque.CopiaComentariosTiendas.Raiz
			CopiaComentariosTienda.Raiz = &copiaCt
			copiaCp := *CopiasGuardadas.Cola.Bloque.CopiaComentariosProductos.Raiz
			CopiaComentariosProducto.Raiz = &copiaCp
			fmt.Println("Archivo Creado"+fecha)
		}
}

//-------------------------------------------------MÉTODOS--------------------------------------------------------------

func remove(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func removerRecorrido(slice []*GrafoRecorrido.ListaRecorrido, s int) []*GrafoRecorrido.ListaRecorrido {
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

func AgregarComentarioProducto(raiz *Inventario.NodoArbol, codigo int) *Inventario.NodoArbol{
	if raiz!=nil {
		if raiz.Codigo == codigo {
			return raiz
		}
		a := AgregarComentarioProducto(raiz.Izq, codigo)
		if a != nil {
			return a
		}
		b := AgregarComentarioProducto(raiz.Der, codigo)
		if b != nil {
			return b
		}
	}
	return nil
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

func inOrdenAlmacenamiento(raiz *Inventario.NodoArbol, codigo int) string{
	if raiz!=nil {
		if raiz.Codigo == codigo {
			return raiz.Almacenamiento
		}
		a := inOrdenAlmacenamiento(raiz.Izq, codigo)
		if a != "" {
			return a
		}
		b := inOrdenAlmacenamiento(raiz.Der, codigo)
		if b != "" {
			return b
		}
	}
	return ""
}

func inOrdenNombreRegresa(raiz *Inventario.NodoArbol){
	if raiz!=nil {
		inOrdenNombreRegresa(raiz.Izq)
		nodoIng := Inventario.NodoArbol{NombreProducto: raiz.NombreProducto, Codigo: raiz.Codigo, Factor: raiz.Factor, Cantidad: raiz.Cantidad, Descripcion: raiz.Descripcion, Imagen: raiz.Imagen, Precio: raiz.Precio, Almacenamiento: raiz.Almacenamiento}
		arregloProductos = append(arregloProductos, nodoIng)
		inOrdenNombreRegresa(raiz.Der)
	}
}

func existeB (pagina *Usuarios.Pagina, dpi int) bool{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
					existe = true
					return true
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := existeB(pagina.Llaves[i].Izq, dpi)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := existeB(pagina.Llaves[i].Der, dpi)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := existeB(pagina.Llaves[i].Der, dpi)
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

func validarDPI (pagina *Usuarios.Pagina, dpi int) bool{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {

					existe = true
					return true
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := validarDPI(pagina.Llaves[i].Izq, dpi)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := validarDPI(pagina.Llaves[i].Der, dpi)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := validarDPI(pagina.Llaves[i].Der, dpi)
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

func validarContra (pagina *Usuarios.Pagina, dpi int, contra string) bool{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
					if pagina.Llaves[i].Usuario.Contra == contra {
						existe = true
						return true
					}
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := validarContra(pagina.Llaves[i].Izq, dpi, contra)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := validarContra(pagina.Llaves[i].Der, dpi, contra)
						if a == true {
							return true
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := validarContra(pagina.Llaves[i].Der, dpi, contra)
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

func tipoUsuario (pagina *Usuarios.Pagina, dpi int, contra string) string{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
					if pagina.Llaves[i].Usuario.Contra == contra {
						existe = true
						return pagina.Llaves[i].Usuario.Cuenta
					}
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := tipoUsuario(pagina.Llaves[i].Izq, dpi, contra)
						if a != "" {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := tipoUsuario(pagina.Llaves[i].Der, dpi, contra)
						if a != "" {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := tipoUsuario(pagina.Llaves[i].Der, dpi, contra)
						if a != "" {
							return a
						}
					}
				}
			}
		}else{
			return ""
		}
	}
	return ""
}

func regresaUsuario (pagina *Usuarios.Pagina, dpi int) string{
	existe := false
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
					existe = true
					return pagina.Llaves[i].Usuario.Cuenta
				}
			}
		}
		if existe == false {
			for i := 0; i < len(pagina.Llaves); i++ {
				if pagina.Llaves[i] != nil {
					if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
						a := regresaUsuario(pagina.Llaves[i].Izq, dpi)
						if a != "" {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
						a := regresaUsuario(pagina.Llaves[i].Der, dpi)
						if a != "" {
							return a
						}
					}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
						a := regresaUsuario(pagina.Llaves[i].Der, dpi)
						if a != "" {
							return a
						}
					}
				}
			}
		}else{
			return ""
		}
	}
	return ""
}

func Respuestas(resp *Comentarios.TablaHash) *[]ComentariosReg{
	var res []ComentariosReg
	for i := 0; i < len(resp.Arreglo); i++ {
		if resp.Arreglo[i] != nil {
			if resp.Arreglo[i].Respuestas != nil {
				a := Respuestas(resp.Arreglo[i].Respuestas)
				com := ComentariosReg{DPI: resp.Arreglo[i].DpiPadre, Fecha: resp.Arreglo[i].FechaComentario, Comentario: resp.Arreglo[i].Comentario, Respuestas: *a}
				res = append(res, com)
			}
		}
	}
	return &res
}

func inOrdenDatos(raiz *Inventario.NodoArbol, codigo int) *Inventario.NodoArbol{
	if raiz!=nil {
		if raiz.Codigo == codigo {
			return raiz
		}
		a := inOrdenDatos(raiz.Izq, codigo)
		if a != nil {
			return a
		}
		b := inOrdenDatos(raiz.Der, codigo)
		if b != nil {
			return b
		}
	}
	return nil
}

func DatosUsuario (pagina *Usuarios.Pagina, dpi int) *Usuarios.NodoUsuario{
	if pagina != nil {
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI == strconv.Itoa(dpi) {
					return pagina.Llaves[i].Usuario
				}
			}
		}
		for i := 0; i < len(pagina.Llaves); i++ {
			if pagina.Llaves[i] != nil {
				if pagina.Llaves[i].Usuario.DPI > strconv.Itoa(dpi) && i == 0 {
					a := DatosUsuario(pagina.Llaves[i].Izq, dpi)
					if a != nil {
						return a
					}
				}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi){
					a := DatosUsuario(pagina.Llaves[i].Der, dpi)
					if a != nil {
						return a
					}
				}else if pagina.Llaves[i].Usuario.DPI < strconv.Itoa(dpi) && i == len(pagina.Llaves) {
					a := DatosUsuario(pagina.Llaves[i].Der, dpi)
					if a != nil {
						return a
					}
				}
			}
		}
	}
	return nil
}

func datosTienda(raiz *ArbolMerkle.Nodo){
	if raiz!=nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil && raiz.NombreTienda != ""{
			existe := CopiaTiendas.ExisteTienda(CopiaTiendas.Raiz, raiz.NombreTienda, raiz.Tipo, raiz.Indice, raiz.Departamento, raiz.NombreTienda, raiz.DescripcionTienda, raiz.ContactoTienda, raiz.Calificacion, raiz.Logo)
			if existe == false {
				tienda := Listas.Tiendas{NombreTienda: raiz.NombreTienda, Descripcion: raiz.DescripcionTienda, Contacto: raiz.ContactoTienda, Calificacion: raiz.Calificacion, Logo: raiz.Logo}
				depa := Listas.Departamentos{NombreDepartamento: raiz.Departamento, Tienda: tienda}
				nuevo := Listas.Nodo{Indice: raiz.Indice, Departamento: depa }
				List.Insertar(&nuevo)
				CopiaTiendas.Insertar(raiz.Hash, raiz.Tipo, nuevo.Indice,nuevo.Departamento.NombreDepartamento, raiz.NombreTienda, raiz.DescripcionTienda, raiz.ContactoTienda,raiz.Calificacion,raiz.Logo)
			}
		}
		datosTienda(raiz.Izquierda)
		datosTienda(raiz.Derecha)
	}
}

func datosProductos(raiz *ArbolMerkle.NodoProductos){
	if raiz!=nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil && raiz.NombreProducto != ""{
			//existe := CopiaProductos.ExisteProducto(CopiaProductos.Raiz, raiz.Hash, raiz.Tipo, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.NombreProducto, raiz.Codigo, raiz.Descripcion, raiz.Precio, raiz.Cantidad, raiz.Imagen, raiz.Almacenamiento)
			//if existe == false {
				Indi := List.Indi()
				Departa := List.Departa()
				NombreTienda := raiz.Tienda
				Departamento := raiz.Departamento
				Calificacion := raiz.Calificacion
				Tercero := posicionTercero(NombreTienda, Departamento,Calificacion, Indi, Departa)
				imp := Vector[Tercero].ListGA.Cabeza
				for imp != nil {
					if imp.NombreTienda == NombreTienda{
						var arbolPosicion *Inventario.Arbol
						if imp.Inventario.Raiz == nil {
							arbolPosicion = imp.Inventario.NuevoArbol()
						}else{
							arbolPosicion = &imp.Inventario
						}
						nodoArbol.NombreProducto = raiz.NombreProducto
						nodoArbol.Codigo = raiz.Codigo
						nodoArbol.Descripcion = raiz.Descripcion
						nodoArbol.Precio =  raiz.Precio
						nodoArbol.Cantidad = raiz.Cantidad
						nodoArbol.Imagen = raiz.Imagen
						nodoArbol.Almacenamiento = raiz.Almacenamiento
						arbolPosicion.Insertar(nodoArbol.NombreProducto, nodoArbol.Codigo, nodoArbol.Descripcion, nodoArbol.Precio, nodoArbol.Cantidad, nodoArbol.Imagen, nodoArbol.Almacenamiento)
						CopiaProductos.Insertar(raiz.Hash, raiz.Tipo, NombreTienda, Departamento, Calificacion, raiz.NombreProducto, raiz.Codigo, raiz.Descripcion, raiz.Precio, raiz.Cantidad, raiz.Imagen, raiz.Almacenamiento)
						imp.Inventario = *arbolPosicion
					}
					imp = imp.Siguiente
				}
			//}
		}
		datosProductos(raiz.Izquierda)
		datosProductos(raiz.Derecha)
	}
}

func datosUsuarios(raiz *ArbolMerkle.NodoUsuarios){
	if raiz!=nil {
		if raiz.Izquierda == nil && raiz.Derecha == nil && raiz.Nombre != ""{
			//existeU := CopiaUsuario.ExisteUsuario(CopiaUsuario.Raiz, raiz.Hash, raiz.Tipo, raiz.DPI, raiz.Nombre, raiz.Correo, raiz.Password, raiz.Cuenta )
			//if existeU == false {
				if raiz.Tipo == "Eliminar" {
					Usuario.ExisteBEliminar(Usuario.Raiz, raiz.DPI, raiz.Password)
					CopiaUsuario.Insertar(raiz.Hash, raiz.Tipo, raiz.DPI, raiz.Nombre, raiz.Correo, raiz.Password, raiz.Cuenta)
				}else{
					existe := existeB(Usuario.Raiz, raiz.DPI)
					if existe == false {
						Usuario.Insertar(Usuarios.NuevaLlave(raiz.DPI, raiz.Nombre, raiz.Correo, raiz.Password, raiz.Cuenta))
						CopiaUsuario.Insertar(raiz.Hash, raiz.Tipo, raiz.DPI, raiz.Nombre, raiz.Correo, raiz.Password, raiz.Cuenta)
					}
				}
			//}

		}
		datosUsuarios(raiz.Izquierda)
		datosUsuarios(raiz.Derecha)
	}
}

func datosPedidos(raiz *ArbolMerkle.NodoPedidos){
	if raiz!=nil {
		if raiz.Derecha == nil && raiz.Izquierda == nil && raiz.Tienda != "" {
			//existeP := CopiaPedidos.ExistePedido(CopiaPedidos.Raiz, raiz.Hash, raiz.Tipo, raiz.Fecha, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.Cliente, raiz.Producto, raiz.Cantidad)
			//if existeP == false {
				Indi := List.Indi()
				Departa := List.Departa()
				Tercero := posicionTercero(raiz.Tienda, raiz.Departamento, raiz.Calificacion, Indi, Departa)
				imp := Vector[Tercero].ListGA.Cabeza
				fecha := raiz.Fecha
				dia,_ := strconv.Atoi(strings.Split(fecha, "-")[0])
				mes,_ := strconv.Atoi(strings.Split(fecha, "-")[1])
				anio,_ := strconv.Atoi(strings.Split(fecha, "-")[2])
				for imp != nil{
					if imp.NombreTienda == raiz.Tienda && imp.Calificacion == raiz.Calificacion {
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

						arbolTiend := imp.Inventario.Raiz
						arbolito := Compras.DescontarProducto(arbolTiend, raiz.Producto, raiz.Cantidad)
						imp.Inventario.Raiz = arbolito
						existeProducto := inOrden(imp.Inventario.Raiz, raiz.Producto)
						if existeProducto == true {
							impa := listaAnioa.Cabeza
							for impa != nil {
								if impa.Anio == anio {
									impr := impa.ListaMatricesMes.Cabeza
									for impr != nil {
										if impr.Mes == mes {
											nombreProducto := inOrdenNombre(imp.Inventario.Raiz, raiz.Producto)
											var recorrido GrafoRecorrido.ListaRecorrido
											for i := 0; i < len(raiz.Recorrido); i++ {
												recorrido.InsertarRec(&raiz.Recorrido[i])
											}
											nodoPedido := metodosMatriz.NuevoNodoPedido(raiz.Fecha, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.Cliente, nombreProducto, raiz.Producto, raiz.Cantidad, strconv.Itoa(dia), &recorrido)
											impr.MatrizMes.Insertar(nodoPedido)
											var recorridoPeido []GrafoRecorrido.NodoRecorrido
											imprp := recorrido.Cabeza
											for imprp != nil{
												nuevo := GrafoRecorrido.NodoRecorrido{Viene: imprp.Viene, Va: imprp.Va, Costo: imprp.Costo, Siguiente: nil, Anterior: nil}
												recorridoPeido = append(recorridoPeido, nuevo)
												imprp = imprp.Siguiente
											}
											CopiaPedidos.Insertar(raiz.Hash, "Crear", fecha, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.Cliente, raiz.Producto, raiz.Cantidad, recorridoPeido)

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
			//}
		}
		datosPedidos(raiz.Izquierda)
		datosPedidos(raiz.Derecha)
	}
}

func comentariosTiendas(raiz *ArbolMerkle.NodoComentario){
	if raiz !=nil{
		if raiz.Derecha == nil && raiz.Izquierda == nil && raiz.Tienda != ""{
			//existeCT := CopiaComentariosTienda.ExisteComentarioT(CopiaComentariosTienda.Raiz, raiz.Hash, raiz.Tipo, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.Respondiendo, raiz.Dpi, raiz.Fecha, raiz.Comentario)
			//if existeCT == false {
				Indi := List.Indi()
				Departa := List.Departa()
				Tercero := posicionTercero(raiz.Tienda, raiz.Departamento, raiz.Calificacion, Indi, Departa)
				Tiendas := Vector[Tercero].ListGA.Cabeza
				for Tiendas != nil {
					if Tiendas.NombreTienda == raiz.Tienda && Tiendas.Calificacion == raiz.Calificacion{
						if raiz.Respondiendo == "" {
							Tiendas.Comentarios.Insertar(raiz.Dpi, raiz.Comentario, raiz.Fecha)
							CopiaComentariosTienda.Insertar(raiz.Hash, raiz.Tipo, raiz.Tienda,raiz.Departamento,raiz.Calificacion,raiz.Respondiendo,raiz.Dpi,raiz.Fecha,raiz.Comentario)
						}else{
							resp := strings.Split(raiz.Respondiendo, "&")
							var comentariosJson []Comentarios.Respuestas
							for i := 0; i < len(resp); i++ {
								x := strings.Split(resp[i], "$")
								dpi,_ := strconv.Atoi(x[0])
								comentariosJson = append(comentariosJson, Comentarios.Respuestas{Dpi: dpi, Comentario: x[1], Fecha: x[2]})
							}
							posicion := 0
							var tabla *Comentarios.TablaHash
							for i := 0; i < len(comentariosJson); i++ {
								if i == 0 {
									posicion = Tiendas.Comentarios.Buscar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									tabla = Tiendas.Comentarios.Arreglo[posicion].Respuestas
								}else if i < len(comentariosJson)-1{
									posicion = tabla.Buscar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									temp := tabla.Arreglo[posicion].Respuestas
									tabla = temp
								}else if i == len(comentariosJson)-1{
									tabla.Insertar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									CopiaComentariosTienda.Insertar(raiz.Hash, raiz.Tipo, raiz.Tienda,raiz.Departamento,raiz.Calificacion,raiz.Respondiendo, comentariosJson[i].Dpi,comentariosJson[i].Fecha,comentariosJson[i].Comentario)
								}
							}
						}
					}
					Tiendas = Tiendas.Siguiente
				}
			//}

		}

		 comentariosTiendas(raiz.Izquierda)
		 comentariosTiendas(raiz.Derecha)
	}
}

func comentariosProductos(raiz *ArbolMerkle.NodoComentarioProducto){
	if raiz !=nil{
		if raiz.Derecha == nil && raiz.Izquierda == nil && raiz.Tienda != ""{
			//existeCP := CopiaComentariosProducto.ExisteComentarioP(CopiaComentariosProducto.Raiz, raiz.Hash, raiz.Tipo, raiz.Tienda, raiz.Departamento, raiz.Calificacion, raiz.CodigoProducto, raiz.Respondiendo, raiz.Dpi, raiz.Fecha, raiz.Comentario)
			//if existeCP == false {
				Indi := List.Indi()
				Departa := List.Departa()
				Tercero := posicionTercero(raiz.Tienda, raiz.Departamento, raiz.Calificacion, Indi, Departa)
				Tiendas := Vector[Tercero].ListGA.Cabeza
				for Tiendas != nil {
					if Tiendas.NombreTienda == raiz.Tienda && Tiendas.Calificacion == raiz.Calificacion{
						Arbol := AgregarComentarioProducto(Tiendas.Inventario.Raiz, raiz.CodigoProducto)
						if raiz.Respondiendo == "" {
							Arbol.Comentarios.Insertar(raiz.Dpi, raiz.Comentario, raiz.Fecha)
							CopiaComentariosProducto.Insertar(raiz.Hash, raiz.Tipo, raiz.Tienda,raiz.Departamento,raiz.Calificacion, raiz.CodigoProducto,raiz.Respondiendo,raiz.Dpi,raiz.Fecha,raiz.Comentario)
						}else{
							resp := strings.Split(raiz.Respondiendo, "&")
							var comentariosJson []Comentarios.Respuestas
							for i := 0; i < len(resp); i++ {
								x := strings.Split(resp[i], "$")
								dpi,_ := strconv.Atoi(x[0])
								comentariosJson = append(comentariosJson, Comentarios.Respuestas{Dpi: dpi, Comentario: x[1], Fecha: x[2]})
							}
							posicion := 0
							var tabla *Comentarios.TablaHash
							for i := 0; i < len(comentariosJson); i++ {
								if i == 0 {
									posicion = Arbol.Comentarios.Buscar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									tabla = Arbol.Comentarios.Arreglo[posicion].Respuestas
								}else if i < len(comentariosJson)-1{
									posicion = tabla.Buscar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									temp := tabla.Arreglo[posicion].Respuestas
									tabla = temp
								}else if i == len(comentariosJson)-1{
									tabla.Insertar(comentariosJson[i].Dpi, comentariosJson[i].Comentario, comentariosJson[i].Fecha)
									CopiaComentariosProducto.Insertar(raiz.Hash, raiz.Tipo, raiz.Tienda,raiz.Departamento,raiz.Calificacion, raiz.CodigoProducto,raiz.Respondiendo, comentariosJson[i].Dpi,comentariosJson[i].Fecha,comentariosJson[i].Comentario)
								}
							}
						}
					}
					Tiendas = Tiendas.Siguiente
				}
			//}
		}
		comentariosProductos(raiz.Izquierda)
		comentariosProductos(raiz.Derecha)
	}
}

//--------------------------------------------------ESTRUCTURAS REGRESO-------------------------------------------------

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
	PrecioP float64 `json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string `json:"Imagen"`
	Almacenamiento string `json:"Almacenamiento"`
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

type CerrarSesion struct{
	Cerrar string `json:"Cerrar"`
}

type Encriptar struct {
	LlaveAntigua string `json:"LlaveA"`
	LlaveNueva string `json:"LlaveN"`
	Tiempo string `json:"Tiempo"`
	
}

type General struct{
	Comentarios []ComentariosReg `json:"General"`
}

type ComentariosReg struct {
	DPI int `json:"Dpi"`
	Fecha string `json:"Fecha"`
	Comentario string `json:"Comentario"`
	Respuestas []ComentariosReg `json:"Respuestas"`
}

type Guardar struct {
	HashBloque string `json:"HashBloque"`
	Indice int
	Fecha string
	Data string
	Nonce int
	PreviousHash string
	ClaveArbolB string
	CopiaTienda *ArbolMerkle.Arbol `json:"Tiendas"`
	CopiaProducto *ArbolMerkle.ArbolProductos `json:"Productos"`
	CopiaPedidos *ArbolMerkle.ArbolPedidos `json:"Pedidos"`
	CopiaUsuarios *ArbolMerkle.ArbolUsuarios `json:"Usuarios"`
	CopiaComentariosTiendas *ArbolMerkle.ArbolComentarios `json:"ComentariosT"`
	CopiaComentariosProductos *ArbolMerkle.ArbolComentariosProducto `json:"ComentariosP"`
	Grafo *GrafoRecorrido.Archivo `json:"Grafo"`
}

type NodoGuardar struct{
	Bloque Guardar
	Siguiente *NodoGuardar
	Anterior *NodoGuardar
}

type ListaGuardar struct{
	Cabeza *NodoGuardar
	Cola *NodoGuardar
}

func (L *ListaGuardar) Insertar(nuevo *NodoGuardar) {
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Cola = nuevo
	}else{
		L.Cola.Siguiente = nuevo
		nuevo.Anterior = L.Cola
		L.Cola = nuevo
	}
}