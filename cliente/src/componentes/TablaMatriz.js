import React from 'react'
import Fila from '../componentes/FilaMatriz'
import { Table } from 'semantic-ui-react'
import '../css/CarritoCompras.css'

function Tabla(props) {
    return (
        <>
        <div className="Tabla">
            <div className="ui segment container ">
                <Table celled inverted selectable>
                    <thead>
                        <tr>
                            {props.encabezados.map((dato) => (
                                <th>{dato}</th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {props.data.map((dato, index) => (
                            <Fila
                                index={index}
                                tienda={dato.Tienda}
                                departamento={dato.Departamento}
                                calificacion = {dato.Calificacion}
                                cliente = {dato.Cliente}
                                nombreCliente = {dato.Nombre}
                                correo = {dato.Correo}
                                nombreProducto ={dato.NombreProducto}
                                codigoProducto = {dato.CodigoProducto}
                                cantidad = {dato.Cantidad}
                            />
                        ))}
                    </tbody>
                </Table>
            </div>
        </div>
        </>
    )
}

export default Tabla
