import React from 'react';
import { Breadcrumb, BreadcrumbItem } from 'reactstrap';
import { statement } from '@babel/template';
import { clickNode } from '../actions/app-actions';
import { connect } from 'react-redux';
import 'bootstrap/dist/css/bootstrap.min.css';
import { clickRelative } from '../actions/app-actions';
import { trackAnalyticsEvent } from '../utils/tracking-utils';
import MatchedText from './matched-text';
import {clickShowTopologyForNode} from '../actions/app-actions';
import ActionTypes from '../constants/action-types';
<<<<<<< HEAD
import {handleShowTopologyForNode} from './node-details'


export class BreadCrumb extends React.Component{
=======
import {handleShowTopologyForNode} from './node-details';
import {getTopoFromId, getLabelAndParentsFromId, getNodeDetails} from '../utils/web-api-utils';
import { clickTopology } from '../actions/app-actions';
import { onTopologyClick} from './topologies'

// import {handleShowTopologyForNode} from '../actions/app-actions'

export class BreadCrumb extends React.Component{

  returnLabel = (label) => {
    return label;
  }

>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
  getLabel(){
    if(this.props.details.toList().toJS()[0]){
      if(this.props.details.toList().toJS()[0]['details']){
        // console.log(this.props.details.toList().toJS()[0]['details']['label'])
        return (this.props.details.toList().toJS()[0]['details']['label'])
      }
    }
  }

  handleClick(ev, id, topologyId) {
    ev.preventDefault();
    this.props.clickRelative(
      id,
      topologyId,
    );
  }

<<<<<<< HEAD
  // clickShowTopologyForNode(topologyId, nodeId) {
  //   console.log('Click Show Topology FOr Node')
  //   return (dispatch, getState) => {
  //     dispatch({
  //       nodeId,
  //       topologyId,
  //       type: ActionTypes.CLICK_SHOW_TOPOLOGY_FOR_NODE
  //     });
  //     updateTopology(dispatch, getState);
  //   };
  // }

  handleShowTopologyForNode = (ev, topologyId, nodeId) => {
    ev.preventDefault();
    console.log('this.props.topId = ', topologyId)
    console.log('this.props.nodeId = ', nodeId)
    this.props.clickShowTopologyForNode(topologyId, nodeId);
    // this.props.clickShowTopologyForNode(topologyId, nodeId);
  }


  // ORIGINAL HANDLECLICK
  // handleClick(ev) {
  //   ev.preventDefault();
  //   trackAnalyticsEvent('scope.node.relative.click', {
  //     topologyId: this.props.topologyId,
  //   });
  //   this.props.dispatch(clickRelative(
  //     this.props.id,
  //     this.props.topologyId,
  //     this.props.label,
  //     this.node.getBoundingClientRect()
  //   ));
  // }

=======
  handleShowTopologyForNode = (ev, topologyId, nodeId) => {
    // console.log("WORKS")
    ev.preventDefault();
    this.props.clickShowTopologyForNode(topologyId, nodeId);
  }

>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
  saveNodeRef(ref) {
    this.node = ref;
  }

