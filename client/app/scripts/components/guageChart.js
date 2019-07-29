import React from 'react';
import Plot from 'react-plotly.js';

class GuageChart extends React.Component {
  render() {
    return (
        <Plot
        

        data={[
        {
        values: [19, 26, 55],
        labels: ['Residential', 'Non-Residential', 'Utility'],
        type: 'pie'
        }
        ]}
        
        layout = {{height: 300, width: 500}}

        // Enter a speed between 0 and 180
        level = {{level: 175}}

        // Trig to calc meter point
        degrees = {{degree: 180 - level}}
        radius = {{radius:0.5}}
        radians = {{radians:degrees * Math.PI / 180}}
        x = {{x:radius*Math.cos(radians)}}
        y = {{y:radius*Math.sin(radians)}}

        // Path: may have to change to create a better triangle
        mainPath = {{one:'M -.0 -0.025 L .0 0.025 L ', 
            pathX:String(x),
            space :' ',
            pathY:String(y),
            pathEnd:' Z'}}
        path = {{path:mainPath.concat(pathX,space,pathY,pathEnd)}}

        data = {[ 
            {
            type: 'scatter',
            x: [0], y:[0],
                marker: {size: 28, color:'850000'},
                showlegend: false,
                name: 'speed',
                text: level,
                hoverinfo: 'text+name'},
            { values: [50/6, 50/6, 50/6, 50/6, 50/6, 50/6, 50],
            rotation: 90,
            text: ['TOO FAST!', 'Pretty Fast', 'Fast', 'Average',
                        'Slow', 'Super Slow', ''],
            textinfo: 'text',
            textposition:'inside',
            marker: {colors:['rgba(14, 127, 0, .5)', 'rgba(110, 154, 22, .5)',
                                    'rgba(170, 202, 42, .5)', 'rgba(202, 209, 95, .5)',
                                    'rgba(210, 206, 145, .5)', 'rgba(232, 226, 202, .5)',
                                    'rgba(255, 255, 255, 0)']},
            labels: ['151-180', '121-150', '91-120', '61-90', '31-60', '0-30', ''],
            hoverinfo: 'label',
            hole: .5,
            type: 'pie',
            showlegend: false
            }
        ]}

        layout = {{
            shapes:[{
                type: 'path',
                path: path,
                fillcolor: '850000',
                line: {
                    color: '850000'
                }
                }],
            title: '<b>Gauge</b> <br> Speed 0-100',
            height: 1000,
            width: 1000,
            xaxis: {zeroline:false, showticklabels:false,
                        showgrid: false, range: [-1, 1]},
            yaxis: {zeroline:false, showticklabels:false,
                        showgrid: false, range: [-1, 1]}
        }}
        />
        );
    }
}

export default GuageChart;

