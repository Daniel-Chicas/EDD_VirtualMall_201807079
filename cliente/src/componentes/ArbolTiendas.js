import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router} from 'react-router-dom'
import {  Header, Icon, Segment } from 'semantic-ui-react'
import ArbolTiendaspdf from '../ArbolesMerkle/ArbolTiendas.pdf'
import '../Inicio.css'

function ArbolTiendas() {
    const [loading, setloading] = useState(false)
    const [existe, setexiste] = useState(false)

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
                    <embed src={ArbolTiendaspdf} type='application/pdf' width="100%" height="625"  /> 
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

export default ArbolTiendas
