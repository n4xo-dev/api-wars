import http from 'k6/http';
import { sleep, group } from 'k6';

const BASE_URL = (__ENV.BASE_URL || 'http://localhost:8080') + '/rest';

export const options = {
  scenarios: {
    cache_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '5s', target: 10000 },
        { duration: '10s', target: 10000 },
      ],
      exec: 'cache',
    },
  },
};

export function cache() {
  group('cache', () => {
    // Ping redis
    http.get(`${BASE_URL}/redis`);
    // Set value
    const key = (Math.random() * 1000000).toFixed(0);
    http.post(`${BASE_URL}/redis`, {
      key: key,
      value: (Math.random() * 1000000000).toString(36),
    });
    // Get value
    http.get(`${BASE_URL}/redis/${key}`);
  });
}
