import React from 'react'
import { List } from 'semantic-ui-react'
import Item from './Item'

function Tree(props) {
    return (
        <div>
            <List ordered>
                {props.listaAnios.map((c, index) => (
                    <Item
                        anio = {c.Anio}
                        mes = {c.Meses}
                    />
                ))}
            </List>
        </div>
    )
}

export default Tree
