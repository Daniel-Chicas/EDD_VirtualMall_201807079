import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import '../css/Nav.css'

const colores=['orange','yellow','green','purple']
const opciones =['Cargar Archivos', 'Listado de Tiendas', 'Carrito de Compras', 'Ver pedidos']
const url =['/CargaArchivos', '/VerTiendas', '/CarritoCompras', '/VerPedidos', '/']

function NavBar() {
    const [activo, setactivo] = useState(colores[10])
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
