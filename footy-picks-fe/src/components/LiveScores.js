import React from "react";
import { Card, Container, Row } from "react-bootstrap";
import { LiveScoreCard } from "./LiveScoreCard";
import LiveScore from "../models/LiveScore"

export default class LiveScores extends React.Component {
  constructor() {
    super();
    this.games = [
      new LiveScore("CHE", "MNU", "4", "0"),
      new LiveScore("SCP", "MNC", "3", "0"),
      new LiveScore("LIV", "PSG", "1", "2"),
      new LiveScore("RMA", "LIL", "4", "0"),
      new LiveScore("BAY", "ATL", "5", "1"),
      new LiveScore("VIL", "BEN", "1", "2"),
      new LiveScore("INT", "BAR", "1", "3"),
      new LiveScore("MIL", "AJA", "2", "2"),
      new LiveScore("DOR", "RBL", "3", "1"),
    ];
  }
  
  render() {
    return (
      <Container fluid>
        <Card bg={'success'} text={'light'}>
          <Card.Header as="h4">Live Scores</Card.Header>
          <Card.Body>
            <Row md={ 9 }>
              {this.games.map(function(object, i){
                return <LiveScoreCard key={"live-score-"+i} liveScore={object}/>
              })}
            </Row>
          </Card.Body>
        </Card>
      </Container>
    );
  }
}