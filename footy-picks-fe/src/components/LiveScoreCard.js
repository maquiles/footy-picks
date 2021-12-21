import React from "react";
import { Card, Col } from "react-bootstrap";

export function LiveScoreCard(props) {
  return (
    <Col>
      <Card className="text-center" bg={'success'} text={'light'} style={{ width: '11rem' }}>
        <Card.Body>
          <Card.Title>
            {props.liveScore.homeTeam} {props.liveScore.homeScore} - {props.liveScore.awayScore} {props.liveScore.awayTeam}
          </Card.Title>
        </Card.Body>
      </Card>
    </Col>
  );
}