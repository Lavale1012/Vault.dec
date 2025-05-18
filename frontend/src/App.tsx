import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home_page from "./pages/Home/Home_page";
import NavBar from "./components/NavBar";

function App() {
  return (
    <Router>
      <NavBar />
      <Routes>
        <Route path="/" element={<Home_page />} />
      </Routes>
    </Router>
  );
}

export default App;
