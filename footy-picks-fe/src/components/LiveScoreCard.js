import React from "react";
import { Card, Col } from "react-bootstrap";

export function LiveScoreCard(props) {
  return (
    <Col>
      <Card className="text-center" bg={'dark'} text={'light'} style={{ width: '11rem' }}>
        <Card.Body>
          <Card.Title>{props.homeTeam} {props.homeScore} - {props.awayScore} {props.awayTeam}</Card.Title>
        </Card.Body>
      </Card>
    </Col>
  );
}