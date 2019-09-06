import React from 'react';
import './App.css';
import ServerForm from './components/ServerForm';
import Server from './components/Server';
import TenantForm from './components/TenantForm';

const App: React.FC = () => {
  const [serverCount, setServerCount] = React.useState<number>(0);

  const upServers = (count: number) => {
    setServerCount(count);
  }

  return (
    <div>
      <div className="container">

        <nav className="navbar navbar-expand-lg navbar-light" style={{ backgroundColor: "#e3f2fd" }}>
          <a className="navbar-brand" href="#">Distributed Counting</a>
        </nav>

        <div className="container">

          {serverCount}

          <div className="row">
            <div className="col-sm">
              <TenantForm />
            </div>
            <div className="col-sm">
              <ServerForm upServers={upServers} />
            </div>
          </div>
          {serverCount}
          <Server serverCount={serverCount} />
        </div>
      </div>
    </div>
  );
}

export default App;
