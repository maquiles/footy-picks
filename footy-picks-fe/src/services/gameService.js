import { activeGameOne, activeGameTwo } from "./getMockData";

export function getActiveGames(player) {
  // TODO: call the backend to get the user's active games

  return [activeGameOne(), activeGameTwo()];
}

export function getPastGames(player) {
  // TODO: call the backend to get the user's active games

  return [activeGameOne()];
}