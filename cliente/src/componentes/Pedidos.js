import {React,  useState} from 'react'
import { PlaceholderParagraph } from 'semantic-ui-react'
import '../css/Pedidos.css'
import Tree from './Tree'
const axios=require('axios').default

function Pedidos() {
    const [anios, setanios] = useState([])
    async function obtener(){
        if(anios.length===0){
        const data = await axios.get('http://localhost:3000/DatosMatriz')
        if (data.data.General !== null) {
            setanios(data.data.General)
            }
        }
    }
    obtener()
    return (
        <div className="Pedidos">
            <Tree 
                listaAnios={anios}
            />
        </div>
    )
}

export default Pedidos
