import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import arbolMedio from '../ArbolesUsuarios/arbolMedio.pdf'
import '../Inicio.css'
import axios from 'axios'

function UsuariosPDF() {
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
        if (loading === false) {
            return (
                <div className="General">
                    <embed src={arbolMedio} type='application/pdf' width="100%" height="625"  /> 
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

export default UsuariosPDF
