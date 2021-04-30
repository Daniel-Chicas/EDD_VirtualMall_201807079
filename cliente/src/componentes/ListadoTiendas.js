import {React, useEffect, useState} from 'react'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import Mosaico from './Mosaico'
import '../css/Mosaico.css'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'

const axios=require('axios').default

function ListadoTiendas() {
    const [tiendas, settiendas] = useState([])
    const [loading, setloading] = useState(false)
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
        async function obtener(){
            if(tiendas.length===0){
                const data = await axios.get('http://localhost:3000/guardar')
                if (data.data.Datos !== null) {
                console.log(data.data.Datos)
                settiendas(data.data.Datos)
                setloading(true)
                }
            }
        }
        obtener()
    })

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
        if(loading === false){
            return(
                <>
                <Router>
                    <NavBar/>   
                </Router>
                <div className="ui segment carga" id="Cargando">
                    <div className="ui active dimmer">
                        <div className="ui text loader">Loading</div>
                    </div>
                    <p />
                </div>
                </>
            )
        }else{
            return (
                <div className="ImportList">       
                <Router>
                    <NavBar/>   
                </Router>
                        <br></br>
                        <Mosaico productos={tiendas} />
                </div>
            )
        }
    }
}

export default ListadoTiendas
