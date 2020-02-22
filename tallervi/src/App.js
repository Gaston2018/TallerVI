import React from 'react';
import logo from './logo.svg';
//import background from '..public/img/Index.jpg'
import './App.css';

function App() {
  return (
    <div className="App" >
      <header class="header">
        <div class="menu">
                <img src="../public/img/logo.png" alt="logo" width="120" class="logo"/>
                <nav class="nav">
                    <ul>
                        <li><a href="#">Solicitar turno</a></li>
                        <li><a href="#">Nuevo paciente</a></li>
                        <li><a href="#">Nuevo veterinario</a></li>
                    </ul>
                </nav>
        </div>

    </header>
    <div id = 'contenedor'>
      <table id='turnos'>
        <thead>
          <td>mascota</td>
        </thead>
        </table> 

    </div>
    <section>
        <article>

        </article>
    </section>
    <aside>

    </aside>
    <footer>

    </footer>
    </div>
  );
}

export default App;
