import React from 'react';
import { Badge, Button } from 'reactstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

export default class ErrorToggle extends React.Component {
    constructor(props) {
        super(props);
        // this.state = {isToggleOn: true};

        // This binding is necessary to make `this` work in the callback
        // this.handleClick = this.handleClick.bind(this);
    }

    // handleClick() {
    //     this.setState(state => (
    //         {isToggleOn: !state.isToggleOn}
    //     ));

    //     this.props.changeVisibility();
    // }

  render() {
    let style = 'err-toggle ';
    if (this.props.newError) {
      style = style.concat('err-new-err');
    } else {
      style = style.concat('err-no-new-err');
    }
    
    return (
      <div>
        <Button onClick={this.props.onClick} className={style} color="danger" outline>
          <Badge color="danger">{this.props.getNumErrors()}</Badge> Alerts 
        </Button>
      </div>
    );
  }
}