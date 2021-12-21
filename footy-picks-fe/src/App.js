import './App.css';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Dashboard from './pages/dashboard';
import Games from './pages/games';
import Navbar from './components/navigation';
import LiveScores from './components/LiveScores';
import NewGameForm from './pages/NewGameForm';


function App() {
  return (
    <Router>
      <Navbar />
      <br />
      <Routes>
        <Route path='/' exact element={<Dashboard />} />
        <Route path='/games' element={<Games />} />
        <Route path='/games/new' element={<NewGameForm />} />
        <Route path='/scores/live' element={<LiveScores />} />
      </Routes>
    </Router>
  );
}

export default App;
