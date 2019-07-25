import React from 'react';
import NodesChart from '../charts/nodes-chart';

// var jumboColor = "jumboColor1";

export function changeColor(incomingJumboColor)
{
  this.state.jumboColor = incomingJumboColor;
  this.forceUpdate();
}

class NodesJumbotron extends React.Component {
  constructor(props) {
    super(props);

    this.state = {jumboColor: ""};

    changeColor = changeColor.bind(this);
  }

    render() {
    return (
      <div id="NodeJumbotron" className={this.state.jumboColor}>
        <h5>Host: Minikube / Pod: AshPod / Container: AshContainer / </h5>
        <NodesChart/>
      </div>
    );
  }
}

export default NodesJumbotron;