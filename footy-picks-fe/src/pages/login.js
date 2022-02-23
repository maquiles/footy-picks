import React from "react";
import { Button, Card, Col, Container, Form, Row } from "react-bootstrap";
import { login } from "../repo/loginService";

export default class Login extends React.Component {
  constructor() {
    super();

    this.state = {
      loginError: false
    }

    this.onFormSubmit = this.onFormSubmit.bind(this);
  }

  onFormSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const formDataObj = Object.fromEntries(formData.entries());
  
    login(formDataObj)
      .then(response => {
        if (response.ok) {
          console.log(response.headers.get('Set-Cookie'));
          return response.json();
        } else {
          throw new Error('Error when logging in >> ' + response.statusText);
        }
      })
      .then(body => {
        this.setState({loginErrror: false});
        console.log(body);
      })
      .catch((error) => {
        this.setState({loginError: true})
      });
  }

  render() {
    let loginErr;
    if (this.state.loginError) {
      loginErr = <p style={{color: "#FF0000"}}>Error logging into account. Please try again.</p>
    }

    return (
      <Container fluid>
        <Row>
          <Col xs md="4"/>
          <Col md="4">
            <Card bg={'dark'} text={'light'}>
              <Card.Header as="h4">Enter Credentials</Card.Header>
              <Card.Body>
                <Form onSubmit={this.onFormSubmit}>
                  <Form.Control className="mb-2" name="email" type="text" placeholder="Email"/>
                  <Form.Control className="mb-2" name="login" type="password" placeholder="Password"/>
                  {loginErr}
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