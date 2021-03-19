import React from 'react'
import { Segment, Image } from 'semantic-ui-react'
import arbol from '../ImagenArbol/ArbolProductos.png'
import '../css/Imagen.css'

function VerArbol() {
    return (
        <div className="Arbol">
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>
            <div>
            <Segment placeholder className="FondoImagen" centered>
            <Image src={arbol} size='300px' centered />
            </Segment>
            </div>
        </div>
    )
}

export default VerArbol