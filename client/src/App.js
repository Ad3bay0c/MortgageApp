import { Fragment } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Navbar from './components/layout/Navbar';
import Home from './components/pages/Home';
import Calculator from './components/pages/Calculator';
import BankState from './context/bank/bankState';
import Alert from './components/layout/Alert';
// import ContactState from './context/contact/ContactState';
import './App.css';

const App = () => {
  return (
    <BankState>
      <Router>
        <Fragment>
          <Navbar />
          <div className="container">
            <Alert />
            <Switch>
              <Route exact path='/' component={Home} />
              <Route exact path='/calculator' component={Calculator} />
            </Switch>
          </div>
        </Fragment>
      </Router>
    </BankState>
  );
}

export default App;
