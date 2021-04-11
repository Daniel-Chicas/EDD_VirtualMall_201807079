import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import { Button, Header, Icon, Segment, Label} from 'semantic-ui-react'
import '../Inicio.css'

import axios from 'axios'

function ArbolesUsuario() {
    
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

    
    const crear = async() =>{
        const data = await axios.get('http://localhost:3000/ArbolesB')
        console.log(data.data)
    }

    const crearVector = async() =>{
        const data = await axios.get('http://localhost:3000/ArbolesB')
        console.log(data.data)
    }

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
                <div className="Eliminar" style={{textAlign: 'center'}}>
                <Router>
                    <NavBar/>
                </Router>

                <Segment inverted className="Fondo">
                <a href="./UsuariosPDF" target="_blank" onClick={crear}>
                    <Label inverted color='red' pointing='right'>ÁRBOL DE CUENTAS (Sin Cifrar)</Label>
                    <Button inverted color='red' onClick={crear}>VER ÁRBOL</Button></a>
                <br/>
                <br/>
                <a href="./UsuariosT" target="_blank" onClick={crear}>
                    <Button inverted color='green' onClick={crear}>VER ÁRBOL</Button>
                    <Label inverted color='green' pointing='left'>ÁRBOL DE CUENTAS (Cifrado)</Label>
                </a>
                <br/>
                <br/>
                <a href="./UsuariosM" target="_blank" onClick={crear}>
                    <Label inverted color='orange' pointing='right'>ÁRBOL DE CUENTAS (Cifrado Sensible)</Label>
                    <Button inverted color='orange' onClick={crear}>VER ÁRBOL</Button></a>
                <br/>
                <br/>
                <a href="./Vector" target="_blank" onClick={crearVector}>
                    <Button inverted color='blue' onClick={crearVector}>VER VECTOR</Button>
                    <Label inverted color='blue' pointing='left'>ESTRUCTURA LINEALIZADA</Label>
                </a>
                </Segment>
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

export default ArbolesUsuario
