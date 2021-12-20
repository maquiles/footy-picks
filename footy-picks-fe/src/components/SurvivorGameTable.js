import React from "react";
import { Card, Container, Table } from "react-bootstrap";


export function SurvivorGameTable(props) {
  return (
    <Container fluid>
      <Card bg="light" text="dark">
        <Card.Header as="h5">{props.gameTitle}</Card.Header>
        <Card.Body>
          <Table striped bordered hover>
            <thead>
              <th>Player</th>
              <th>Round 1</th>
              <th>Round 2</th>
              <th>Round 3</th>
            </thead>
            <tbody>
              <tr>
                <td>Player 1</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'red'}}>MNU</td>
              </tr>
              <tr>
                <td>Player 2</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'red'}}>MNU</td>
              </tr>
              <tr>
                <td>Player 3</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'red'}}>MNU</td>
              </tr>
              <tr>
                <td>Player 4</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'red'}}>MNU</td>
              </tr>
              <tr>
                <td>Player 5</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'green'}}>MNC</td>
                <td style={{ backgroundColor: 'green'}}>LIV</td>
              </tr>
              <tr>
                <td>Player 6</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'green'}}>MNC</td>
                <td style={{ backgroundColor: 'red'}}>TOT</td>
              </tr>
              <tr>
                <td>Player 7</td>
                <td style={{ backgroundColor: 'green'}}>CHE</td>
                <td style={{ backgroundColor: 'green'}}>MNC</td>
                <td style={{ backgroundColor: 'red'}}>TOT</td>
              </tr>
            </tbody>
          </Table>
        </Card.Body>
      </Card>
    </Container>
  );
}