import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router} from 'react-router-dom'
import { Button, Header, Icon, Segment, Label, Input} from 'semantic-ui-react'
import '../Inicio.css'

import axios from 'axios'
import UsuariosPDF from './UsuariosPDF'

function ArbolesUsuario() {
    
    const [loading, setloading] = useState(false)
    const [existe, setexiste] = useState(false)
    const [clave, setClave] = useState("")

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

    
    const UsuarioT = ()=>{

        const LlaveA = ""
            const LlaveN = clave.toString()
            var Usuario = {
                LlaveA,
                LlaveN
            }

        if (clave == "") {
            document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
            document.getElementById("Antigua").style.visibility = 'visible'
        }else{
            document.getElementById("Antigua").style.visibility = 'hidden'
            axios.post("http://localhost:3000/ArbolesB", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    console.log(response.data.Alerta)
                if (response.data.Alerta === "No") {
                    document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
                    document.getElementById("Antigua").style.visibility = 'visible'
                }else{
                    window.location.href = "http://localhost:8001/UsuariosT"
                }
                }).catch(error=>{
                    console.log(error);
                })
        }
    }

    const UsuariosPDF = ()=>{

        const LlaveA = ""
            const LlaveN = clave.toString()
            var Usuario = {
                LlaveA,
                LlaveN
            }

        if (clave == "") {
            document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
            document.getElementById("Antigua").style.visibility = 'visible'
        }else{
            document.getElementById("Antigua").style.visibility = 'hidden'
            axios.post("http://localhost:3000/ArbolesB", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    console.log(response.data.Alerta)
                if (response.data.Alerta === "No") {
                    document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
                    document.getElementById("Antigua").style.visibility = 'visible'
                }else{
                    window.location.href = "http://localhost:8001/UsuariosPDF"
                }
                }).catch(error=>{
                    console.log(error);
                })
        }
    }

    const UsuariosM = ()=>{

        const LlaveA = ""
            const LlaveN = clave.toString()
            var Usuario = {
                LlaveA,
                LlaveN
            }

        if (clave == "") {
            document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
            document.getElementById("Antigua").style.visibility = 'visible'
        }else{
            document.getElementById("Antigua").style.visibility = 'hidden'
            axios.post("http://localhost:3000/ArbolesB", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    console.log(response.data.Alerta)
                if (response.data.Alerta === "No") {
                    document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
                    document.getElementById("Antigua").style.visibility = 'visible'
                }else{
                    window.location.href = "http://localhost:8001/UsuariosM"
                }
                }).catch(error=>{
                    console.log(error);
                })
        }
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
                
                <Header icon>
                       <Icon name='angle double down' />
                </Header>
                        <br/>
                        <Input type="password" name="Usuario" label='Clave' placeholder='' className="Datos" onChange={e => setClave(e.target.value)}/>
                        <br/>
                        <Label pointing prompt color="red" id="Antigua" className="Alerta">Debe ingresar una cotraseña</Label>
                        <br/>
                        <br/>
                        <br/>

                
                <a target="_blank" onClick={UsuariosPDF}>
                    <Label inverted color='red' pointing='right'>ÁRBOL DE CUENTAS (Sin Cifrar)</Label>
                    <Button inverted color='red' onClick={UsuariosPDF}>VER ÁRBOL</Button></a>
                <br/>
                <br/>
                <a target="_blank" onClick={UsuarioT}>
                    <Button inverted color='green' onClick={UsuarioT}>VER ÁRBOL</Button>
                    <Label inverted color='green' pointing='left'>ÁRBOL DE CUENTAS (Cifrado)</Label>
                </a>
                <br/>
                <br/>
                <a  target="_blank" onClick={UsuariosM}>
                    <Label inverted color='orange' pointing='right'>ÁRBOL DE CUENTAS (Cifrado Sensible)</Label>
                    <Button inverted color='orange' onClick={UsuariosM}>VER ÁRBOL</Button></a>
                <br/>
                <br/>
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
