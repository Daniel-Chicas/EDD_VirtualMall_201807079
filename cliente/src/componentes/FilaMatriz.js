import React from 'react'
import UltimaFila from './UltimaFila'

function Fila(props) {
    if(props.codigoProducto === "ACEPTAR TODO"){
       return(
           <UltimaFila
                todo={props.todo}
           />
       )
    }
    return (
        <>
        <tr key={props.index}>
            <td>{props.index}</td>
                <td>{props.tienda}</td>
                <td>{props.departamento}</td>
                <td>{props.calificacion}</td>
                <td>{props.cliente}</td>
                <td>{props.nombreCliente}</td>
                <td>{props.correo}</td>
                <td>{props.nombreProducto}</td>
                <td>{props.codigoProducto}</td>
                <td>{props.cantidad}</td>
        </tr>
        </>
    )
}

export default Fila
