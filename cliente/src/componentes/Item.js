import React from 'react'
import { List } from 'semantic-ui-react'
import Mes from './ItemsMeses'
import { Label } from 'semantic-ui-react'

function Item(props) {
    return (
        <List.Item>
                <Label.Group color='orange'>
                    <Label as='a'>Año: {props.anio}</Label>
                </Label.Group>
            <List.List>
                {props.mes.map((c) => (
                    <Mes
                        anio = {props.anio}
                        mes = {c.MesA}
                    />
                ))}
            </List.List>
        </List.Item>
    )
}

export default Item
