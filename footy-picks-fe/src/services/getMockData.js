import Player from "../models/Player"

export function MockUserPlayer() {
  return new Player("Matthew", "mattaquiles@gmail.com", ["1", "2", "3", "4"]);
}

export function activeGameOne() {
  return {
    id: "1",
    name: "Test Survivor Game",
    league: "EPL",
    rows: [
      {
        player: "Matthew",
        rounds: [
          {
            round: 1,
            pick: "CHE",
            correct: true
          },
          {
            round: 2,
            pick: "MNU",
            correct: true
          },
          {
            round: 3,
            pick: "LIV",
            correct: true
          },
          {
            round: 4,
            pick: "MNC",
            correct: null
          }
        ]
      },
      {
        player: "Player 2",
        rounds: [
          {
            round: 1,
            pick: "CHE",
            correct: true
          },
          {
            round: 2,
            pick: "MNU",
            correct: true
          },
          {
            round: 3,
            pick: "LIV",
            correct: false
          }
        ]
      },
      {
        player: "Player 3",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 4",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 5",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 6",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      }
    ]
  }
}

export function activeGameTwo() {
  return {
    id: "2",
    name: "Test Survivor Game 2",
    league: "UCL",
    rows: [
      {
        player: "Matthew",
        rounds: [
          {
            round: 1,
            pick: "CHE",
            correct: true
          },
          {
            round: 2,
            pick: "MNU",
            correct: true
          },
          {
            round: 3,
            pick: "LIV",
            correct: true
          },
          {
            round: 4,
            pick: "MNC",
            correct: null
          }
        ]
      },
      {
        player: "Player 2",
        rounds: [
          {
            round: 1,
            pick: "CHE",
            correct: true
          },
          {
            round: 2,
            pick: "MNU",
            correct: true
          },
          {
            round: 3,
            pick: "LIV",
            correct: false
          }
        ]
      },
      {
        player: "Player 3",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 4",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 5",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      },
      {
        player: "Player 6",
        rounds: [
          {
            round: 1,
            pick: "CHE"
          },
          {
            round: 2,
            pick: "MNU"
          },
          {
            round: 3,
            pick: "LIV"
          },
          {
            round: 4,
            pick: "MNC"
          }
        ]
      }
    ]
  }
}