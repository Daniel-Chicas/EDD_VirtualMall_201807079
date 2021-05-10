import {React, useState} from 'react'
import { Segment,  } from 'semantic-ui-react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, } from 'react-router-dom'
import {  Header, Icon } from 'semantic-ui-react'
import arbol from '../ImagenArbol/ArbolProductos.pdf'
import '../css/Imagen.css'
import axios from 'axios'

function VerArbol() {
    const [loading, setloading] = useState(false)
    const [existe, setexiste] = useState(false)


    axios.post('http://localhost:3000/UsuarioLinea')
    .then(response=>{
        if (response.data === "Usuario") {
            setloading(true)
        }else if (response.data === "no"){
            setexiste(true)
        }
    }).catch(error=>{
        console.log(error);
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
            return (
                <div className="Arbol">
                    <Router>
                        <NavBar/>
                    </Router>
                    <embed src={arbol} type='application/pdf' width="100%" height="600"/>
                </div>
            )
        }else{
            return(
                <div className="General">
                    <Router>
                        <NavBar/>
                    </Router>
                    <Segment placeholder id="SubirArchivo">
                    <Header icon>
                        <Icon name='user secret' />
                            DEBE TENER PERMISOS DE ADMINISTRADOR PARA INGRESAR A ESTA PÁGINA
                        <br/><br/>
                    </Header>
                    </Segment>
                </div>
            )
        }
    }
}

export default VerArbol