import React, { Component } from 'react'
import Axios from 'axios'

export default class Turnos extends Component {
    
    state = {
        turnos :[]
    }

    async componentDidMount() {
        const res = await Axios.get('http://fierce-harbor-58840.herokuapp.com/turnos');
        console.log(res)
    }

    render() {
        return (
            <div>
                CREAR TURNOS
            </div>
        )
    }
}
