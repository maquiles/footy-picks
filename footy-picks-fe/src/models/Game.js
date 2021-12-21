export default class Game {
  constructor(name, gametype, startDate, status, league, private, passcode) {
    this.name = name;
    this.gametype = gametype;
    this.startDate = startDate;
    this.status = status;
    this.league = league;
    this.private = private;
    this.passcode = passcode;
  }
}