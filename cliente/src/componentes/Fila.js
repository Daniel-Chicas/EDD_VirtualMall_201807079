import React from 'react'
import { Input } from 'semantic-ui-react'
import UltimaFila from './UltimaFila'
import UltimaFilaTotal from './UltimaFilaTotal'

function Fila(props) {
    if(props.codigoProducto === "ACEPTAR TODO"){
       return(
           <UltimaFila
                todo={props.todo}
           />
       )
    }else if (props.codigoProducto === "Sumatoria"){
        return(
            <UltimaFilaTotal
                 todo={props.todo}
            />
        )
    }

    function suma(){
        var cantidad = parseInt(document.getElementById("Producto"+props.index).value)
        var precio = props.precio
        document.getElementById("SumaParcial"+props.index).innerHTML = "Q"+(cantidad*precio)
        var total = 0
        props.todo.map((x, index)=>{
            var cantidad = parseInt(document.getElementById("Producto"+index).value)
                var precio = x.precio
                total += precio*cantidad
        })
        document.getElementById("SumaTotal").innerHTML = "Q"+(total)
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
                <td>Q{props.precio}</td>
                <td>{props.CantidadMax}</td>
                <td><Input type="number" id={"Producto"+props.index} name="tentacles" min="0" max={props.CantidadMax} onChange={suma}/></td>
                <td><label id={"SumaParcial"+props.index}></label></td>
        </tr>
        </>
    )
}

export default Fila
