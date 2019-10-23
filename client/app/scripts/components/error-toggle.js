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
      style = style.concat('err-new-err ');
    } else {
      style = style.concat('err-no-new-err ');
    }

    let style2 = '';
    if (this.props.getNumErrors() === 0){
      style2 = style2.concat('err-disabled');
      style = style.concat('err-disabled');
    }

    return (
      <div>
        <Button onClick={this.props.onClick} className={style} color="danger" outline>
          <Badge color="danger" className={style2} >{this.props.getNumErrors()}</Badge> Alerts 
        </Button>
      </div>
    );
  }
}