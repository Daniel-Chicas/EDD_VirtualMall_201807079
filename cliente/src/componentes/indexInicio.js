import React from 'react';
import ReactDOM from 'react-dom';
import App from '../App';
import 'semantic-ui-css/semantic.min.css'
import './css/index.css'
import 'react-calendar/dist/Calendar.css';

ReactDOM.render(
      <React.StrictMode>
        <App/>
      </React.StrictMode>,
      document.getElementById('root')
);
