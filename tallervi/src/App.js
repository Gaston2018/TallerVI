import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import './App.css';


import Index from './components/Index'
import CreateUser from './components/CreateUser'
import Turnos from './components/Turnos'
import Navigation from './components/Navigation'
import CreateClient from './components/CreateClients'
import CreateMascota from './components/CreateMascota'


function App() {
  return (
    <Router>
      <Navigation/>
      <Route path="/" exact component={Index}/>
      <Route path="/user" component={CreateUser}/>
      <Route path="/turnos" component={Turnos}/>
      <Route path="/cliente" component={CreateClient}/>
      <Route path="/mascota" component={CreateMascota}/>
    </Router>
  );
}

export default App;
