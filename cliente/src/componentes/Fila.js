import React from 'react'
import { Input } from 'semantic-ui-react'
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
                <td>{props.nombreProducto}</td>
                <td>{props.codigoProducto}</td>
                <td>{props.CantidadMax}</td>
                <td><Input type="number" id={"Producto"+props.index} name="tentacles" min="0" max={props.CantidadMax}/></td>
                <td>
                </td>
                <td>
                </td>
        </tr>
        </>
    )
}

export default Fila
