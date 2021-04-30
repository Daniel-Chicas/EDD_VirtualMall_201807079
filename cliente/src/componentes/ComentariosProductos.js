import {React, useEffect, useState} from 'react'
import '../css/Comentarios.css'
import { Button, Header, Icon, Segment, Form, Comment, Input, Label} from 'semantic-ui-react'
import NavBar from '../componentes/NavBar'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import Com from '../componentes/ComentarioProducto'
const axios=require('axios').default

function ComentariosProductos() {
    const [tiendas, settiendas] = useState([])
    const [listado, setlistado] = useState("")
    const [usuario, setusuario] = useState("")
    var URLactual = window.location;
    var direccion = URLactual.toString().split("/")
    var direccion2 = direccion[4].split("%20")
    var cadena = ""
    for (let step = 0; step < direccion2.length; step++) {
        if(step === 0){
            cadena = direccion2[step]
        }else{
            cadena = cadena+" "+direccion2[step]
        }
      }
    
    useEffect(()=>{
        async function obtener(){
            if(tiendas.length===0){
                const data = await axios.get('http://localhost:3000/ComentariosP/'+cadena)
                if (data.data.General !== null) {
                    settiendas(data.data.General)
                }
            }
            const user = await axios.get('http://localhost:3000/DatosLinea')
            if (user.data !== null) {
                setusuario(user.data)
            }
        }
        obtener()
    })

    const comentar = async()=>{
        let comment = localStorage.getItem('Comentario')
        var datos = JSON.parse(comment)
        var comm = document.getElementById("Area").value
        const DPI = usuario
        const Comentario = comm
        const Fecha = Date(Date.now())

        if (datos === null) {
            const Producto = parseInt(cadena.split("&")[3])
            const Comentarios = []
            const Tienda = ""
            const Departamento = ""
            const Calificacion = ""

        var viene={
            DPI,
            Comentario,
            Fecha
        }

        Comentarios.push(viene)
        
        var json={
            Tienda,
            Departamento,
            Calificacion,
            Producto,
            Comentarios
        }
            console.log(JSON.stringify(json))
            axios.post("http://localhost:3000/Comentarios/"+cadena, JSON.stringify(json) , {headers:{ 'Content-Type':'multipart/form-data'}})
            .then(response=>{
                console.log(response.data);
            }).catch(error=>{
                console.log(error);
            })
        }else{
            var viene={
                DPI,
                Comentario,
                Fecha
            }
            datos.Comentarios.push(viene)
            datos.Producto = parseInt(cadena.split("&")[3])
            axios.post("http://localhost:3000/Comentarios/"+cadena, datos , {headers:{ 'Content-Type':'multipart/form-data'}})
            .then(response=>{
                console.log(response.data);
            }).catch(error=>{
                console.log(error);
            })
        }

        const data = await axios.get('http://localhost:3000/ComentariosP/'+cadena)
        if (data.data.General !== null) {
            settiendas(data.data.General)
        }
        localStorage.clear()
    }

    return (
        <div>
                <Router>
                    <NavBar/>   
                </Router>
            <div className="GeneralLogin">
                <div className="Comentarios">
                <Comment.Group threaded>
                <Header as='h3' dividing>
                    <h1>{"COMENTARIOS ACERCA DEL PRODUCTO"}</h1>
                </Header>
                {tiendas.map((dato, index) => (
                    <Com    
                        Nombre = {dato.Dpi}
                        Fecha = {dato.Fecha}
                        Comentario = {dato.Comentario}
                        Respuestas = {dato.Respuestas}
                        Todo = {tiendas[index]}                   
                    />
                ))}
                <Form reply>
                <Label color="brown" size="huge">Agrega tu comentario :D</Label>
                <Form.TextArea id="Area"/>
                <Button content='Comentar' labelPosition='right' icon='edit' size="huge" primary onClick={comentar}/>
                </Form>
            </Comment.Group>
                </div>
            </div>
        </div>
    )
}

export default ComentariosProductos
