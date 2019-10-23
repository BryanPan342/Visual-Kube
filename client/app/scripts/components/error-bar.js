import React from 'react';
import { Toast, ToastBody} from 'reactstrap';
import { connect } from 'react-redux';
import { shownNodesSelector } from '../selectors/node-filters';
import { getNodesbyTopology } from '../actions/app-actions';

import { warning } from 'react-icons-kit/typicons/warning';
import { Icon } from 'react-icons-kit';
import { GRAPH_VIEW_MODE } from '../constants/naming';
import { trackAnalyticsEvent } from '../utils/tracking-utils';
import { isDashboardViewModeSelector } from '../selectors/topology'

import { APIcall } from '../utils/web-api-utils';
import { clickNode } from '../actions/app-actions';

import ErrorToggle from './error-toggle';

export const ErrorIcon = () => <Icon icon={warning} />;

export const index_topoById = (topo, data) => {
 var sanity_check;
 for(var i = 0; i < data.length; i++){
   sanity_check = Object.keys(data[i])[0]
   if(sanity_check && sanity_check.slice(-topo.length) === topo)
     return i;
 }
 return -1;
}

export function formatData(nodes, topologyId){
  var return_data;
  if(topologyId === "pods")
    return_data = [];
  else if (topologyId === "hosts")
    return_data = { cpu: { value: 0, max: 0}, memory: { value: 0, max: 0}};
 
  if(!nodes.get(topologyId))
    return return_data;

  var data = nodes.get(topologyId).toList().toJS();
  let i, currNode, test;
  for(i = 0; i < data.length; i++){
    currNode = data[i]
    if(topologyId === "pods" && currNode.hasOwnProperty("parents") && currNode['metadata'][0]['value'] !== "Running"){
      test = APIcall(currNode.metadata[2].value, currNode.id, currNode.label).then(data => {console.log(data); return data;});
      console.log(test);
      return_data.push(test);
    }
    else if(topologyId === "hosts" && data[i]['metrics']){
      return_data={cpu: {value: data[i]['metrics'][0]['value'], max: data[i]['metrics'][0]['max']}, memory: {value: data[i]['metrics'][1]['value'], max: data[i]['metrics'][1]['max']}}
    }
  }
  return return_data;
}



var numErrors = 0;

export function setNumErrors(nodes) {
  numErrors = nodes.length;
}

export class ErrorBar extends React.Component {
  constructor(props){
    super(props);
    this.error_data = new Map();

    this.state = {
      newError: false,
      isToggleOn: true
    }
  }

  getNumErrors = () => {
    return numErrors;
  }

  componentDidUpdate(prevProps) {
    if (this.props.data.length > prevProps.data.length) {
      this.setState({
        newError: true,
      })
    }
  }

  clickToggle = () => {
    this.setState({
      newError: false,
      isToggleOn: !this.state.isToggleOn
    })

  }

  onClickErr(ev, node) {
    this.props.clickNode(node.id, node.label, ev.target.getBoundingClientRect(), 'pods');
  }

  render() {
    const { isDashboardViewMode, data } = this.props;
    const { newError, isToggleOn } = this.state;
    setNumErrors(data);
    var allGoodMsg = false;
   if (data.length === 0 && isDashboardViewMode) {
    allGoodMsg = true;
   }
   if (data[0] && Array.isArray(data)){
    return (
      <div className='err-bar' >
        { !isDashboardViewMode && <ErrorToggle 
          isToggleOn={isToggleOn} 
          onClick={this.clickToggle} 
          newError={newError} 
          getNumErrors={this.getNumErrors}
          />
        }
        { allGoodMsg ? 
          <div>You have no errors! All good!</div> : 
          (isToggleOn ? 
            <div>
              {data.map((element) => 
                <Toast >
                  <ToastBody className="err-item" onClick={ev => this.onClickErr(ev, element)} >
                    {/* <ErrorIcon /> */}
                    <b className='err-icon'> <i className="fas fa-exclamation-circle"></i> &nbsp; Container: {element.label} </b> <b> &nbsp; - &nbsp; {element.status} </b>
                    </ToastBody>      
                </Toast>
              )}
            </div> : 
            <div></div>
          )
        }
        </div>
    );
  }
    else {
      return(<div>
        { !isDashboardViewMode && <ErrorToggle
          isToggleOn={isToggleOn} 
          onClick={this.clickToggle} 
          newError={newError} 
          getNumErrors={this.getNumErrors}
          />
        }
      </div>);
    }
  }
}

const mapStatetoProps = (state) => ({
  state: state,
  current_nodes: state.get('nodesByTopology'),
  nodes: shownNodesSelector(state),
  currentTopology: state.get('currentTopology'),
  isDashboardViewMode: isDashboardViewModeSelector(state),
  data: state.get('errorData').toList().toJS()
	})

const mapDispatchToProps = (dispatch) => ({
    getNodesbyTopology: (topoId) => dispatch(getNodesbyTopology(topoId)),
    clickNode: (id, label, ev, pod) => dispatch(clickNode(id, label, ev, pod))
})

export default connect(
  mapStatetoProps, mapDispatchToProps
)(ErrorBar);   
