import {React, useState} from 'react'
import { Input, Header, Icon, Segment, Button, Label } from 'semantic-ui-react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router} from 'react-router-dom'
import '../css/CargaArchivos.css'
import '../Inicio.css'
import axios from 'axios'
var crypto = require('crypto');

function CambiarClave() {

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

    
    const [ant, setant] = useState("")
    const [nueva, setnueva] = useState("")
    const [tiempo, settiempo] = useState(0)

    const insertarArchivos=async()=>{
        if (ant === "") {
            document.getElementById("Antigua").style.visibility = 'visible'
            document.getElementById("Nueva").style.visibility = 'hidden'
        }else if (nueva === ""){
            document.getElementById("Nueva").style.visibility = 'visible'
            document.getElementById("Antigua").style.visibility = 'hidden'
        }else if (nueva !== "" &&  ant !== "") {
            document.getElementById("Nueva").style.visibility = 'hidden'
            document.getElementById("Antigua").style.visibility = 'hidden'
            const LlaveA = crypto.createHash('sha256').update(ant.toString()).digest('hex')
            const LlaveN = crypto.createHash('sha256').update(nueva.toString()).digest('hex')
            const Tiempo = "0"           
            var Usuario = {
                LlaveA,
                LlaveN,
                Tiempo
            }
            console.log(JSON.stringify(Usuario))
            axios.post("http://localhost:3000/CambiarContra", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    console.log(response.data.Alerta)
                if (response.data.Alerta === "No") {
                    document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
                    document.getElementById("Antigua").style.visibility = 'visible'
                }else{
                    document.getElementById("Acepta").style.visibility = 'visible'
                }
                }).catch(error=>{
                    console.log(error);
                })
        }
    }

    const cambiarTiempo=async()=>{
        document.getElementById("Nueva").style.visibility = 'hidden'
        document.getElementById("Antigua").style.visibility = 'hidden'
            const LlaveA = ""
            const LlaveN = ""
            const Tiempo = tiempo           
            var Usuario = {
                LlaveA,
                LlaveN,
                Tiempo
            }
            console.log(JSON.stringify(Usuario))

            axios.post("http://localhost:3000/CambiarContra", JSON.stringify(Usuario) , {headers:{ 'Content-Type':'multipart/form-data'}})
                .then(response=>{
                    console.log(response.data.Alerta)
                if (response.data.Alerta === "No") {
                    document.getElementById("Antigua").innerHTML = 'Contraseña incorrecta'
                    document.getElementById("Antigua").style.visibility = 'visible'
                }else{
                    document.getElementById("Acepta").style.visibility = 'visible'
                }
                }).catch(error=>{
                    console.log(error);
                })
        
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
        if(loading === false){
            return (
            <>
            <div className="GeneralLogin">
                <Router>
                    <NavBar/>
                </Router>
                <div>
                    <Segment placeholder id="SubirArchivo">
                        <Header icon>
                            <Icon name='exchange' />
                            CAMBIAR CLAVE DE ACCESO PARA VER ÁRBOL B
                            <br/><br/>
                            <Input type="password" name="Usuario" label='Clave Antigua' placeholder='' className="Datos" onChange={e => setant(e.target.value)}/>
                            <br/>
                            <Label pointing prompt color="red" id="Antigua" className="Alerta">Debe ingresar una cotraseña</Label>
                            <br/>
                            <Input type="password" label='Nueva Clave' placeholder='' className="Datos" onChange={e => setnueva(e.target.value)}/>
                            <br/>
                            <Label pointing prompt color="purple" id="Acepta" className="Alerta">Contraseña Cambiada</Label>
                            <Label pointing prompt color="red" id="Nueva" className="Alerta">Debe ingresar una cotraseña</Label>
                            <br/>
                            <Button color='green' className="Datos" onClick={insertarArchivos}>Cambiar Clave</Button>
                        </Header>
                    </Segment> 
                </div>
                <div>
                <Segment placeholder id="SubirArchivo">
                    <Header icon>
                        <Icon name='exchange' />
                        CAMBIAR EL TIEMPO DE DURACIÓN DE AUTOGUARDADO
                        <br/><br/>
                        <Input name="Tiempo" label='Tiempo (mins)' placeholder='' className="Datos" onChange={e => settiempo(e.target.value)}/>
                        <br/>
                        <Button color='green' className="Datos" onClick={cambiarTiempo}>Cambiar Tiempo</Button>
                    </Header>
                </Segment> 
                </div>
            </div>
            </>
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



export default CambiarClave