  makeBreadcrumb(level){
    if(this.props.details.toList().toJS()[0]){
      if(this.props.details.toList().toJS()[0]['details']){
        if(this.props.details.toList().toJS()[0]['details']['parents']){
<<<<<<< HEAD
          // Keep if we get rid of processes breadcrumb
=======
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
          if(level == 'processes'){
            return null;
          }
          // the host is the last element in the array
          else if(level == 'pods'){
            let parents = this.props.details.toList().toJS()[0]['details']['parents']
            let id = parents[parents.length-1]['id']
            let topologyId = parents[parents.length-1]['topologyId']
            return(
              <div>
              <BreadcrumbItem className = "breadcrumbitem"
              // onClick={ev => this.handleClick(ev,id, topologyId)}
<<<<<<< HEAD
              onClick={ev => this.handleShowTopologyForNode(ev, topologyId, id)}
              >
                <span className = 'level'>Host :    </span> 
              {parents[parents.length-1]['label'].toString()}
              </BreadcrumbItem>
              <BreadcrumbItem className = "breadcrumbitem">
                <span className = 'level'>Pod :       </span>
                {this.getLabel()}
=======
              // onClick={ev => this.handleShowTopologyForNode(ev, getTopoFromId(parents[parents.length-1]['id']), parents[parents.length-1]['id'])}
              // onClick={ev => this.handleShowTopologyForNode(ev, topologyId, id)}
              >
                <span className = 'level'>Host:    </span>
              {parents[parents.length-1]['label'].toString()}
              </BreadcrumbItem>
              <BreadcrumbItem className = "breadcrumbitem">
                <b>
                <span className = 'level'>Pod:       </span>
                {this.getLabel()}
                </b>
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
                </BreadcrumbItem>
              </div>
            );
          }
<<<<<<< HEAD
          else if(level == 'containers'){
=======
          else if(level == 'containers' && this.props.details.toList().toJS()[0]['details']['parents'][2]){
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
            let parents = this.props.details.toList().toJS()[0]['details']['parents']
            return(
              <div>
              <BreadcrumbItem 
              className = "breadcrumbitem"
<<<<<<< HEAD
              onClick={ev => this.handleShowTopologyForNode(ev, parents[2]['topologyId'], parents[2]['id'])}
              // onClick={ev => this.handleClick(
              //   ev,
              //   parents[2]['id'], 
              //   parents[2]['topologyId'])}
=======
              // onClick={ev => this.handleShowTopologyForNode(ev)}
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
                ref={this.saveNodeRef}>
                <span className = 'level'>Host:   </span> {this.props.details.toList().toJS()[0]['details']['parents'][2]['label']}
                </BreadcrumbItem>
              <BreadcrumbItem 
              className = "breadcrumbitem"
<<<<<<< HEAD
              onClick={ev => this.handleShowTopologyForNode(ev, parents[1]['topologyId'], parents[1]['id'])}
              // onClick={ev => this.handleClick(
              //   ev,
              //   parents[1]['id'], 
              //   parents[1]['topologyId'])}
                ref={this.saveNodeRef}
              ><span className = 'level'>Pod:   </span> {this.props.details.toList().toJS()[0]['details']['parents'][1]['label']}</BreadcrumbItem>
              <BreadcrumbItem className = "breadcrumbitem">
                <span className = 'level'>Container:   </span>{this.getLabel()}</BreadcrumbItem>
              </div>
            );
          }
          // Commented out because we got rid of processes breadcrumb
          // else if(level == 'processes'){
          //   if(this.props.details.toList().toJS()[0]['details']['parents'][0] && this.props.details.toList().toJS()[0]['details']['parents'][2]){
          //     return(<div>
          //       <BreadcrumbItem className = "breadcrumbitem"><a href=''>{this.props.details.toList().toJS()[0]['details']['parents'][2]['label'].toString()}</a></BreadcrumbItem>
          //       <BreadcrumbItem className = "breadcrumbitem"><a href=''>{this.props.details.toList().toJS()[0]['details']['parents'][1]['label'].toString()}</a></BreadcrumbItem>
          //       <BreadcrumbItem className = "breadcrumbitem" active>{this.getLabel()}</BreadcrumbItem>
          //       </div>)
          //   }
          //   else{
          //     return(
          //       <div>
          //       <BreadcrumbItem className = "breadcrumbitem"><a href=''>{this.props.details.toList().toJS()[0]['details']['parents'][0]['label']}</a></BreadcrumbItem>
          //       <BreadcrumbItem className = "breadcrumbitem" active>{this.getLabel()}</BreadcrumbItem>
          //       </div>
          //     );
          //   }
          // }
        }
        if(level != 'Processes'){
          return(<BreadcrumbItem className = "breadcrumbitem"> <span className = 'level'> Host:   </span>{this.getLabel()}</BreadcrumbItem>)
=======
              // onClick={ev => this.handleShowTopologyForNode(ev)}
                ref={this.saveNodeRef}>
                <span className = 'level'>Pod:   </span> {this.props.details.toList().toJS()[0]['details']['parents'][1]['label']}</BreadcrumbItem>
              <BreadcrumbItem className = "breadcrumbitem">
                <b>
                <span className = 'level'>Container:   </span>{this.getLabel()}
                </b></BreadcrumbItem>
              </div>
            );
          }
        }
        if(level != 'Processes'){
          return(<BreadcrumbItem className = "breadcrumbitem"
          // onClick={ev => this.handleShowTopologyForNode(ev, this.props.details.toList().toJS()[0]['topologyId'], this.props.details.toList().toJS()[0]['id'])}
          ><b><span className = 'level'> Host:   </span>{this.getLabel()}</b></BreadcrumbItem>)
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
        }
      }
    }
  }

<<<<<<< HEAD
  render() {
    let label = this.getLabel();
=======

  render() {
    let label = this.getLabel();
    // IF A NODE IS SELECTED
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
    if(this.props.details.toList().toJS()[0]){
      if(this.props.details.toList().toJS()[0]['details']){
        return (
          <div>
            <Breadcrumb className = "breadcrumb">
<<<<<<< HEAD
              {this.makeBreadcrumb(this.props.details.toList().toJS()[0]['topologyId'])}
=======
              {/* {console.log("TOPOLOGYID 1 = ", this.props.details.toList().toJS()[0]['topologyId'])}
              {console.log("TOPOLOGYID 2 = ", getTopoFromId(this.props.details.toList().toJS()[0]['id']))} */}
              {this.makeBreadcrumb(getTopoFromId(this.props.details.toList().toJS()[0]['id']))}
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
            </Breadcrumb>
          </div>
        );
      }
      else{
<<<<<<< HEAD
        return (<Breadcrumb></Breadcrumb>)
      }
    }
    else{
=======
        // if(this.props.viewingNodeId){
        // }
        return (<Breadcrumb></Breadcrumb>)
      }
    }
    // VIEWING NODE, NO SELECTED NODE : 
    else{
      if(this.props.viewingNodeId){
        if(getTopoFromId(this.props.viewingNodeId) == 'hosts'){
          console.log("BREADCRUMB TOPOLOGY = ", this.props.topologies.toList().toJS())
          return(
            <Breadcrumb className = "breadcrumb">
            <BreadcrumbItem className = "breadcrumbitem"
            onClick={ev => this.onTopologyClick(ev, this.props.topology)}
            // onClick={ev => this.handleShowTopologyForNode(ev)}
            >
            <b><span className = 'level'> Host:  </span>
            {this.props.breadcrumb[0]}
            </b>
            </BreadcrumbItem>
            </Breadcrumb>
            )
        }
        else if(getTopoFromId(this.props.viewingNodeId) == 'pods' && this.props.breadcrumb[2]["2"]){
          return(
            <Breadcrumb className = "breadcrumb">
            <BreadcrumbItem className = "breadcrumbitem"
            onClick={ev => this.handleShowTopologyForNode(ev, this.props.breadcrumb[2]['2'].topologyId, this.props.breadcrumb[2]['2'].id)}> 
            <span className = 'level'> Host:  </span>
            {this.props.breadcrumb[2]["2"].label}
            {/* {this.props.breadcrumb[1]["0"]} */}
            </BreadcrumbItem>
            <BreadcrumbItem className = "breadcrumbitem"
            // onClick={ev => this.handleShowTopologyForNode(ev, getTopoFromId(this.props.breadcrumb[1]), this.props.breadcrumb[1])}
            > 
            <b><span className = 'level'> Pod:  </span>
            {this.props.breadcrumb[0]}
            </b>
            </BreadcrumbItem>
            </Breadcrumb>
          )
        }
      }
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
      return (<Breadcrumb></Breadcrumb>)
    }
  }
}

// export default Example;

function mapStatetoProps(state){
    return {
<<<<<<< HEAD
   state: state,
   details : state.get('nodeDetails')
=======
   topologies: state.get('topologies'),
   state: state,
   details : state.get('nodeDetails'),
   viewingNodeId : state.get('viewingNodeId'),
   breadcrumb : state.get('breadcrumb')
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
    };
}

export default connect(
 mapStatetoProps,
<<<<<<< HEAD
 { clickRelative, clickShowTopologyForNode }
=======
 { clickRelative, 
  clickShowTopologyForNode, 
  handleShowTopologyForNode, 
  getNodeDetails, 
  clickTopology,
  onTopologyClick }
>>>>>>> b3e1c13e6473ed5ce64cf1f9efb68aecd4c69c6c
)(BreadCrumb);