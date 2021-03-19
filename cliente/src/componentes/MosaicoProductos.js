import React from 'react'
import Carta from './CartasProductos'


function MosaicoProductos(props) {
    const numero = 1
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.Productos.map((c, index) => (
                    <Carta 
                        tienda={props.productos.Tienda}
                        departamento={props.productos.Departamento}
                        calificacion={props.productos.Calificacion}
                        imagen={c.Imagen}
                        nombre={c.Nombre}
                        codigo ={c.Codigo}
                        CantidadMax = {c.Cantidad}
                        cantidad = {numero}
                        descripcion={c.Descripcion}
                        precio={c.Precio}
                        id={c.Codigo}
                        key={c.Codigo}
                    />

                ))}
            </div>
        </div>
    )
}

export default MosaicoProductos
