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
  const [fibCount, setFibCount] = useState(10);
  const [fibResults, setFibResults] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

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

  const fetchFibonacciNumbers = async () => {
    setIsLoading(true);
    setError(null);
    setFibResults([]);

    try {
      let fibArray = [];
      // Fetch each Fibonacci number individually since the API returns one number at a time
      for (let i = 0; i < fibCount; i++) {
        const response = await fetch(`http://localhost:8080/fibonacci?n=${i}`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        fibArray.push(data.value);
      }
      setFibResults(fibArray);
    } catch (error) {
      console.error("Error fetching Fibonacci numbers:", error);
      setError("Error fetching Fibonacci numbers. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h1>Fibonacci Calculator</h1>

        <div style={{ marginBottom: "20px" }}>
          <p>
            Response from server:{" "}
            <span style={{ color: "yellow" }}>{message}</span> on{" "}
            <span style={{ color: "yellow" }}>
              {new Date(serverTimestamp).toLocaleString()}
            </span>
          </p>
        </div>

        <div
          style={{
            margin: "20px 0",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <label htmlFor="fibCount" style={{ marginBottom: "10px" }}>
            Enter number of Fibonacci numbers to calculate:
          </label>
          <div style={{ display: "flex", alignItems: "center" }}>
            <input
              id="fibCount"
              type="number"
              min="1"
              max="93"
              value={fibCount}
              onChange={(e) =>
                setFibCount(
                  Math.min(93, Math.max(1, parseInt(e.target.value) || 0))
                )
              }
              style={{
                padding: "8px",
                fontSize: "16px",
                marginRight: "10px",
                width: "100px",
                textAlign: "center",
                borderRadius: "4px",
                border: "none",
              }}
            />
            <button
              onClick={fetchFibonacciNumbers}
              disabled={isLoading}
              style={{
                padding: "8px 16px",
                fontSize: "16px",
                backgroundColor: "#61dafb",
                color: "#282c34",
                border: "none",
                borderRadius: "4px",
                cursor: isLoading ? "not-allowed" : "pointer",
              }}
            >
              {isLoading ? "Calculating..." : "Calculate"}
            </button>
          </div>
          {error && <p style={{ color: "red" }}>{error}</p>}

          {fibResults.length > 0 && (
            <div style={{ marginTop: "20px", width: "80%", maxWidth: "800px" }}>
              <h3>First {fibResults.length} Fibonacci Numbers:</h3>
              <div
                style={{
                  display: "flex",
                  flexWrap: "wrap",
                  justifyContent: "center",
                  padding: "10px",
                  backgroundColor: "rgba(255, 255, 255, 0.1)",
                  borderRadius: "8px",
                }}
              >
                {fibResults.map((num, index) => (
                  <div
                    key={index}
                    style={{
                      margin: "5px",
                      padding: "8px 12px",
                      backgroundColor: "rgba(97, 218, 251, 0.2)",
                      borderRadius: "4px",
                      minWidth: "40px",
                      textAlign: "center",
                    }}
                  >
                    <div style={{ fontSize: "12px", color: "#aaa" }}>
                      {index}
                    </div>
                    <div>{num.toString()}</div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      </header>
    </div>
  );
}

export default App;
