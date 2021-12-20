import React from "react";
import { Accordion, Button, Card, Col, Container, ListGroup, Row, Tab } from "react-bootstrap";
import { SurvivorGameTable } from "../components/SurvivorGameTable";

function Games() {
  return (
    <Container fluid>
      <Row>
        <Tab.Container defaultActiveKey="#link1">
          <Col sm={2}>
            <Accordion>
              <Accordion.Item eventKey="0">
                <Accordion.Header>Active Games</Accordion.Header>
                <Accordion.Body>
                  <ListGroup>
                    <ListGroup.Item action href="#game1">Game 1</ListGroup.Item>
                    <ListGroup.Item action href="#game2">Game 2</ListGroup.Item>
                    <ListGroup.Item action href="#game3">Game 3</ListGroup.Item>
                    <ListGroup.Item action href="#game4">Game 4</ListGroup.Item>
                  </ListGroup>
                </Accordion.Body>
              </Accordion.Item>
              <Accordion.Item eventKey="1">
                <Accordion.Header>Past Games</Accordion.Header>
                <Accordion.Body>
                  <ListGroup>
                    <ListGroup.Item action href="#pastGame1">Past Game 1</ListGroup.Item>
                    <ListGroup.Item action href="#pastGame2">Past Game 2</ListGroup.Item>
                    <ListGroup.Item action href="#pastGame3">Past Game 3</ListGroup.Item>
                    <ListGroup.Item action href="#pastGame4">Past Game 4</ListGroup.Item>
                  </ListGroup>
                </Accordion.Body>
              </Accordion.Item>
            </Accordion>
            <Button className="mt-2" variant="primary" href="/games/new"> New Game</Button> 
          </Col>
          <Col>
            <Tab.Content>
              <Tab.Pane eventKey="#newGame" style={{ color: "white"}}>Trying to create a new game...</Tab.Pane>
              <Tab.Pane eventKey="#link1" style={{ color: "white"}}>Choose a game brah</Tab.Pane>
              <Tab.Pane eventKey="#game1">
                <SurvivorGameTable gameTitle="Game 1"/>
              </Tab.Pane>
              <Tab.Pane eventKey="#game2">
                <SurvivorGameTable gameTitle="Game 2"/>
              </Tab.Pane>
              <Tab.Pane eventKey="#game3">Link 1</Tab.Pane>
              <Tab.Pane eventKey="#game4">Link 2</Tab.Pane>
              <Tab.Pane eventKey="#pastGame1">Link 1</Tab.Pane>
              <Tab.Pane eventKey="#pastGame2">Link 2</Tab.Pane>
              <Tab.Pane eventKey="#pastGame3">Link 1</Tab.Pane>
              <Tab.Pane eventKey="#pastGame4">Link 2</Tab.Pane>
            </Tab.Content>
          </Col>
        </Tab.Container>
      </Row>
    </Container>
  );
}

export default Games;