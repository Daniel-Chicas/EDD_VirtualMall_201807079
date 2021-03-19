import {React, useEffect, useState} from 'react'
import Tabla from './Tabla'
import '../css/CarritoCompras.css'

function CarritoCompras() {
    const encabezado=['id', 'Tienda', 'Departamento', 'Calificación', 'Nombre Producto', 'Código Producto', 'Disponibilidad en stock', 'Cantidad', 'Confirmar', "Eliminar"]
    const [listado, setlistado] = useState([
        ["-----", "-----", "-----", "-----", "-----", "-----", 0]
    ])
    useEffect(()=>{
        let data = localStorage.getItem('prueba1')
        if(data != null){
            setlistado(JSON.parse(data))
        }
    }, [])

    return (
        <div className="Carrito">
            <Tabla  data = {listado}
                    encabezados={encabezado}
            />
            <div>
                
            </div>
        </div>
    )
}

export default CarritoCompras
