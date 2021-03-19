import React from 'react'
import { Segment, Image } from 'semantic-ui-react'
import matriz from '../ImagenMatriz/Matriz.png'
import '../css/Imagen.css'
import DatosMatriz from './DatosMatriz'


function ImagenMatriz(props) {
    return (
        <div className="Matriz">
            <div>
            <Segment placeholder className="FondoImagen" centered>
            <Image src={matriz} size='300px' centered />
            </Segment>
            </div>
            <DatosMatriz/>
        </div>
    )
}

export default ImagenMatriz
