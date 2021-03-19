import React from 'react'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import CargaArchivosjson from './componentes/CargaArchivosjson'
import ListadoTiendas from './componentes/ListadoTiendas'
import NavBar from './componentes/NavBar'
import Productos from './componentes/ProductosTienda'
import Carrito from './componentes/CarritoCompras'
import Pedidos from './componentes/Pedidos'
import Imagen from './componentes/ImagenMatriz'
import ImagenArbol from './componentes/VerArbol'
import './App.css'

function App() {
  return (
  <>
      <Router>
        <NavBar/>
        <Route path="/CargaArchivos" component={CargaArchivosjson} />
        <Route path="/VerTiendas" component={ListadoTiendas}/>
        <Route path="/Productos" component={Productos}/>
        <Route path="/CarritoCompras" component={Carrito}/>
        <Route path="/VerPedidos" component={Pedidos}/>
        <Route path="/VerMatriz" component={Imagen} />
        <Route path="/VerArbol" component={ImagenArbol} />
      </Router>
   </>
  )
}

export default App;
