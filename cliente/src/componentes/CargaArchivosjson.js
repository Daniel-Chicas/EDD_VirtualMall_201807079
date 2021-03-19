import {React, useState} from 'react'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'
import '../css/CargaArchivos.css'
import axios from 'axios'


function CargaArchivosjson() {
    const [archivos, setArchivos]=useState(null);

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
            await axios.post("http://localhost:3000/cargarArchivos", f , {headers:{ 'Content-Type':'multipart/form-data'}})
            .then(response=>{
                alert("El archivo se ha cargado")
                console.log(response.data);
            }).catch(error=>{
                console.log(error);
            })
        }else{
            return(
                alert("No se ha podido cargar el archivo, revise la entrada.")
            )
        }
    }

    return (
    <>
    <div className="General">
        <Segment placeholder id="SubirArchivo">
            <Header icon>
                <Icon name='folder open' />
                La extensión del archivo debe ser tipo .json
                <br/><br/>
                <input type="file" name="Files" multiple onChange={(e)=>subirArchivos(e.target.files)}></input>
            </Header>
            <Button id="json" positive onClick={()=>
                insertarArchivos()
                }>Subir Archivo :D</Button>
        </Segment>
    </div>
    </>
    )
}



export default CargaArchivosjson
