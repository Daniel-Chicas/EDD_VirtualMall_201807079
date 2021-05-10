import React from 'react'
import { Segment, Image } from 'semantic-ui-react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router} from 'react-router-dom'
import matriz from '../ImagenMatriz/Matriz.png'
import '../css/Imagen.css'
import DatosMatriz from './DatosMatriz'


function ImagenMatriz() {
    return (
        <div className="Matriz">
            <Router>
                <NavBar/>
            </Router>
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
