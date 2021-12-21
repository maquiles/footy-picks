import './App.css';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Dashboard from './pages/dashboard';
import Games from './pages/games';
import Navbar from './components/navigation';
import LiveScores from './components/LiveScores';
import NewGameForm from './pages/NewGameForm';
import React from 'react';
import { getActiveGames, getPastGames } from './services/gameService';
import { MockUserPlayer } from './services/getMockData';


export default class App extends React.Component {
  constructor() {
    super();

    const player = MockUserPlayer;
    localStorage.activeGames = JSON.stringify(getActiveGames(player));
    localStorage.pastGames = JSON.stringify(getPastGames(player));
  }

  render() {
    return (
      <Router>
        <Navbar />
        <br />
        <Routes>
          <Route path='/' exact element={<Dashboard />} />
          <Route path='/games' element={<Games />} />
          <Route path='/games/new' element={<NewGameForm />} />
          <Route path='/scores' element={<LiveScores />} />
        </Routes>
      </Router>
    );
  }
}