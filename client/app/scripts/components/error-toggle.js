import React from 'react';
import { Badge, Button } from 'reactstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

export default class ErrorToggle extends React.Component {
    constructor(props) {
        super(props);
    }

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