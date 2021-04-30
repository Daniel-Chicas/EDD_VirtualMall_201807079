import {React, useEffect, useState} from 'react'
import {BrowserRouter as Router} from 'react-router-dom'
import NavBar from '../componentes/NavBar'
import MosaicoProductos from './MosaicoProductos'

const axios=require('axios').default

function ProductosTienda() {
    var URLactual = window.location;
    var direccion = URLactual.toString().split("/")
    var direccion2 = direccion[4].split("%20")
    var cadena = ""
    for (let step = 0; step < direccion2.length; step++) {
        if(step === 0){
            cadena = direccion2[step]
        }else{
            cadena = cadena+" "+direccion2[step]
        }
      }
      console.log(cadena)

      const [tiendas, settiendas] = useState([])
      const [loading, setloading] = useState(false)
      useEffect(()=>{
          async function obtener(){
              if(tiendas.length===0){
                  const data = await axios.get('http://localhost:3000/Tienda/'+cadena)
                  settiendas(data.data[0])
                  if (data.data[0].Productos!=null) {
                    setloading(true)
                  }else{
                      alert("Debe cargar un archivo de productos o quizá esta tienda no tiene productos")
                  }
              }
          }
          obtener()
      }) 

    if(loading === false){
        return(
            <div className="ui segment carga">
                <div className="ui active dimmer">
                    <div className="ui text loader">Loading</div>
                </div>
                <p />
            </div>
        )
    }else{
        return (
            <div className="ImportList">
                <Router>
                    <NavBar/>   
                </Router>   
                <MosaicoProductos productos={tiendas}/>
            </div>
        )
    }
}

export default ProductosTienda
