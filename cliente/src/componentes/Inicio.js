import React from 'react'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import Login from '../componentesLogin/Login'
import CrearUsuario from '../componentesLogin/CrearUsuario'

import CargaArchivosjson from '../componentes/CargaArchivosjson'
import ListadoTiendas from '../componentes/ListadoTiendas'
import Productos from '../componentes/ProductosTienda'
import Carrito from '../componentes/CarritoCompras'
import Pedidos from '../componentes/Pedidos'
import Imagen from '../componentes/ImagenMatriz'
import ImagenArbol from '../componentes/VerArbol'
import EliminarUsuario from '../componentes/EliminarUsuario'
import ArbolesUsuario from '../componentes/ArbolesUsuario'
import UsuariosPDF from '../componentes/UsuariosPDF'
import UsuariosTodo from '../componentes/ArbolUsuarioTodo'
import UsuariosMedio from '../componentes/ArbolesMedio'
import Vector from '../componentes/VectorLinealizado'
import Recorrido from '../componentes/Recorrido'

import '../Inicio.css'

function Inicio() {
  return (
    <>
      <Router>
        <Route path="/Login" component={Login} />
        <Route path="/CrearUsuario" component={CrearUsuario}/>
        <Route path="/CargaArchivos" component={CargaArchivosjson} />
        <Route path="/VerTiendas" component={ListadoTiendas}/>
        <Route path="/Productos" component={Productos}/>
        <Route path="/CarritoCompras" component={Carrito}/>
        <Route path="/VerPedidos" component={Pedidos}/>
        <Route path="/VerMatriz" component={Imagen} />
        <Route path="/VerArbol" component={ImagenArbol} />
        <Route path="/EliminarUsuario" component={EliminarUsuario} />
        <Route path="/ArbolesUsuario" component={ArbolesUsuario} />
        <Route path="/UsuariosPDF" component={UsuariosPDF} />
        <Route path="/UsuariosT" component={UsuariosTodo} />
        <Route path="/UsuariosM" component={UsuariosMedio} />
        <Route path="/Vector" component={Vector} />
        <Route path="/Recorrido" component={Recorrido} />
      </Router>
    </>
  )
} 

export default Inicio
