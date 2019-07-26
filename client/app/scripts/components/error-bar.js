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

export const ErrorIcon = () => <Icon icon={warning} size={18}/>;

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



var isVisible = true;

export function changeVisibility(){
  isVisible = !isVisible;
  this.forceUpdate();
}

var numErrors = 0;

export function setNumErrors(nodes) {
  numErrors = nodes.length;
}

export function getNumErrors() {
  return numErrors;
}

export class ErrorBar extends React.Component {
  constructor(props){
    super(props);
    this.error_data = new Map();

    changeVisibility = changeVisibility.bind(this);
  }

  onClickErr(ev, node) {
    this.props.clickNode(node.id, node.label, ev.target.getBoundingClientRect(), 'pods');
  }
  render() {
    const { isDashboardViewMode } = this.props;
    var data = this.props.state.get('errorData').toList().toJS();
    setNumErrors(data);
    var allGoodMsg = false;
   if (data.length === 0 && isDashboardViewMode) {
    allGoodMsg = true;
   }
   if (isVisible && data[0] && Array.isArray(data)){
    return (
      <div className='err-bar' >
        { allGoodMsg ? 
          <div>You have no errors! All good!</div> :
          <div>
            {data.map((element) => 
              <Toast >
                <ToastBody className="err-item" onClick={ev => this.onClickErr(ev, element)} >
                  {/* <div style={{color: '#000000'}}><ErrorIcon  /></div>  */}
                  <div></div>
                  <div> <i className="fas fa-exclamation-circle error-icon"></i> &nbsp; <b>[Container] {element.label} &nbsp;- </b> &nbsp; {element.status}</div>
                </ToastBody>      
              </Toast>
            )}
          </div>
       }  
      </div>
    );
  }
    else {
      return(<div></div>);
    }
  }
}

const mapStatetoProps = (state) => ({
  state: state,
  current_nodes: state.get('nodesByTopology'),
  nodes: shownNodesSelector(state),
  currentTopology: state.get('currentTopology'),
  isDashboardViewMode: isDashboardViewModeSelector(state)
	})

const mapDispatchToProps = (dispatch) => ({
    getNodesbyTopology: (topoId) => dispatch(getNodesbyTopology(topoId)),
    clickNode: (id, label, ev, pod) => dispatch(clickNode(id, label, ev, pod))
})

export default connect(
  mapStatetoProps, mapDispatchToProps
)(ErrorBar);   
