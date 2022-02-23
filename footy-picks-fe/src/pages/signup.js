import React from "react";
import { Button, Card, Col, Container, Form, Row } from "react-bootstrap";
import { createNewAccount } from "../repo/playerService";

export default class SignUp extends React.Component {
  constructor() {
    super();
    
    this.state = {
      created: false,
      creationError: false,
      loginConfErr: false
    }

    this.onFormSubmit = this.onFormSubmit.bind(this);
  }

  onFormSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formDataObj = Object.fromEntries(formData.entries());

    var pwConf = formDataObj["login"] == formDataObj["confirm-login"]
    if (!pwConf) {
      this.setState({loginConfErr: true});
      return;
    }

    delete formDataObj["confirm-login"];
    createNewAccount(formDataObj)
      .then(response => {
        if (!response.ok) {
          throw new Error('Error creating new player >> ' + response.statusText);
        } else {
          response.json();
        }
      })
      .then(body => {
        this.setState({
          created: true,
          creationError: false,
          loginConfErr: false
        });
      })
      .catch((error) => {
        this.setState({creationError: true})
      });
  }

  render() {
    if (this.state.created) {
      return (
        <Container fluid>
          <Row>
            <Col xs md="4"/>
            <Col md="4">
              <Card bg={'dark'} text={'light'}>
                <Card.Header as="h4">Account Created</Card.Header>
                <Card.Body>
                  <Button variant="primary" type="submit" href="/login">Log In</Button>
                </Card.Body>
              </Card>
            </Col>
            <Col xs md="4"/>
          </Row>
        </Container>
      );
    }

    let loginConfirmed;
    if (this.state.loginConfErr) {
      loginConfirmed = <p style={{color: "#FF0000"}}>Passwords don't match. Try again.</p>
    }

    let creationErr;
    if (this.state.creationError) {
      creationErr = <p style={{color: "#FF0000"}}>Error creation new account. Please try again.</p>
    }

    return (
      <Container fluid>
        <Row>
          <Col xs md="4"/>
          <Col xs md="4">
            <Card bg={'dark'} text={'light'}>
              <Card.Header as="h4">Sign Up</Card.Header>
              <Card.Body>
                <Form onSubmit={this.onFormSubmit}>
                  <Form.Control className="mb-2" name="email" type="text" placeholder="Email"/>
                  <Form.Control className="mb-2" name="name" type="text" placeholder="Name"/>
                  <Form.Control className="mb-2" name="login" type="password" placeholder="Password"/>
                  <Form.Control className="mb-2" name="confirm-login" type="password" placeholder="Confirm Password"/>
                  {loginConfirmed}
                  {creationErr}
                  <Button variant="primary" type="submit">Create</Button>
                </Form>
              </Card.Body>
            </Card>
          </Col>
          <Col xs md="4"/>
        </Row>
      </Container>
    )
  }
}