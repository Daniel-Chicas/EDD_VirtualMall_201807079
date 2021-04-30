import React from 'react'
import '../css/Carta.css'
import Comentarios from './Comentarios'
const axios=require('axios').default

function Cartas(props) {
    const ver = async()=>{
        var cadena = props.Departamento+"&"+props.nombre+"&"+props.calificacion
        const data = await  axios.get('http://localhost:3000/Arbol/'+cadena)
        window.location="http://localhost:8001/VerArbol"
    }

    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen} />
                </div>
                <div className="content">
                    <div className="header">{props.nombre}</div>
                    <div className="meta">
                        <a>{props.Departamento}</a>
                        <a>{props.calificacion}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={()=>{ 
                        window.location="http://localhost:8001/Productos/"+props.Departamento+"&"+props.nombre+"&"+props.calificacion; 
                        console.log(props.id)}}>
                        Ver Productos
                    </div>
                    <div className="ui basic blue button center fluid" onClick={ver}>
                        Ver árbol de la tienda
                    </div>
                    <div className="ui basic green button center fluid" onClick={()=>{ 
                        window.location="http://localhost:8001/Comentarios/"+props.Departamento+"&"+props.nombre+"&"+props.calificacion}}>
                        Ver Comentarios
                    </div>
                </div>
                <div className="extra content">
                    <span><i className="tty" />Contacto: {props.Contacto}</span>
                </div>
            </div>
        </div>
    )
}

export default Cartas
