import axios from 'axios';
import { BACKEND_URL, LOGIN_URL } from "./utils";

export function login(loginBody) {
  return fetch(BACKEND_URL+LOGIN_URL, {
    method: 'POST',
    body: JSON.stringify(loginBody),
    credentials: 'include'
  });
}
