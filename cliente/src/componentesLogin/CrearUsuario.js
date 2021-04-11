import {React, useState} from 'react'
import {BrowserRouter as Router} from 'react-router-dom'
import NavBarInicio from '../componentesLogin/NavBarIncio'
import { Input, Header, Icon, Segment, Button, Label } from 'semantic-ui-react'
import { useHistory } from "react-router-dom";
import '../Inicio.css'
const axios=require('axios').default


function CrearUsuario() {
    const [dpi, setusuario] = useState("")
    const [Password, setcontra] = useState("")
    const [Correo, setcorreo] = useState("")
    const [Nombre, setnombre] = useState("")
    const history = useHistory();

    const iniciar = () => {
        const Usuarios = []
        var expReg= /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/;
        var esValio = expReg.test(Correo)
        if(esValio === false){
            document.getElementById("CorreoValido").style.visibility = 'visible'
            document.getElementById("CorreoValido").innerHTML = "INGRESE UN CORREO VÁLIDO"
        }else{
            document.getElementById("CorreoValido").style.visibility = 'hidden'
            var Cuenta = "Usuario"
            var aux = parseInt(dpi)
            const DPI = aux
            if(isNaN(DPI) === true){
                document.getElementById("UsuarioValido").style.visibility = 'visible'
                document.getElementById("UsuarioValido").innerHTML = "INGRESE UN DPI VÁLIDO"
            }else{
                document.getElementById("UsuarioValido").style.visibility = 'hidden'
                var Usuario = {
                    DPI,
                    Nombre,
                    Correo,
                    Password,
                    Cuenta
                }
                Usuarios.push(Usuario)
                var General = {
                    Usuarios
                }
                axios.post("http://localhost:3000/Usuarios", JSON.stringify(General) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    if (response.data.Alerta === undefined) {
                        document.getElementById("UsuarioValido").style.visibility = 'visible'
                        document.getElementById("UsuarioValido").innerHTML = "ESTE  DPI  YA  ESTÁ  EN  USO"
                    }else{
                        history.push(`/Login`)
                    }
                }).catch(error=>{
                    console.log(error);
                })
            }
        }
    }
    return (
        <div className="General">
            <Router>
                <NavBarInicio/>
            </Router>
            <Segment placeholder id="CrearUsuario" className="General">
                <Header icon>
                    <Icon name='add user'/>
                    CREAR USUARIO
                    <br/>
                    <Input type="text" name="Usuario" label='DPI' placeholder='123456789' onChange={e => setusuario(e.target.value)} className="Datos"/>
                    <br/>
                    <Label pointing prompt color="red" id="UsuarioValido" className="Alerta"></Label>
                    <br/>
                    <Input type="text" label='Correo' placeholder='12345@gmail.com' onChange={e => setcorreo(e.target.value)} className="DatosCrear"/>
                    <br/>
                    <Label pointing prompt color="red" id="CorreoValido" className="Alerta"></Label>
                    <br/>
                    <Input type="password" label='Password' placeholder='contraseña123' onChange={e => setcontra(e.target.value)} className="DatosCrear"/>
                    <br/>
                    <br/>
                    <Input type="text" label='Nombre' placeholder='Jesse Pinkman' onChange={e => setnombre(e.target.value)} className="DatosCrear"/>
                    <br/>
                    <br/>
                    <Button color='green' className="DatosCrear" onClick={iniciar}>Crear Usuario</Button>
                </Header>
                
            <div className="burbujas">
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
                <div className="burbuja"></div>
            </div>
            </Segment>
        </div>
    )
}

export default CrearUsuario
