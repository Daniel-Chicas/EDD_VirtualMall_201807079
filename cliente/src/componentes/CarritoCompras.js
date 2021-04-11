import {React, useEffect, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import { Label } from 'semantic-ui-react'
import Tabla from './Tabla'
import '../css/CarritoCompras.css'
import '../Inicio.css'
import axios from 'axios'

function CarritoCompras() {
    const encabezado=['id', 'Tienda', 'Departamento', 'Calificación', 'Nombre Producto', 'Código Producto', 'Precio', 'Disponibilidad en stock', 'Cantidad', 'Total por Producto']
    const [listado, setlistado] = useState([
        ["-----", "-----", "-----", "-----", "-----", "-----", 0]
    ])

    const [existe, setexiste] = useState(false)
    axios.post('http://localhost:3000/UsuarioLinea')
    .then(response=>{
        if (response.data === "no"){
            setexiste(true)
        }
    }).catch(error=>{
        console.log(error);
    })


    useEffect(()=>{
        let data = localStorage.getItem('prueba1')
        if(data != null){
            setlistado(JSON.parse(data))
        }
    }, [])

    if (existe === true) {return(
        <div className="General">
            <Router>
                <NavBar/>
            </Router>
            <Segment placeholder id="SubirArchivo">
            <Header icon>
                <Icon name='user secret' />
                    DEBE INICIAR SESIÓN
                <br/><br/>
            </Header>
            </Segment>
        </div>
    )
    }else{
        return (
            <div className="Carrito">
            <Router>
                <NavBar/>
            </Router>
            <Label color="green"  size="big" className="Alerta" id="comprado">FIN COMPRAS</Label>
                <Tabla  data = {listado}
                        encabezados={encabezado}
                />
                <div>
                    
                </div>
            </div>
        )
    }
}

export default CarritoCompras
