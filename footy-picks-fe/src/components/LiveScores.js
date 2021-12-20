import React from "react";
import { Card, Container, Row } from "react-bootstrap";
import { LiveScoreCard } from "./LiveScoreCard";

export function LiveScores() {
  return (
    <Container fluid>
      <Card bg={'success'} text={'light'}>
          <Card.Header as="h4">Live Scores</Card.Header>
          <Card.Body>
            <Row md={ 9 }>
              <LiveScoreCard homeTeam="CHE" homeScore="1" awayTeam="MNU" awayScore="0" />
              <LiveScoreCard homeTeam="SCP" homeScore="1" awayTeam="MNC" awayScore="0" />
              <LiveScoreCard homeTeam="LIV" homeScore="1" awayTeam="PSG" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
              <LiveScoreCard homeTeam="RMA" homeScore="1" awayTeam="LIL" awayScore="0" />
            </Row>
          </Card.Body>
        </Card>
      </Container>
  );
}