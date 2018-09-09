import React, { Component } from 'react';
import { Provider } from 'react-redux';
import store from './store';
import BillList from './components/Billlist';
import BillForm from './components/Billform';

import './App.css';

class App extends Component {
  render() {
    return (
      <Provider store={ store }>
        <div className="App">
          <header className="App-header">
            <h1 className="App-title">Test header</h1>
          </header>

          <p className="App-intro">
            Ting tong
          </p>

          <BillList/>

          <hr/>

          <BillForm/>

          <footer>
            © Kaspars Kadikis, 2018
          </footer>
        </div>
      </Provider>
    );
  }
}

export default App;
