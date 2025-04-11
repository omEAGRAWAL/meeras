import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { BrowserRouter, Routes, Route } from "react-router";
import Home from './pages/Home'
import Signup from './pages/Signup';
import venue from "./pages/Venue.jsx";
import Venue from "./pages/Venue.jsx";

function App() {
  const [count, setCount] = useState(0)

  return (
    <BrowserRouter>
    <Routes>
    
      <Route path="/" element={<Home />} />
      <Route path="/signup" element={<Signup />} />
      <Route path="/v" element={<Venue/>} />
    </Routes>
  </BrowserRouter>
  )
}

export default App
