import React from 'react'
import { Button, Header, Icon, Segment, Form, Comment, Input} from 'semantic-ui-react'
import Respuesta from '../componentes/Respuesta'
import '../css/Comentarios.css'

function Comentario(props) {
    var bandera = props.bandera


    const responder = () =>{
        var final = []
        var casi = []
        var resp =props.Nombre+"|"+props.Comentario+"|"+props.Fecha+"&"+props.Padre
        var sep = resp.split("&")
        for (let index = 0; index < sep.length; index++) {
            if (sep[index] !== "undefined") {
                casi.push(sep[index])
            }
        }
        
        for (let index = 0; index < casi.length; index++) {
            var a = casi[index].split("->")
            if (a[1] === "undefined") {
                final.push(casi[index])
            }
            if(index === casi.length-1){
                var b = casi[index].split("->")
                final.push(b[0])
                final.push(b[1])
            }         
        }
        for (let index = 0; index < final.length; index++) {
            if (final[index] === undefined) {
                final[index] = ""
            }
        }
        
        const fin = final.reverse()
        var ete = []
        for (let index = 0; index < fin.length; index++) {
            if (final[index] !== "") {
                ete.push(final[index])
            }
        }

        const Producto = -1
        const Comentarios = []
        var DPI = 0
        var Comentario = ""
        var Fecha = ""
        const Tienda = ""
        const Departamento = ""
        const Calificacion = ""

        for (let index = 0; index < ete.length; index++) {
            var a = ete[index].split("|")
            DPI = parseInt(a[0])
            Comentario = a[1]
            Fecha = a[2]
            var viene = {
                DPI,
                Comentario,
                Fecha
            }
            Comentarios.push(viene)
        }


        var json={
            Tienda,
            Departamento,
            Calificacion,
            Producto,
            Comentarios
        }

        localStorage.setItem('Comentario', JSON.stringify(json))
    }
    
    if (props.Respuestas !== null && bandera === undefined) {
        return (
            <div className="Comment">
                <Comment>
                <Comment.Avatar as='a' src='http://placeimg.com/640/480/cats' />
                <Comment.Content>
                    <Comment.Author as='a'>{props.Nombre}</Comment.Author>
                    <Comment.Metadata>
                    <span >{props.Fecha}</span>
                    
                    </Comment.Metadata>
                    <Comment.Text>{props.Comentario}</Comment.Text>
                    <Comment.Actions>
                    <a onClick={responder}>Responder</a>
                    </Comment.Actions>
                    
                    {props.Respuestas.map((dato, index) => (
                        <Respuesta   
                            Padre = {props.Nombre+"|"+props.Comentario+"|"+props.Fecha+"&"+props.Padre}
                            Nombre = {dato.Dpi}
                            Fecha = {dato.Fecha}
                            Comentario = {dato.Comentario}
                            Respuestas = {dato.Respuestas}
                            Todo = {props.Todo}
                        />
                    ))}
                </Comment.Content>
                </Comment>
                <br/>
            </div>
        )
    }else if (bandera === "Respuesta" && props.Respuestas !== null) {
        return (
            <div className="Comment">
                <Comment>
                <Comment.Content>
                    {props.Respuestas.map((dato, index) => (
                        <Respuesta    
                            Padre = {props.Nombre+"&"+props.Padre}
                            Nombre = {dato.Dpi}
                            Fecha = {dato.Fecha}
                            Comentario = {dato.Comentario}
                            Respuestas = {dato.Respuestas}
                            Todo = {props.Todo}
                        />
                    ))}
                </Comment.Content>
                </Comment>
                <br/>
            </div>
        )
    }else if (bandera !== "Respuesta" && props.Respuestas === null){
        return (
            <div className="Comment">
                <Comment>
                <Comment.Avatar as='a' src='http://placeimg.com/640/480/cats' />
                <Comment.Content>
                    <Comment.Author as='a' id="Nombre">{props.Nombre}</Comment.Author>
                    <Comment.Metadata>
                    <span >{props.Fecha}</span>
                    </Comment.Metadata>
                    <Comment.Text>{props.Comentario}</Comment.Text>
                    <Comment.Actions>
                    <a onClick={responder} text-right>Responder</a>
                    </Comment.Actions>
                </Comment.Content>
                </Comment>
                <br/>
            </div>
        )
    }else{
        return(
            <div>
                
            </div>
        )
    }
}

export default Comentario
