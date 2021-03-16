import React from 'react'
import Carta from './CartasProductos'


function MosaicoProductos(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    <Carta 
                        imagen={c.Imagen}
                        nombre={c.Nombre}
                        codigo ={c.Codigo}
                        cantidad = {c.Cantidad}
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
