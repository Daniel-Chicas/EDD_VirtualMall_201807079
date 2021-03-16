import React from 'react'
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
                        <a>Código: {props.codigo}</a>
                        <br></br>
                        <a>Cantidad Disponible: {props.cantidad}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={()=>{console.log(props.id)}}>Comprar</div>
                </div>
                <div className="extra content">
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default Cartas
