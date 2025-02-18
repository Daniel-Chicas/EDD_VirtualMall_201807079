import React from 'react'
import '../css/Carta.css'

function Cartas(props) {
    const tienda = props.tienda
    const departamento = props.departamento
    const calificacion = props.calificacion
    const nombreP = props.nombre
    const producto = props.id
    const CantidadMax = props.CantidadMax
    const cantidad = props.cantidad
    const precio = props.precio
    const almacenamiento = props.almacenamiento

    const enviar = ()=>{
        var json={
            tienda,
            departamento,
            calificacion,
            nombreP,
            cantidad,
            CantidadMax,
            producto,
            precio,
            almacenamiento
        }
        var datos = localStorage.getItem('prueba1')
        if (datos === null || datos === undefined) {
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
                        <br></br>
                        <a>Lugar almacenado: {props.almacenamiento}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={enviar}>Añadir al carrito</div>
                    <div className="ui basic red button center fluid" onClick={()=>{window.location="http://localhost:8001/ComentariosP/"+props.departamento+"&"+props.tienda+"&"+props.calificacion+"&"+props.id}}>VER COMENTARIOS</div>
                </div>
                <div className="extra content">
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default Cartas
