import React from "react";
import { Button, Card, Col, Container, Form, Row } from "react-bootstrap";

export default class Login extends React.Component {
  constructor() {
    super();
  }

  onFormSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formDataObj = Object.fromEntries(formData.entries());
    // TODO: make login call to backend and update fe with 
  }

  render() {
    return (
      <Container fluid>
        <Row>
          <Col xs md="4"/>
          <Col xs md="4">
            <Card bg={'dark'} text={'light'}>
              <Card.Header as="h4">Enter Credentials</Card.Header>
              <Card.Body>
                <Form onSubmit={this.onFormSubmit}>
                  <Form.Control className="mb-2" name="email" type="text" placeholder="Email"/>
                  <Form.Control className="mb-2" name="password" type="password" placeholder="Password"/>
                  <Button variant="primary" type="submit">Login</Button>
                </Form>
              </Card.Body>
            </Card>
            <p style={{color: "#FFFFFF"}}>
              Don't have an account? &ensp;
              <a href='/signup'>Sign Up</a>
            </p>
          </Col>
          <Col xs md="4"/>
        </Row>
      </Container>
    )
  }
}