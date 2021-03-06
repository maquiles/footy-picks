import React from "react";
import { Navbar, Nav, Container } from "react-bootstrap";

const Navigation = () => {
  return (
    <Navbar collapseOnSelect fixed='top' expand='sm' bg='dark' variant='dark'>
      <Container>
        <Navbar.Brand href='/'>Footy Picks</Navbar.Brand>
        <Navbar.Toggle aria-controls='reponsive-navbar-nav' />
        <Navbar.Collapse id='responsive-navbar-nav'>
          <Nav>
            <Nav.Link href='/scores'>Scores</Nav.Link>
            <Nav.Link href='/games'>Games</Nav.Link>
            <Nav.Link href='/login'>Login</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Navigation;