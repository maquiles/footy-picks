import React from "react";
import { Card, Col, Container, Row } from "react-bootstrap";
import { getActiveGames } from "../repo/gameService";
import { MockUserPlayer } from "../repo/getMockData";
import SurvivorGameTable from "./SurvivorGameTable";


export default class ActiveGames extends React.Component {
  constructor() {
    super();
    this.state = {
      games: {},
      fetched: false
    };
  }

  componentDidMount() {
    getActiveGames(MockUserPlayer)
      .then(response => response.json())
      .then(body => {
        this.setState({
          fetched: true,
          games: body
        });
      });
  }

  handleFetching() {
    if (!this.state.fetched) {
      return <p>Fetching data...</p>
    } else {
      return (
        <Row>
          {this.state.games.map(function(game) {
            return (
              <Col key={game.name}>
                <SurvivorGameTable gameTitle={game.name} league={game.league} rows={game.rows}/>
              </Col>
            );
          })}          
        </Row>
      );
    }
  }

  render() {
    return (
      <Container fluid>
        <Card bg="dark" text="light">
          <Card.Body>
            <Card.Title as="h4">Active Games</Card.Title>
            {this.handleFetching()}
          </Card.Body>
        </Card>
      </Container>
    );
  }
}