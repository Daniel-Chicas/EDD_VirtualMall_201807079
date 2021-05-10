import {React, useEffect, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router} from 'react-router-dom'
import { Button, Header, Icon, Segment, Label} from 'semantic-ui-react'
import '../Inicio.css'
import axios from 'axios'

function ArbolesMerkle() {
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

    useEffect(()=>{
        async function obtener(){
            const user = await axios.get('http://localhost:3000/HacerArboles')
            console.log(user.data)
            if (user.data.Alerta === "NO SE HA CREADO NINGUNA COPIA") {
                alert("AÚN NO SE HA CREADO NINGUNA COPIA")
            }
        }
        obtener()
    })
    
    const ArbolT = ()=>{
        window.location.href = "http://localhost:8001/ArbolTiendas"
    }
    const ArbolP = ()=>{
        window.location.href = "http://localhost:8001/ArbolProductos"
    }
    const ArbolPP = ()=>{
        window.location.href = "http://localhost:8001/ArbolPedidos"
    }
    const ArbolU = ()=>{
        window.location.href = "http://localhost:8001/ArbolUsuarios"
    }
    const ArbolCT = ()=>{
        window.location.href = "http://localhost:8001/ArbolComentariosT"
    }
    const ArbolCP = ()=>{
        window.location.href = "http://localhost:8001/ArbolComentariosP"
    }
    const Arreglar = () =>{
        async function obtener(){
            const user = await axios.get('http://localhost:3000/VerificarArboles')
            alert("Arreglado")
        }
        obtener()
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
                       <Icon name='sitemap' />  
                <br/>
                       <Button  color='GRAY' onClick={Arreglar}>
                            ARREGLAR LOS ARCHIVOS QUE ESTÉN DAÑADOS 
                        </Button>
                </Header>
                <br/>
                <br/>
                <br/>
                <br/>
                
                <a target="_blank" onClick={ArbolT}>
                    <Label inverted color='red' pointing='right'>Arbol de Tiendas</Label>
                    <Button inverted color='red' >VER ÁRBOL</Button>
                    </a>
                <br/>
                <br/>
                <a target="_blank" onClick={ArbolP}>
                    <Button inverted color='green' >VER ÁRBOL</Button>
                    <Label inverted color='green' pointing='left'>Arbol de Productos</Label>
                    </a>
                <br/>
                <br/>
                <a  target="_blank" onClick={ArbolPP}>
                    <Label inverted color='orange' pointing='right'>Arbol de Pedidos</Label>
                    <Button inverted color='orange' >VER ÁRBOL</Button>
                   </a>
                <br/>
                <br/>
                <a target="_blank" onClick={ArbolU}>
                    <Button inverted color='purple' >VER ÁRBOL</Button>
                    <Label inverted color='purple' pointing='left'>Arbol de Usuarios</Label>
                    </a>
                <br/>
                <br/>
                
                <a  target="_blank" onClick={ArbolCT}>
                    <Label inverted color='brown' pointing='right'>Arbol de Comentarios de las Tiendas</Label>
                    <Button inverted color='brown' >VER ÁRBOL</Button>
                    </a>
                <br/>
                <br/>
                <a  target="_blank" onClick={ArbolCP}>
                    <Button inverted color='blue' >VER ÁRBOL</Button>
                    <Label inverted color='blue' pointing='left'>Arbol de Comentarios de los Productos</Label>
                    </a>
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

export default ArbolesMerkle
