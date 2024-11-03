import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";

function App() {
  // initialise a piece of state that will hold
  // json data
  const [message, setMessage] = useState("<Loading message...>");
  const [serverTimestamp, setServerTimestamp] = useState(
    "<Loading timestamp...>"
  );
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/");
        const data = await response.json();
        setMessage(data.message);
        setServerTimestamp(data.timestamp);
      } catch (error) {
        console.log(error);
      }
    };
    fetchData();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Setting up the visualisation service (watch this space)</p>
        <p>
          Response from server:{" "}
          <span style={{ color: "yellow" }}>{message}</span> on{" "}
          <span style={{ color: "yellow" }}>
            {new Date(serverTimestamp).toLocaleString()}
          </span>
        </p>
      </header>
    </div>
  );
}

export default App;
