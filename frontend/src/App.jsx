import "./App.css";
import Search from "./search";
import About from "./About";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Search />} />
      <Route path="/about" element={<About />} />
    </Routes>
  );
}
