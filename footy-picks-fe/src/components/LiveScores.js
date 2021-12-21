import React from "react";
import { Card, Container, Row } from "react-bootstrap";
import { LiveScoreCard } from "./LiveScoreCard";
import Score from "../models/Score"

export default class LiveScores extends React.Component {
  constructor() {
    super();
    this.games = [
      new Score("CHE", "MNU", "4", "0"),
      new Score("SCP", "MNC", "3", "0"),
      new Score("LIV", "PSG", "1", "2"),
      new Score("RMA", "LIL", "4", "0"),
      new Score("BAY", "ATL", "5", "1"),
      new Score("VIL", "BEN", "1", "2"),
      new Score("INT", "BAR", "1", "3"),
      new Score("MIL", "AJA", "2", "2"),
      new Score("DOR", "RBL", "3", "1"),
    ];
  }
  
  render() {
    return (
      <Container fluid>
        <Card bg={'dark'} text={'light'}>
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