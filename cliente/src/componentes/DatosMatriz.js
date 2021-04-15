import {React,  useState} from 'react'
import Calendar from 'react-calendar';
import Tabla from './TablaMatriz'
import { Button, Segment } from 'semantic-ui-react'
import '../css/Imagen.css'
const axios=require('axios').default

function DatosMatriz() {
    var URLactual = window.location;
    var direccion = URLactual.toString().split("/")
    var direccion2 = direccion[4].split("&")
    var cadena = direccion2[1]+"-1-"+direccion2[0]
    var fecha = Date.parse(cadena)
    var fechaG = new Date(fecha)
    const [value, onChange] = useState(fechaG);
    const [tiendas, settiendas] = useState([])
    const encabezado=['id', 'Tienda', 'Departamento', 'Calificación', 'DPI', 'Nombre Cliente', 'Correo', 'Nombre Producto', 'Código Producto', 'Cantidad']
    
    const ver = ()=>{
        async function obtener(){
            const dia = value.getDate().toString()
            const mes = value.getMonth()+1
            const anio = value.getFullYear()
            const data = await axios.get('http://localhost:3000/Pedido/'+anio+"&"+mes+"&"+dia)
            if(data.data.Datos!=null){
                settiendas(data.data.Datos)
            }
        }
        obtener()
    }
    return (
        <div>
            <Segment placeholder className="FondoImagen">
                    <Calendar
                        onChange={onChange}
                        value={value}
                    />
                    <br/>
                    <br/>
                    <Button inverted color='red' onClick={ver}>Ver Pedidos del {value.getDate().toString()} del {value.getMonth()+1}</Button>
                    <br/>
                    <br/>
                    <a href="/Recorrido"  target="_blank" >
                        <Button inverted color='black' onClick={ver}>Ver Recorrido</Button>
                    </a>
            </Segment>
            <br/>
            <Tabla  data = {tiendas}
                    encabezados={encabezado}
            />
        </div>
    )
}

export default DatosMatriz
