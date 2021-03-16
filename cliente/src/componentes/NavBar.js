import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import '../css/Nav.css'

const colores=['orange','yellow','green','purple','red','brown']
const opciones =['Cargar Archivos', 'Listado de Tiendas','Busqueda Tiendas Específica','Busqueda por Posición' , 'Eliminar Tienda', 'Carrito de Compras']
const url =['/CargaArchivos', '/VerTiendas', '/TiendaEspecifica', 'id/{numero}', '/EliminarTienda', '/CarritoCompras', '/']

function NavBar() {
    const [activo, setactivo] = useState(colores[0])
    return (
       <Menu inverted className="Nav">
           {colores.map((c,iterador)=>(
               <Menu.Item as={Link} to={url[iterador]}
                    key={c}
                    name={opciones[iterador]}
                    active={activo===c}
                    color={c}
                    onClick={()=>setactivo(c)}
               />
           ))}
       </Menu>
    )
}

export default NavBar
