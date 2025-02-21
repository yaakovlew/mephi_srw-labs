import { useState, useEffect } from 'react';

const Timer = () => {
  const initialTime = 24 * 60 * 60; // 1.5 часа в секундах

  // Попытка извлечь состояние завершения и оставшегося времени из localStorage
  const savedTimeRemaining = parseInt(localStorage.getItem('timeRemaining'), 10) || initialTime;
  const isDone = localStorage.getItem('isDone') === '1';

  const [timeRemaining, setTimeRemaining] = useState(isDone ? 0 : savedTimeRemaining);

  useEffect(() => {
    if (isDone) return; // Если таймер завершен, остановить логику

    const interval = setInterval(() => {
      setTimeRemaining((prev) => {
        const updatedTime = prev - 1;

        if (updatedTime <= 0) {
          clearInterval(interval);
          localStorage.setItem('timeRemaining', '0');
          localStorage.setItem('isDone', '1');
          return 0;
        }

        localStorage.setItem('timeRemaining', updatedTime);
        return updatedTime;
      });
    }, 1000);

    return () => clearInterval(interval);
  }, [isDone]);

  const formatTime = (seconds) => {
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const secs = seconds % 60;

    return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`;
  };

  return (
    <div style={styles.timer}>
      <p>{formatTime(timeRemaining)}</p>
    </div>
  );
};

const styles = {
  timer: {
    position: 'absolute',
    top: '10px',
    left: '10px',
    fontSize: '24px',
    fontWeight: 'bold',
    color: 'white',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    padding: '5px',
    borderRadius: '5px',
  },
};

export default Timer;