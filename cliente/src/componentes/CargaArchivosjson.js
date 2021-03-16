import {React, useState} from 'react'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import '../css/CargaArchivos.css'
import { Message } from 'semantic-ui-react'
import axios from 'axios'


function CargaArchivosjson() {
    const [archivos, setArchivos]=useState(null);

    const subirArchivos=e=>{
        setArchivos(e);
    }

    const insertarArchivos=async()=>{
        const f = new FormData();
        if (archivos != null) {
            for (let index = 0; index<archivos.length; index++){
                f.append("Indice",archivos[index])
            }
            await axios.post("http://localhost:3000/cargarArchivos", f , {headers:{ 'Content-Type':'multipart/form-data'}})
            .then(response=>{
                alert("EL ARCHIVO HA SIDO GUARDADO EXISTOSAMENTE!.")
                console.log(response.data);
            }).catch(error=>{
                console.log(error);
            })
        }else{
            alert("Debe ingresar un archivo")
        }
    }

    return (
    <>
    <div className="General">
        <Segment placeholder id="SubirArchivo">
            <Header icon>
                <Icon name='folder open' />
                La extensión del archivo debe ser tipo .json
            </Header>
            <div>
                <input type="file" name="Files" multiple onChange={(e)=>subirArchivos(e.target.files)}></input>
            </div>
            <br></br>
            <Button id="json" positive onClick={()=>insertarArchivos()}>Subir Archivo :D</Button>
        </Segment>
    </div>
    </>
    )
}



export default CargaArchivosjson
