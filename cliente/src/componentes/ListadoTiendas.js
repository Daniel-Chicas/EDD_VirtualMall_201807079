import {React, useEffect, useState} from 'react'
import Mosaico from './Mosaico'
import '../css/Mosaico.css'
const axios=require('axios').default

function ListadoTiendas() {
    const [tiendas, settiendas] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(()=>{
        async function obtener(){
            if(tiendas.length===0){
                const data = await axios.get('http://localhost:3000/guardar')
                if (data.data.Datos !== null) {
                console.log(data.data.Datos)
                settiendas(data.data.Datos)
                setloading(true)
                }
            }
        }
        obtener()
    }) 

    if(loading === false){
        return(
            <div className="ui segment carga">
                <div className="ui active dimmer">
                    <div className="ui text loader">Loading</div>
                </div>
                <p />
            </div>
        )
    }else{
        return (
            <div className="ImportList">
                    <br></br>
                    <Mosaico productos={tiendas} />
            </div>
        )
    }
}

export default ListadoTiendas
