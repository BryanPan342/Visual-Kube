import React from 'react';
import NodesChart from '../charts/nodes-chart';

var jumboColor = "jumboColor1";

export function changeColor(incomingJumboColor)
{
  jumboColor = incomingJumboColor;
  this.forceUpdate();
}

class NodesJumbotron extends React.Component {
  constructor(props) {
    super(props);

    changeColor = changeColor.bind(this);
  }

    render() {
    return (
      <div id="NodeJumbotron" className={jumboColor}>
        <h5>Host: Minikube / Pod: AshPod / Container: AshContainer / </h5>
        <NodesChart/>
      </div>
    );
  }
}

export default NodesJumbotron;