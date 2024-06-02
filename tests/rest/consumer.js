import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    // { duration: '2s', target: 1 },
    { duration: '15s', target: 10000 },
    { duration: '1m', target: 10000 },
    { duration: '15s', target: 0 },
  ],
};

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  const baseURL = 'http://localhost:8080';
  const ACTION_SLEEP = 60;
  // Get users and pick one
  const users = http.get(`${baseURL}/rest/users/`).json();
  const randomUserIndex = Math.floor(Math.random() * users.length);
  const randomUserId = users[randomUserIndex].id;
  http.get(`${baseURL}/rest/users/${randomUserId}`);
  sleep(ACTION_SLEEP);
  // Get user's posts and pick one
  const posts = http.get(`${baseURL}/rest/users/${randomUserId}/posts`).json();
  const randomPostIndex = Math.floor(Math.random() * posts.length);
  const randomPostId = posts[randomPostIndex].id;
  http.get(`${baseURL}/rest/posts/${randomPostId}`);
  sleep(ACTION_SLEEP);
  // Get post's comments and pick one
  const comments = http.get(`${baseURL}/rest/posts/${randomPostId}/comments`).json();
  const randomCommentIndex = Math.floor(Math.random() * comments.length);
  const randomCommentId = comments[randomCommentIndex].id;
  http.get(`${baseURL}/rest/comments/${randomCommentId}`);
  sleep(ACTION_SLEEP);
  // Get chats and pick one
  const chats = http.get(`${baseURL}/rest/chats`).json();
  const randomChatIndex = Math.floor(Math.random() * chats.length);
  const randomChatId = chats[randomChatIndex].id;
  const chat = http.get(`${baseURL}/rest/chats/${randomChatId}?eager=true`).json();
  sleep(ACTION_SLEEP);
  // Get messages in chat and pick one
  const messages = http.get(`${baseURL}/rest/chats/${randomChatId}/messages`).json();
  const randomMessageIndex = Math.floor(Math.random() * messages.length);
  const randomMessageId = messages[randomMessageIndex].id;
  http.get(`${baseURL}/rest/messages/${randomMessageId}`);
  sleep(ACTION_SLEEP);
  // Get messages from a specific user in chat
  const randomChatUserIndex = Math.floor(Math.random() * chat.participants.length);
  const randomChatUserId = chat.participants[randomChatUserIndex].id;
  http.get(`${baseURL}/rest/chats/${randomChatId}/user/${randomChatUserId}/messages`);
  sleep(ACTION_SLEEP);
}
