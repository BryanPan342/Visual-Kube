import React from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';
import { Card, CardImg } from 'reactstrap';
import TutorialCarousel from './tutorial-carousel';
import 'bootstrap/dist/css/bootstrap.min.css';

import img1 from './images/describedfigmahierarchy.png';

class TutorialModal extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        modal: false
      };
  
      this.toggle = this.toggle.bind(this);
    }
  
    toggle() {
      this.setState(prevState => ({
        modal: !prevState.modal
      }));
    }
  
    render() {
      return (
        <div>
          <Button color="primary" onClick={this.toggle}>{this.props.buttonLabel}Tutorial</Button>
          <Modal id="AshModal" isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
            <ModalHeader toggle={this.toggle}>Tutorial</ModalHeader>
            <ModalBody>
            <Card>
              <CardImg top width="100%" src={img1} alt="Stuff should be here" />
            </Card>
            </ModalBody>
            <ModalFooter>
              <Button color="primary" onClick={this.toggle}>Close</Button>
            </ModalFooter>
          </Modal>
        </div>
      );
    }
  }
  
  export default TutorialModal;