import React from 'react';
import ReactDOM from 'react-dom';
import Inicio from './componentes/Inicio';
import 'semantic-ui-css/semantic.min.css'
import './css/index.css'
import 'react-calendar/dist/Calendar.css';

ReactDOM.render(
      <React.StrictMode>
        <Inicio/>
      </React.StrictMode>,
      document.getElementById('root')
);
