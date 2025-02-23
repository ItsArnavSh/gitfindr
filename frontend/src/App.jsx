import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import SearchResults from "./searchResults";
import Search from "./search";
import About from "./About";

function App() {
  return (
    <>
      <Search />
      <About />
    </>
  );
}

export default App;
