import axios from 'axios'
import { BACKEND_URL, PLAYER_URL } from "./utils"

export function createNewAccount(player) {
  return axios.post(BACKEND_URL+PLAYER_URL, { player })
}