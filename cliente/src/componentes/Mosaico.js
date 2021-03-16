import React from 'react'
import Carta from './Cartas'


function Mosaico(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    c.Departamentos.map((x, contador)=>(
                        x.Tiendas.map((y, cuenta)=>(
                            <Carta 
                                Departamento={x.Nombre}
                                nombre={y.Nombre}
                                calificacion ={y.Calificacion}
                                descripcion={y.Descripcion}
                                imagen={y.Logo}
                                Contacto={y.Contacto}
                                id={y.PosicionVector}
                                key={y.PosicionVector}
                            />
                        ))
                    ))

                ))}
            </div>
        </div>
    )
}

export default Mosaico
