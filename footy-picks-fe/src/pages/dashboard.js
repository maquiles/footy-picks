import React from "react";
import { Container, Row } from "react-bootstrap";
import LiveScores from "../components/LiveScores";
import ActiveGames from "../components/ActiveGames";

export default class Dashboard extends React.Component {
  constructor() {
    super();
  }

  render() {
    return (
      <Container fluid>
        <Row>
          <LiveScores />
        </Row>
        <br />
        <Row>
          <ActiveGames />
        </Row>
      </Container>
    );
  }
}