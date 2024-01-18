import React, { useState, useEffect } from 'react';

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
    <div>
      <input
        type="text"
        value={counterName}
        onChange={e => setCounterName(e.target.value)}
        placeholder="Counter Name"
      />
      <button onClick={createCounter}>Create Counter</button>
      <h2> Counters:</h2>

      <ul>
        {Object.entries(counters).map(([name, value]) => (
          <li key={name}>
            {name}: {value}
            <button onClick={() => incrementCounter(name)}>Increment</button>
          </li>
        ))}
      </ul>
    </div>
  );
};
export default App;
