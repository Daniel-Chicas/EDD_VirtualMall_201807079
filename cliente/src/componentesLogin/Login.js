import {React, useState} from 'react'
import {BrowserRouter as Router} from 'react-router-dom'
import NavBarInicio from '../componentesLogin/NavBarIncio'
import '../Inicio.css'
import '../css/CargaArchivos.css'
import { Input, Header, Icon, Segment, Button, Label } from 'semantic-ui-react'
import { useHistory } from "react-router-dom";
const axios=require('axios').default

//Usuario.Insertar(Usuarios.NuevaLlave(1234567890101, "EDD2021", " auxiliar@edd.com", "1234", "Administrador"))

function Login() {
    const Cerrar = "si"

    var CerrarSesion = {
        Cerrar
    }

    axios.post("http://localhost:3000/IniciarSesion", JSON.stringify(CerrarSesion) , {headers:{ 'Content-Type':'multipart/form-data'}})
    .then(response=>{
        console.log(response)
    }).catch(error=>{
        console.log(error);
    })


    const [nombre, setnombre] = useState("")
    const [Password, setPassword] = useState("")
    const history = useHistory();

    const iniciar = () => {
        if (Password === '') {
            document.getElementById("ContraValido").style.visibility = 'visible'
            document.getElementById("ContraValido").innerHTML = "Contraseña Incorrecta"
        }else{
            document.getElementById("ContraValido").style.visibility = 'hidden'
            const Nombre = parseInt(nombre)
            if(isNaN(Nombre) === true){
                document.getElementById("UsuarioValido").style.visibility = 'visible'
                document.getElementById("UsuarioValido").innerHTML = "DPI Incorrecto"
            }else{
                document.getElementById("UsuarioValido").style.visibility = 'hidden'
                var Usuario = {
                    Nombre,
                    Password
                }
                axios.post("http://localhost:3000/IniciarSesion", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                    .then(response=>{
                    if(response.data[0] === "si"){
                        if(response.data[1] === "si") {
                            if (response.data[2] === "Administrador" || response.data[2] === "Admin") {
                                history.push('/CargaArchivos')   
                            }else{
                                history.push('/VerTiendas')   
                            }
                        }else{
                            document.getElementById("ContraValido").style.visibility = 'visible'
                            document.getElementById("ContraValido").innerHTML = "Contraseña Incorrecta"
                        }
                    }else{
                        document.getElementById("UsuarioValido").style.visibility = 'visible'
                        document.getElementById("UsuarioValido").innerHTML = "DPI Incorrecto"
                    }


                    }).catch(error=>{
                        console.log(error);
                    })
            }
        }
    }
    return (
        <div className="GeneralLogin">
      <Router>
        <NavBarInicio/>
      </Router>
            <Segment placeholder id="SubirArchivo">
                <Header icon>
                    <Icon name='user'/>
                        INICIAR SESIÓN
                        <br></br>
                    <Input type="text" name="Usuario" label='DPI' placeholder='1234567890' onChange={e => setnombre(e.target.value)} className="Datos"/>
                    <br/>
                    <Label pointing prompt color="red" id="UsuarioValido" className="Alerta"></Label>
                    <br/>
                    <Input type="password" label='Contraseña' placeholder='' onChange={e => setPassword(e.target.value)} className="Datos"/>
                    <br/>
                    <Label pointing prompt color="red" id="ContraValido" className="Alerta"></Label>
                    <br/>
                    <Button color='green' className="Datos" onClick={iniciar}>INICIAR SESIÓN</Button>
                </Header>
            </Segment>
        </div>
    )
}

export default Login
