import http from 'k6/http';
import { sleep } from 'k6';

const BASE_URL = 'http://localhost:8080/rest';

export const options = {
  stages: [
    // { duration: '2s', target: 1 },
    { duration: '15s', target: 1000 },
    { duration: '1m30s', target: 1000 },
    { duration: '15s', target: 0 },
  ],
};

export function test() {
  console.log('@@@ >>> TEST <<< @@@');
  http.get(`${BASE_URL}/user/1`);
}

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  const ACTION_SLEEP = 20;
  // Get users and pick one
  const users = http.get(`${BASE_URL}/users/`).json();
  const randomUserIndex = Math.floor(Math.random() * users.length);
  const randomUserId = users[randomUserIndex].id;
  http.get(`${BASE_URL}/users/${randomUserId}`);
  sleep(ACTION_SLEEP);
  // Get user's posts and pick one
  const posts = http.get(`${BASE_URL}/users/${randomUserId}/posts`).json();
  const randomPostIndex = Math.floor(Math.random() * posts.length);
  const randomPostId = posts[randomPostIndex].id;
  http.get(`${BASE_URL}/posts/${randomPostId}`);
  sleep(ACTION_SLEEP);
  // Get post's comments and pick one
  const comments = http.get(`${BASE_URL}/posts/${randomPostId}/comments`).json();
  const randomCommentIndex = Math.floor(Math.random() * comments.length);
  const randomCommentId = comments[randomCommentIndex].id;
  http.get(`${BASE_URL}/comments/${randomCommentId}`);
  sleep(ACTION_SLEEP);
  // Get chats and pick one
  const chats = http.get(`${BASE_URL}/chats`).json();
  const randomChatIndex = Math.floor(Math.random() * chats.length);
  const randomChatId = chats[randomChatIndex].id;
  const chat = http.get(`${BASE_URL}/chats/${randomChatId}?eager=true`).json();
  sleep(ACTION_SLEEP);
  // Get messages in chat and pick one
  const messages = http.get(`${BASE_URL}/chats/${randomChatId}/messages`).json();
  const randomMessageIndex = Math.floor(Math.random() * messages.length);
  const randomMessageId = messages[randomMessageIndex].id;
  http.get(`${BASE_URL}/messages/${randomMessageId}`);
  sleep(ACTION_SLEEP);
  // Get messages from a specific user in chat
  const randomChatUserIndex = Math.floor(Math.random() * chat.participants.length);
  const randomChatUserId = chat.participants[randomChatUserIndex].id;
  http.get(`${BASE_URL}/chats/${randomChatId}/user/${randomChatUserId}/messages`);
  sleep(ACTION_SLEEP);
}
