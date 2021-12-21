export class SurvivorGame {
  constructor(id, title, players, startDate, active, league) {
    this.id = id;
    this.title = title;
    this.players = players;
    this.startDate = startDate;
    this.active = active;
    this.league = league;
  }
}

export class SurvivorRound {
  constructor(round, date, matches, game) {
    this.round = round;
    this.date = date;
    this.matches = matches;
    this.game = game;
  }
}

export class SurvivorPick {
  constructor(player, pick, round, game) {
    this.player = player;
    this.pick = pick;
    this.round = round;
    this.game = game;
  }
}