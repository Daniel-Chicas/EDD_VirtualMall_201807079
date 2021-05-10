import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import '../css/Nav.css'

const colores=['light blue','brown']
const opciones =['Iniciar Sesion', 'Crear Usuario']

function NavBarIncio() {
    const [activo, setactivo] = useState(colores[10])
    if (activo === "brown") {
        window.location.href = "http://localhost:8001/CrearUsuario"
    }else if (activo === "light blue"){
        window.location.href = "http://localhost:8001/Login"
    }
    return (
       <Menu inverted className="Nav" >
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

export default NavBarIncio
