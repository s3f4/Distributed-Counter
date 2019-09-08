import React from 'react';
import './App.css';
import NodeForm from './components/NodeForm';
import NodeList from './components/NodeList';
import TenantForm from './components/TenantForm';
import { getNodes, upNodes } from './api';

const App: React.FC = () => {
  const [nodesData, setNodesData] = React.useState<any>([]);
  const [error, setError] = React.useState<any>(false);

  const nodesRequest = () => {
    setError(false);
    getNodes().then(data => {
      if (data.error) {
        setError(data.error.message);
      } else {
        setNodesData(data);
      }
    });
  }

  const up = (count: number) => {
    upNodes(count).then(data => {
      if (data.error) {
        setError(data.error.message);
      } else {
        nodesRequest();
      }
    });
  }

  React.useEffect(() => {
    nodesRequest();
  }, []);


  const alert = () => {
    if (error) {
      return (
        <div className="mt-3 alert alert-danger" role="alert">
          {error}
        </div>
      );
    }
  }

  return (
    <div>
      <div className="container">
        <nav className="navbar navbar-expand-lg navbar-light" style={{ backgroundColor: "#e3f2fd" }}>
          <a className="navbar-brand" href="/">Distributed Counting</a>
        </nav>
        <div className="container">
          {alert()}
          <div className="row">
            <div className="col-sm">
              <TenantForm reRender={nodesRequest} />
            </div>
            <div className="col-sm">
              <NodeForm reRender={nodesRequest} upNodes={up} />
            </div>
          </div>
          <NodeList nodes={nodesData.nodes} />
        </div>
      </div>
    </div>
  );
}

export default App;
