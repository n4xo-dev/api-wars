import http from 'k6/http';
import { group } from 'k6';
import { Rate } from 'k6/metrics';

const BASE_URL = __ENV.BASE_URL || 'http://localhost:3000/graphql';

export const options = {
  scenarios: {
    cache_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '1m', target: 10000 },
        { duration: '30s', target: 10000 },
      ],
      exec: 'cache',
    },
  },
};

const errorRate = new Rate('graphql_error_rate');

export function cache() {
  group('cache', () => {
    const key = (Math.random() * 1000000).toFixed(0);
    // Set a Redis record
    let mutation = `
      mutation {
        createRedisRecord(input: { key: "${key}", value: "${Math.random().toString(36)}" }) {
          key
        }
      }
    `;
    let response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    let body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);

    // Get the Redis record
    let query = `
      query {
        redisRecord(key: "${key}") {
          value
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query: query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
  });
}
