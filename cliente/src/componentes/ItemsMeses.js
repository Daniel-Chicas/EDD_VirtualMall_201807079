import {React} from 'react'
import { Button } from 'semantic-ui-react'
const axios=require('axios').default

function ItemsMeses(props) {    
    const matriz = async()=>{
            var cadena = props.anio+"&"+props.mes
            const data = await  axios.get('http://localhost:3000/ImagenMatriz/'+cadena)
            console.log(data.data)
            window.location="http://localhost:8001/VerMatriz/"+props.anio+"&"+props.mes
        }

    return (
        <>
        <div>
            <Button inverted color='olive' onClick={matriz}>Mes: {props.mes}</Button>
        </div>
        </>
    )
}

export default ItemsMeses
