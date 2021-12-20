import React from "react";
import { Card, Container, Table } from "react-bootstrap";


export function ScorePredictorGameTable(props) {
  return (
    <Container fluid>
      <Card bg='light' text='dark'>
        <Card.Header as="h5">{props.gameTitle}</Card.Header>
        <Card.Body>
          <Table striped bordered hover>
            {/* TODO: table order matters */}
            {/* TODO: add tab to card that shows score picks for current round */}
            <thead>
              <th>Player</th>
              <th>Points</th>
              <th>Exact</th>
              <th>Close</th>
              <th>Outcome</th>
            </thead>
            <tbody>
              <tr>
                <td>Player 1</td>
                <td>12</td>
                <td>2</td>
                <td>2</td>
                <td>3</td>
              </tr>
              <tr>
                <td>Player 2</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
              <tr>
                <td>Player 3</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
              <tr>
                <td>Player 4</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
              <tr>
                <td>Player 5</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
              <tr>
                <td>Player 6</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
              <tr>
                <td>Player 7</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
            </tbody>
          </Table>
        </Card.Body>
      </Card>
    </Container>  
  );
}