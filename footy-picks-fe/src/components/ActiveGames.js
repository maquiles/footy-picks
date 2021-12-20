import React from "react";
import { Card, Col, Container, Row } from "react-bootstrap";
import { SurvivorGameTable } from "./SurvivorGameTable";
import { ScorePredictorGameTable } from "./ScorePredictorGameTable";


export function ActiveGames() {
  return (
    <Container fluid>
      <Card bg="dark" text="light">
        <Card.Body>
          <Card.Title as="h4">Active Games</Card.Title>
          <Row>
            <Col>
              <SurvivorGameTable gameTitle="Test Survivor Game" />
            </Col>
            <Col>
              <ScorePredictorGameTable gameTitle="Test Score Predictor Game" />
            </Col>
            <Col>
              <SurvivorGameTable gameTitle="Test Survivor Game 2" />
            </Col>           
          </Row>
        </Card.Body>
      </Card>
    </Container>
  );
}