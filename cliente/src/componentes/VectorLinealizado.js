import {React, useState} from 'react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import '../pdf.css'
import axios from 'axios'


function UsuariosPDF() {
    const [loading, setloading] = useState(false)
    const [existe, setexiste] = useState(false)
    const [archivos, setArchivos]=useState(null);

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

    const subirArchivos=e=>{
        setArchivos(e);
    }

    const insertarArchivos=async()=>{
        localStorage.clear()
        const f = new FormData();
        if (archivos != null) {
            for (let index = 0; index<archivos.length; index++){
                f.append("Indice",archivos[index])
            }
            let pdffFileURL = URL.createObjectURL(archivos[0]);
            document.querySelector('#VistaPrevia').setAttribute(('src'), pdffFileURL);
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
                <div className="GeneralLogin">
                    <Segment placeholder id="SubirArchivo">
                    <Header icon>
                        <Icon name='pdf file' />
                        <br/><br/>
                        <input type="file" name="Files" multiple onChange={(e)=>subirArchivos(e.target.files)}></input>
                    </Header>
                    <Button id="json" positive onClick={()=>insertarArchivos()}>Ver Archivo :D</Button>
                </Segment>
                <embed type='application/pdf' width="100%" height="400" id="VistaPrevia"/> 
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
