import React from 'react'
import { Button, Icon } from 'semantic-ui-react'
const axios=require('axios').default

function UltimaFila(props) {
    const todo = () =>{
        const Compras = []
        props.todo.map((c, index)=>{
            var a = document.getElementById("Producto"+index).value
            var habilitado = document.getElementById("Producto"+index).disabled
            if(c.producto!== undefined){
                if(habilitado !== true){
                    const Tienda = c.tienda
                    const Productos = []
                    const Departamento = c.departamento
                    const Calificacion = c.calificacion
                    const Codigo = c.producto
                    const Cantidad = parseInt(a)
                    if(a !== ""){
                        if(a !== "0"){
                            var Producto={
                                Codigo,
                                Cantidad
                            }
                            const fecha = new Date()
                            const Fecha = fecha.getDate()+"-"+(fecha.getMonth()+1)+"-"+fecha.getFullYear()
                            Productos.push(Producto)
                            var viene={
                                Fecha,
                                Tienda,
                                Departamento,
                                Calificacion,
                                Productos
                            }
                            Compras.push(viene)
                            var compras ={
                                Compras
                            }
                            if(index === props.todo.length-1){
                                axios.post("http://localhost:3000/carritoCompras", JSON.stringify(compras) , {headers:{ 'Content-Type':'multipart/form-data'}})
                                .then(response=>{
                                    console.log(response.data);
                                }).catch(error=>{
                                    console.log(error);
                                })
                            }  
                        }
                    }
                }
            }
            document.getElementById("Producto"+index).disabled = true
            localStorage.clear()
        })
    }
    const eliminarTodo = () =>{
        localStorage.clear()
        props.todo.map((c, index)=>{
            document.getElementById("Producto"+index).disabled = true
        }
    )}
    return (
        <>
        <tr key={props.index}>
            <td>{props.index}</td>
                <td>{props.tienda}</td>
                <td>{props.departamento}</td>
                <td>{props.calificacion}</td>
                <td>{props.nombreProducto}</td>
                <td>{props.codigoProducto}</td>
                <td>
                <div>
                </div>
                </td>
                <td>
                <div>
                <Button  color='green' animated onClick={todo}>
                <Button.Content visible >Comprar Todo</Button.Content>
                <Button.Content hidden>
                <Icon name='shop' />
                </Button.Content>
                </Button>
                </div>
                </td>
                <td>
                <div>
                <Button  color='red' animated onClick={eliminarTodo}>
                <Button.Content visible >Eliminar Todo</Button.Content>
                <Button.Content hidden>
                <Icon name='shop' />
                </Button.Content>
                </Button>
                </div>
                </td>
        </tr>
        </>
    )
}

export default UltimaFila
