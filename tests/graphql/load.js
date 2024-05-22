import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    // { duration: '2s', target: 1 },
    { duration: '10s', target: 20 },
    { duration: '50s', target: 20 },
    { duration: '10s', target: 0 },
  ],
};

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  const usersRes = http.post('http://localhost:3000/graphql', JSON.stringify({
    query: `
      query User {
        users {
            id
            name
            email
            createdAt
            updatedAt
            deletedAt
        }
      }  
    `
  }).replace(/\s+/g, ' '), {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  const users = usersRes.json('data.users');
  const randomUserIndex = Math.floor(Math.random() * users.length);
  const randomUserId = users[randomUserIndex].id;
  http.post('http://localhost:3000/graphql/', JSON.stringify({
    query: `
      query User {
        user(id: ${randomUserId}) {
            id
            name
            email
            createdAt
            updatedAt
            deletedAt
        }
      }  
    `
  }).replace(/\s+/g, ' '), {
    headers: {
      'Content-Type': 'application/json',
    },
  });
  sleep(1);
}
