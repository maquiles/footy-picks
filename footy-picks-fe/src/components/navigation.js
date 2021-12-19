import React from "react";
import { Link } from "react-router-dom"
import { Navbar, Nav, Container } from "react-bootstrap"

const Navigation = () => {
  return (
    <Navbar collapseOnSelect fixed='top' expand='sm' bg='dark' variant='dark'>
      <Container>
        <Navbar.Brand href='/'>Footy Picks</Navbar.Brand>
        <Navbar.Toggle aria-controls='reponsive-navbar-nav' />
        <Navbar.Collapse id='responsive-navbar-nav'>
          <Nav>
            <Nav.Link href='/'>Dashboard</Nav.Link>
            <Nav.Link href='/games'>Games</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Navigation;