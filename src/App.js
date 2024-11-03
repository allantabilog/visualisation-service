import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("Loading...");
  const [serverTimestamp, setServerTimestamp] = useState("");
  useEffect(() => {
    fetch("http://localhost:8080/")
      .then((response) => response.json())
      .then((data) => setMessage(data.message))
      .catch((error) => console.log(error));
  });

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Setting up the visualisation service (watch this space)</p>
        <p>Response from server: {message}</p>
      </header>
    </div>
  );
}

export default App;
