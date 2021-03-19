import React, { useState } from 'react'
import '../css/Carta.css'

function Cartas(props) {
    console.log(props)
    const tienda = props.tienda
    const departamento = props.departamento
    const calificacion = props.calificacion
    const nombreP = props.nombre
    const producto = props.id
    const CantidadMax = props.CantidadMax
    const cantidad = props.cantidad

    const enviar = ()=>{
        var json={
            tienda,
            departamento,
            calificacion,
            nombreP,
            cantidad,
            CantidadMax,
            producto
        }
        var datos = localStorage.getItem('prueba1')
        if (datos == null || datos == undefined) {
            console.log(datos)
            localStorage.setItem('prueba1', JSON.stringify([json]))
        }else{
            datos=JSON.parse(datos)
            datos.push(json)
            console.log(datos)
            localStorage.setItem('prueba1', JSON.stringify(datos))
        }
        console.log(datos)
    }


    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen} />
                </div>
                <div className="content">
                    <div className="header" onChange={e => (e.target.value)}>{props.nombre}</div>
                    <div className="meta">
                        <a>Código: {props.codigo}</a>
                        <br></br>
                        <a>Cantidad Disponible: {props.CantidadMax}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={enviar}>Añadir al carrito</div>
                </div>
                <div className="extra content">
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default Cartas
