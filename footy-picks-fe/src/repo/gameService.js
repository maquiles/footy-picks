import axios from 'axios';
import { activeGameOne, activeGameTwo } from "./getMockData";
import { BACKEND_URL, GAMES_URL } from "./utils";

export function getActiveGames(player) {
  return axios.get(BACKEND_URL+GAMES_URL+"?user=Matthew");
}

export function getPastGames(player) {
  // TODO: call the backend to get the user's active games

  return [activeGameOne()];
}