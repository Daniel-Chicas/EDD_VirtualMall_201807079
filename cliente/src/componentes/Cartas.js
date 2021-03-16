import React from 'react'
import ProductosTienda from './ProductosTienda'
import '../css/Carta.css'

function Cartas(props) {
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
                </div>
                <div className="extra content">
                    <span><i className="tty" />Contacto: {props.Contacto}</span>
                </div>
            </div>
        </div>
    )
}

export default Cartas
