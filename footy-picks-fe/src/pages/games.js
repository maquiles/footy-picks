import React from "react";
import { Accordion, Button, Card, Col, Container, ListGroup, Row, Tab } from "react-bootstrap";
import SurvivorGameTable from "../components/SurvivorGameTable";

export default class Games extends React.Component {
  constructor() {
    super();

    this.activeGames = JSON.parse(localStorage.activeGames);
    this.pastGames = JSON.parse(localStorage.pastGames);
  }

  render() {
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
                      {this.activeGames.map((game) => {
                        const href = "#active-game-"+game.id; 
                        return <ListGroup.Item action href={href} key={game.id}>{game.name}</ListGroup.Item>
                      })}
                    </ListGroup>
                  </Accordion.Body>
                </Accordion.Item>
                <Accordion.Item eventKey="1">
                  <Accordion.Header>Past Games</Accordion.Header>
                  <Accordion.Body>
                    <ListGroup>
                      {this.pastGames.map((game) => {
                        const href = "#past-game-"+game.id; 
                        return <ListGroup.Item action href={href} key={game.id}>{game.name}</ListGroup.Item>
                      })}
                    </ListGroup>
                  </Accordion.Body>
                </Accordion.Item>
              </Accordion>
              <Button className="mt-2 mb-2" variant="primary" href="/games/new"> New Game</Button> 
            </Col>
            <Col>
              <Tab.Content>
                <Tab.Pane eventKey="#newGame" style={{ color: "white"}}>Trying to create a new game...</Tab.Pane>
                <Tab.Pane eventKey="#link1" style={{ color: "white"}}>Choose a game brah</Tab.Pane>
                {this.activeGames.map((game) => {
                  const href = "#active-game-"+game.id;
                  return (
                    <Tab.Pane eventKey={href} key={game.id}>
                      <SurvivorGameTable gameTitle={game.name} league={game.league} rows={game.rows} />
                    </Tab.Pane>
                  );
                })}
                {this.pastGames.map((game) => {
                  const href = "#past-game-"+game.id;
                  return (
                    <Tab.Pane eventKey={href} key={game.id}>
                      <SurvivorGameTable gameTitle={game.name} league={game.league} rows={game.rows} />
                    </Tab.Pane>
                  );
                })}
              </Tab.Content>
            </Col>
          </Tab.Container>
        </Row>
      </Container>
    );
  }
}