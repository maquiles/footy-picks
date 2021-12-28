import { activeGameOne, activeGameTwo } from "./getMockData";

const BACKEND_URL = "http://localhost:8000"
const GAMES_URL = "/games"

export function getActiveGames(player) {
  return fetch(BACKEND_URL+GAMES_URL+"?user=Matthew")
}

export function getPastGames(player) {
  // TODO: call the backend to get the user's active games

  return [activeGameOne()];
}