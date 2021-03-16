import React from 'react'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import CargaArchivosjson from './componentes/CargaArchivosjson'
import ListadoTiendas from './componentes/ListadoTiendas'
import NavBar from './componentes/NavBar'
import Productos from './componentes/ProductosTienda'
import './App.css'

function App() {
  return (
  <>
      <Router>
        <NavBar/>
        <Route path="/CargaArchivos" component={CargaArchivosjson} />
        <Route path="/VerTiendas" component={ListadoTiendas}/>
        <Route path="/Productos" component={Productos}/>
        <Route path="/EliminarTienda" component={ListadoTiendas}/>
        <Route path="/CarritoCompras" component={ListadoTiendas}/>
      </Router>
   </>
  )
}

export default App;
