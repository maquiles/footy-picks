import React from "react";
import { Button, Col, Container, Form } from "react-bootstrap";
import { renderMatches } from "react-router-dom";


export default class NewGameForm extends React.Component {
  constructor() {
    super();
    this.state = {
      privateGame: false
    };

    this.handleCheck = this.handleCheck.bind(this);
  }

  handleCheck(e) {
    this.setState({ privateGame: e.target.checked});
  }

  onFormSubmit(e) {
    e.preventDefault()
    const formData = new FormData(e.target);
    const formDataObj = Object.fromEntries(formData.entries());
    console.log(formDataObj);
  }

  render() {
    return (
      <Container fluid>
        <Col xs lg="2">
          <Form onSubmit={this.onFormSubmit}>
            <Form.Floating className="mb-3">
              <Form.Control name="gameTitle" type="text" placeholder="Game Title"/>
              <label htmlFor="gameTitle">Game Title</label>
            </Form.Floating>

            <Form.Control as="select" className="mb-3" name="league">
              <option>Select League</option>
              <option value="EPL">English Premier League - EPL</option>
              <option value="UCL">UEFA Champions League - UCL</option>
            </Form.Control>

            <Form.Control as="select" className="mb-3" name="gametype">
              <option>Select Game Type</option>
              <option value="survivor">Survivor</option>
              <option value="predictor">Score Predictor</option>
            </Form.Control>

            <Form.Group className="mb-3">
              <Form.Check 
                id="privateGame" 
                type="checkbox" 
                label="Private Game" 
                style={{ color: 'white'}}
                checked={this.state.privateGame}
                onChange={this.handleCheck}/>
              <Form.Control name="gamePassword" hidden={!this.state.privateGame} type="password" placeholder="Private Game Code" />
            </Form.Group>

            <Button variant="primary" type="submit">Create Game</Button>
          </Form> 
        </Col>
      </Container>
    );
  }
}