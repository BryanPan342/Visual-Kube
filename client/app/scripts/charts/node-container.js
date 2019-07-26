import React from 'react';
import { connect } from 'react-redux';
import { List as makeList } from 'immutable';
import { GraphNode } from 'weaveworks-ui-components';

import {
  getMetricValue,
  getMetricColor,
} from '../utils/metric-utils';
import { clickNode, enterNode, leaveNode } from '../actions/app-actions';
import { trackAnalyticsEvent } from '../utils/tracking-utils';
import { getNodeColor } from '../utils/color-utils';
import MatchedResults from '../components/matched-results';
import { GRAPH_VIEW_MODE } from '../constants/naming';

import NodeNetworksOverlay from './node-networks-overlay';
import { rgb } from 'polished';

import { changeColor } from '../components/nodes-jumbotron';
import { setNodeColor } from '../components/node-details';

var ashColorCode = 0;

class NodeContainer extends React.Component {

  saveRef = (ref) => {
    this.ref = ref;
  };

  handleMouseClick = (nodeId, ev) => {
    ev.stopPropagation();
    trackAnalyticsEvent('scope.node.click', {
      layout: GRAPH_VIEW_MODE,
      parentTopologyId: this.props.currentTopology.get('parentId'),
      topologyId: this.props.currentTopology.get('id'),
    });
    this.props.clickNode(nodeId, this.props.label, this.ref.getBoundingClientRect());
  };

  ashShape(ashShapeString)
  {
    if (ashShapeString === 'circle')
    {
      ashColorCode = 1;
      if (this.props.viewingNodeId) 
        {changeColor("jumboColor1");}

    }
    else if (ashShapeString === 'heptagon')
    {
      ashColorCode = 2;
      if (this.props.viewingNodeId) 
        {changeColor("jumboColor2");}
    }
    else if (ashShapeString === 'hexagon')
    {
      ashColorCode = 3;
      if (this.props.viewingNodeId) 
        {changeColor("jumboColor3");}
    }
    else if (ashShapeString === 'square')
    {
      ashColorCode = 4;
      if (this.props.viewingNodeId) 
        {changeColor("jumboColor4");}
    }

    if (ashShapeString != 'cloud')
    {
      ashShapeString = 'visualkube';
    }

    return ashShapeString;
  }

  ashColor(rank, label, pseudo)
  {
    if(this.props.error) {
      setNodeColor("rgba(255,0,0,0.7)");
      return rgb(255,0 ,0);
    }
    else if(this.props.parentErrors.has(this.props.id)){
      setNodeColor("rgba(255,165,0,0.7)");
      return rgb(255,165,0);
    }

    if (ashColorCode === 1)
    {
      setNodeColor("rgba(50, 7, 172, 0.7)");
      return rgb(50, 7, 172);
    }
    else if (ashColorCode === 2)
    {
      setNodeColor("rgb(71, 88, 239)");
      return rgb(71, 88, 239);
    }
    else if (ashColorCode === 3)
    {
      setNodeColor("rgb(7, 182, 220)");
      return rgb(7, 182, 220);
    }
    else if (ashColorCode === 4)
    {
      setNodeColor("rgb(82, 214, 214)");
      return rgb(82, 214, 214);
    }
    else {
      return rgb(0,0,0);
    }
  }

  ashMetric(metricFormattedValue)
  {
    if (metricFormattedValue)
    {
      return metricFormattedValue;
    }
    else {
      switch(ashColorCode) {
        case 1:
          return "Host";
        case 2:
          return "Pod";
        case 3:
          return "Cont.";
        case 4:
          return "App";
        default:
          return "Node";
      }
    }
  }

  renderPrependedInfo = () => {
    const { showingNetworks, networks } = this.props;
    if (!showingNetworks) return null;

    return (
      <NodeNetworksOverlay networks={networks} />
    );
  };

  renderAppendedInfo = () => {
    const matchedMetadata = this.props.matches.get('metadata', makeList());
    const matchedParents = this.props.matches.get('parents', makeList());
    const matchedDetails = matchedMetadata.concat(matchedParents);
    return (
      <MatchedResults matches={matchedDetails} searchTerms={this.props.searchTerms} />
    );
  };

  render() {
    const {
      rank, label, pseudo, metric, showingNetworks, networks
    } = this.props;
    const { hasMetric, height, formattedValue } = getMetricValue(metric);
   
    const metricFormattedValue = !pseudo && hasMetric ? formattedValue : '';
    const labelOffset = (showingNetworks && networks) ? 10 : 0;
    return (
      <GraphNode
        id={this.props.id}
        shape={this.ashShape(this.props.shape)}
        tag={this.props.tag}
        label={this.props.label}
        labelMinor={this.props.labelMinor}
        labelOffset={labelOffset}
        stacked={this.props.stacked}
        highlighted={this.props.highlighted}
        color={this.ashColor(rank, label, pseudo)}
        // color={getNodeColor(rank, label, pseudo)}
        size={this.props.size}
        isAnimated={this.props.isAnimated}
        contrastMode={this.props.contrastMode}
        forceSvg={this.props.exportingGraph}
        searchTerms={this.props.searchTerms}
        metricColor={getMetricColor(metric)}
        metricFormattedValue={this.ashMetric(metricFormattedValue)}
        metricNumericValue={height}
        renderPrependedInfo={this.renderPrependedInfo}
        renderAppendedInfo={this.renderAppendedInfo}
        onMouseEnter={this.props.enterNode}
        onMouseLeave={this.props.leaveNode}
        onClick={this.handleMouseClick}
        graphNodeRef={this.saveRef}
        x={this.props.x}
        y={this.props.y}
      />
    );
  }
}

function mapStateToProps(state) {
  return {
    parentErrors: state.get('errorParents'),
    contrastMode: state.get('contrastMode'),
    currentTopology: state.get('currentTopology'),
    exportingGraph: state.get('exportingGraph'),
    searchTerms: [state.get('searchQuery')],
    showingNetworks: state.get('showingNetworks'),
    nodeDetails: state.get('nodeDetails'),
    viewingNodeId: state.get("viewingNodeId"),
  };
}

export default connect(
  mapStateToProps,
  { clickNode, enterNode, leaveNode }
)(NodeContainer);
