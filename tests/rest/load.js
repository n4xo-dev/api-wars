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
  const usersRes = http.get('http://localhost:8080/rest/users/');  
  const randomUserIndex = Math.floor(Math.random() * usersRes.json().length);
  const randomUserId = usersRes.json()[randomUserIndex].id;
  http.get(`http://localhost:8080/rest/users/${randomUserId}`);
  sleep(1);
}
