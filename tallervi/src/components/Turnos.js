import React, { Component } from 'react'
import Axios from 'axios'

export default class Turnos extends Component {

    state = {
        turnos: [],
        turno: ''
    }

    async componentDidMount() {
        const res = await Axios.get('http://fierce-harbor-58840.herokuapp.com/turnos');
        this.setState({ turnos: res.data });
        console.log(this.state.turnos)
    }

    onChangeTurnos = (e) => {
        this.setState({
            turno: e.target.value
        })

    }

    onSubmit = e => {
        Axios.post('http://fierce-harbor-58840.herokuapp.com/turnos', {
            turno: this.state.turnos
        })
        e.preventDefault();
    }
    render() {
        return (
            <div className="row">
                <div className="col-md-4">
                    <div className="card card-body">
                        <h3>Crear Turno</h3>
                        <form onSubmit={this.onSubmit}>
                            <div className="form-group">
                                <input 
                                    type="text"
                                    className="form-control"
                                    onChange={this.onChangeTurnos}
                                    />
                            </div>
                            <button type="submit" className="btn btn-primary">
                                Save
                            </button>
                        </form>
                    </div>
            </div>
                <div className="col-md-8">
                    <ul className="list-group">
                        {
                            this.state.turnos.map(turnos => (
                                <li className="list-group-item list-group-action" key={turnos.id_turno}>
                                {turnos.id_turno}
                            </li>))
                        }
                    </ul>
                </div>
            </div>
        )
    }
}
