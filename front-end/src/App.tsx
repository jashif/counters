import React, { useState, useEffect } from 'react';
import './App.css';
const App = () => {
  const [counters, setCounters] = useState([]);
  const [counterName, setCounterName] = useState('');
  const apiUrl = process.env.REACT_APP_API_URL || 'http://localhost:8080';
  useEffect(() => {
    fetchCounters();
  }, []);

  const fetchCounters = async () => {
    try {
      const response = await fetch(`${apiUrl}/counters`);
      if (response.ok) {
        const data = await response.json();
        setCounters(data);
      }
    } catch (error) {
      console.error('Error fetching counters:', error);
    }
  };

  const createCounter = async () => {
    try {
      await fetch(`${apiUrl}/create`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: counterName }),
      });
      fetchCounters();
    } catch (error) {
      console.error('Error creating counter:', error);
    }
  };

  const incrementCounter = async name => {
    try {
      await fetch(`${apiUrl}/increment?name=${name}`, {
        method: 'GET',
      });
      fetchCounters();
    } catch (error) {
      console.error('Error incrementing counter:', error);
    }
  };

  return (
    <div className="app-container">
      <header>
        <h1>Simple Counter Application</h1>
      </header>

      <div className="main-content">
        <div className="counter-creator">
          <input
            type="text"
            value={counterName}
            onChange={e => setCounterName(e.target.value)}
            placeholder="Counter Name"
          />
          <button onClick={createCounter}>Create Counter</button>
        </div>

        <h2>Counters:</h2>
        <ul>
          {counters.map((counter: any) => (
            <li key={counter.name}>
              {counter.name}: {counter.value}
              <button onClick={() => incrementCounter(counter.name)}>
                Increment
              </button>
            </li>
          ))}
        </ul>
      </div>

      <footer>
        <p>© 2024 Simple Counter App</p>
      </footer>
    </div>
  );
};
export default App;
