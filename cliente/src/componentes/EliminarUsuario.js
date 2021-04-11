import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import { Input, Label } from 'semantic-ui-react'
import '../Inicio.css'
import axios from 'axios'

function EliminarUsuario() {
    
    const [loading, setloading] = useState(false)
    const [nombre, setnombre] = useState("")
    const [Password, setPassword] = useState("")
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
    
    const eliminar = () => {
        if (Password === '') {
            document.getElementById("ContraValido").style.visibility = 'visible'
            document.getElementById("ContraValido").innerHTML = "Contraseña Incorrecta"
        }else{
            document.getElementById("ContraValido").style.visibility = 'hidden'
            const Nombre = parseInt(nombre)
            if(isNaN(Nombre) === true){
                document.getElementById("UsuarioValido").style.visibility = 'visible'
                document.getElementById("UsuarioValido").innerHTML = "Debe ingresar un dpi válido"
            }else{
                document.getElementById("UsuarioValido").style.visibility = 'hidden'
                document.getElementById("Eliminado").style.visibility = 'hidden'
                var Usuario = {
                    Nombre,
                    Password
                }
                axios.post("http://localhost:3000/EliminarUsuario", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                    .then(response=>{
                        console.log(response.data)
                        if(response.data[0] === "si"){
                            if(response.data[1] === "si") {
                                if (response.data[2] === "EL USUARIO HA SIDO ELIMINADO" ){
                                    document.getElementById("Eliminado").style.visibility = 'visible'
                                }
                            }else{
                                document.getElementById("ContraValido").style.visibility = 'visible'
                                document.getElementById("ContraValido").innerHTML = "Contraseña Incorrecta"
                            }
                        }else{
                            document.getElementById("UsuarioValido").style.visibility = 'visible'
                            document.getElementById("UsuarioValido").innerHTML = "Debe ingresar un dpi válido..."
                        }
                    }).catch(error=>{
                        console.log(error);
                    })
            }
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
                <div className="Eliminar">
                <Router>
                    <NavBar/>
                </Router>

                <Segment placeholder id="SubirArchivo">
                    <Header icon>
                        <Icon name='user delete' />ELIMINAR USUARIO :(
                            <br></br>
                        <Label pointing="right" prompt color="red" id="UsuarioValido" className="Alerta"></Label>
                        <Input type="text" name="Usuario" label='DPI' placeholder='1234567890' onChange={e => setnombre(e.target.value)} className="Datos"/>
                        <Label color='green' pointing='left' prompt id="Eliminado"  className="Alerta">Usuario eliminado con éxito!</Label>
                        <br/>
                        <br/>
                        <Label pointing="right" prompt color="red" id="ContraValido" className="Alerta"></Label>
                        <Input type="password" label='Contraseña del Usuario' placeholder='' onChange={e => setPassword(e.target.value)} className="Datos"/>
                        <br/>
                        <br/>
                        <Button color='green' className="Datos" onClick={eliminar}>ELIMINAR</Button>
                    </Header>
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

export default EliminarUsuario
