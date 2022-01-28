import React from "react";
import { Button, Card, Col, Container, Form, Row } from "react-bootstrap";

export default class SignUp extends React.Component {
  constructor() {
    super();
    
    this.state = {
      created: false,
      creationError: false,
    }

    this.onFormSubmit = this.onFormSubmit.bind(this);
  }

  onFormSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formDataObj = Object.fromEntries(formData.entries());
    // TODO: make request to backend and redirest to login or catch error
    this.setState({created: true});
  }

  render() {
    if (this.state.created) {
      return (
        <Container fluid>
          <Row>
            <Col xs md="4"/>
            <Col xs md="4">
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