import http from 'k6/http';
import { sleep, group } from 'k6';

const BASE_URL = (__ENV.BASE_URL || 'http://localhost:8080') + '/rest';
const ACTION_SLEEP = parseInt(__ENV.ACTION_SLEEP) || 1;

export const options = {
  scenarios: {
    consumer_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 1000 },
      ],
      exec: 'consumer',
      startTime: '0s',
    },
    producer_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 1000 },
      ],
      exec: 'producer',
      startTime: '2m',
    },
    big_consumer_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 1000 },
      ],
      exec: 'consumer',
      startTime: '4m',
    },
    updater_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 1000 },
      ],
      exec: 'updater',
      startTime: '6m',
    },
  },
};

export function consumer() {
  group('consumer', () => {
    // Get users and pick one
    const users = http.get(`${BASE_URL}/users`).json();
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
  });
}

export function producer() {
  group('producer', () => {
    // Create a user
    const user = http.post(`${BASE_URL}/users/`, {
      name: `New User ${Date.now()}`,
      email: `${Date.now()}+${(Math.random() * 100).toFixed(0)}@mail.com`,
    }).json();
    sleep(ACTION_SLEEP);
    // Create a post with that user
    const post = http.post(`${BASE_URL}/posts`, {
      title: `Post title ${Date.now()}`,
      content: 'Post content',
      userId: user.id,
    }).json();
    sleep(ACTION_SLEEP);
    // Write a comment on that post
    http.post(`${BASE_URL}/comments`, {
      content: 'Comment content',
      userId: user.id,
      postId: post.id,
    });
    sleep(ACTION_SLEEP);
    // Creat a chat and add user
    const chat = http.post(`${BASE_URL}/chats`, null).json();
    http.post(`${BASE_URL}/chats/${chat.id}/users`, JSON.stringify([user.id]))
    sleep(ACTION_SLEEP);
    // Write a message on the chat with the user
    http.post(`${BASE_URL}/messages`, {
      content: 'Message content',
      userId: user.id,
      chatId: chat.id,
    }).json();
    sleep(ACTION_SLEEP);
  });
}

export function updater() {
  group('updater', () => {
    const randomId = () => (Math.random() * 100 + 1).toFixed(0);
    // Update a user
    http.patch(`${BASE_URL}/users/${randomId()}`, {
      email: `updated${Date.now()}@mail.com`,
    }).json();
    sleep(ACTION_SLEEP);
    // Update a post
    http.patch(`${BASE_URL}/posts/${randomId()}`, {
      content: 'Post content updated',
    }).json();
    sleep(ACTION_SLEEP);
    // Update a comment
    http.patch(`${BASE_URL}/comments/${randomId()}`, {
      content: 'Comment content updated',
    });
    sleep(ACTION_SLEEP);
    // Update a message
    http.patch(`${BASE_URL}/messages/${randomId()}`, {
      content: 'Message content updated',
    }).json();
    sleep(ACTION_SLEEP);
  });
}