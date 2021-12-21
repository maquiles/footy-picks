import React from "react";
import { Card, Col, Container, Row } from "react-bootstrap";
import SurvivorGameTable from "./SurvivorGameTable";


export default class ActiveGames extends React.Component {
  constructor() {
    super();
    this.games = JSON.parse(localStorage.activeGames);
  }

  render() {
    return (
      <Container fluid>
        <Card bg="dark" text="light">
          <Card.Body>
            <Card.Title as="h4">Active Games</Card.Title>
            <Row>
              {this.games.map(function(game) {
                return (
                  <Col key={game.name}>
                    <SurvivorGameTable gameTitle={game.name} league={game.league} rows={game.rows}/>
                  </Col>
                )
              })}          
            </Row>
          </Card.Body>
        </Card>
      </Container>
    );
  }
}