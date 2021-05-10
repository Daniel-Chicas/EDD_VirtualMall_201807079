import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import '../css/Nav.css'
var colores = ['orange', 'yellow','green', "brown", 'purple', 'red','black', 'pink', 'white', 'blue']
var opciones = ['Ver Tiendas', 'Carrito de Compras', 'Ver pedidos','Cargar Archivos', 'Eliminar Usuario', 'Arboles Usuario','Cambiar Clave', 'Ver Vector', 'Árboles de Merkle','Cerrar Sesión', '/']

function NavBar() {
    const [activo, setactivo] = useState(colores[10])
    if (activo === "blue") {
        window.location.href = "http://localhost:8001/Login"
    }else if(activo === "brown"){
        window.location.href = "http://localhost:8001/CargaArchivos"
    }else if(activo === "green"){
        window.location.href = "http://localhost:8001/VerPedidos"
    }else if(activo === "yellow"){
        window.location.href = "http://localhost:8001/CarritoCompras"
    }else if(activo === "orange"){
        window.location.href = "http://localhost:8001/VerTiendas"
    }else if(activo === "purple"){
        window.location.href = "http://localhost:8001/EliminarUsuario"
    }else if(activo === "red"){
        window.location.href = "http://localhost:8001/ArbolesUsuario"
    }else if(activo === "black"){
        window.location.href = "http://localhost:8001/CambiarClave"
    }else if(activo === "pink"){
        window.location.href = "http://localhost:8001/Vector"
    }else if(activo === "white"){
        window.location.href = "http://localhost:8001/ArbolesMerkle"
    }
    return (
       <Menu inverted className="Nav">
           {colores.map((c,iterador)=>(
               <Menu.Item 
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
