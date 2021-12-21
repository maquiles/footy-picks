import React from "react";
import { Card, Container, Table } from "react-bootstrap";

function getCellColor(correct) {
  if (correct == true) {
    return 'green';
  } else if (correct == false) {
    return 'red';
  } else if (correct == null) {
    return 'white';
  }
}

export default class SurvivorGameTable extends React.Component {
  constructor(props) {
    super(props);

    this.rows = props.rows;
    this.numRounds = this.rows.length > 0 ? this.rows[0]["rounds"].length : 0;
  }

  render() {
    return (
      <Container fluid>
        <Card bg="light" text="dark">
          <Card.Header as="h5">{this.props.gameTitle} - {this.props.league}</Card.Header>
          <Card.Body>
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>Player</th>
                  {[...Array(this.numRounds)].map((x, i) => {
                    const round = i+1;
                    const colTitle = "Round "+round;
                    return (<th key={i}>{colTitle}</th>);
                  })}
                </tr>
              </thead>
              <tbody>
                {this.rows.map(function(row) {
                  return (
                    <tr key={row["player"]}>
                      <td>{row["player"]}</td>
                      {row["rounds"].map(function(round) {
                        return (
                          <td key={round["round"]} style={{backgroundColor: getCellColor(round["correct"])}}>{round["pick"]}</td>
                        );
                      })}
                    </tr>
                  );
                })}
              </tbody>
            </Table>
          </Card.Body>
        </Card>
      </Container>
    );
  }
}