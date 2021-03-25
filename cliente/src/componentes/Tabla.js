import React from 'react'
import Fila from '../componentes/Fila'
import { Table } from 'semantic-ui-react'
import { Button, Icon } from 'semantic-ui-react'
import { useHistory } from "react-router-dom";
import '../css/CarritoCompras.css'

function Tabla(props) {
    const history = useHistory();
    const mandar = () =>{
        history.push(`/VerTiendas`)
    }
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
                                tienda={dato.tienda}
                                departamento={dato.departamento}
                                calificacion = {dato.calificacion}
                                nombreProducto ={dato.nombreP}
                                codigoProducto = {dato.producto}
                                precio = {dato.precio}
                                CantidadMax = {dato.CantidadMax}
                                cantidad = {dato.cantidad}
                                todo = {props.data}
                            />
                        ))}
                        <Fila
                                index={props.data.length}
                                tienda={""}
                                departamento={""}
                                calificacion = {""}
                                nombreProducto ={""}
                                codigoProducto = {"Sumatoria"}
                                cantidad = {""}
                                CantidadMax = {""}
                                todo = {props.data}
                        />
                        <Fila
                                index={props.data.length}
                                tienda={""}
                                departamento={""}
                                calificacion = {""}
                                nombreProducto ={""}
                                codigoProducto = {"ACEPTAR TODO"}
                                cantidad = {""}
                                CantidadMax = {""}
                                todo = {props.data}
                        />
                    </tbody>
                </Table>
            </div>
        </div>
        <br/>
        <div>
            <Button  color='red' animated onClick={mandar}>
                <Button.Content visible >SEGUIR COMPRANDO</Button.Content>
                <Button.Content hidden>
                <Icon name='arrow left' />
                </Button.Content>
            </Button>
        <br/>
        </div>
        <br/>
        <br/>
        </>
    )
}

export default Tabla
