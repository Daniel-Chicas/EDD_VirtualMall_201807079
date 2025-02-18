import {React,  useState} from 'react'
import NavBar from '../componentes/NavBar'
import { Header, Icon, Segment } from 'semantic-ui-react'
import {BrowserRouter as Router} from 'react-router-dom'
import '../css/Pedidos.css'
import Tree from './Tree'
const axios=require('axios').default

function Pedidos() {
    
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

    const [anios, setanios] = useState([])
    async function obtener(){
        if(anios.length===0){
        const data = await axios.get('http://localhost:3000/DatosMatriz')
        if (data.data.General !== null) {
            console.log(data.data.General)
            setanios(data.data.General)
            }
        }
    }
    obtener()

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
                <div className="Pedidos">
                    <Router>
                        <NavBar/>
                    </Router>
                    <Tree 
                        listaAnios={anios}
                    />
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

export default Pedidos
