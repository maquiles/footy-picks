import './App.css';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Dashboard from './pages/dashboard';
import Games from './pages/games';
import Navbar from './components/navigation';


function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path='/' exact element={<Dashboard />} />
        <Route path='/games' element={<Games />} />
      </Routes>
    </Router>
  );
}

export default App;
