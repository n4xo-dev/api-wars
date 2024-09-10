import http from 'k6/http';
import { sleep, group } from 'k6';
import { Rate } from 'k6/metrics';

const BASE_URL = __ENV.BASE_URL || 'http://localhost:3000/graphql';
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

const errorRate = new Rate('graphql_error_rate');

export function consumer() {
  group('consumer', () => {
    // List users
    let query = `
      query {
        users {
          id
        }
      }
    `;
    let response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    let body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const users = body.data.users;
    const randomUserIndex = Math.floor(Math.random() * users.length);
    const randomUserId = users[randomUserIndex].id;
    sleep(ACTION_SLEEP);

    // Get user details and posts
    query = `
      query {
        user(id: "${randomUserId}") {
          id
          name
          email
          posts { id }
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const user = body.data.user;
    sleep(ACTION_SLEEP);

    // Get one post's details and the comments
    const randomPostId = user.posts[Math.floor(Math.random() * user.posts.length)].id;
    query = `
      query {
        post(id: "${randomPostId}") {
          id
          title
          content
          userId
          comments { id }
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const post = body.data.post;
    sleep(ACTION_SLEEP);

    // Get comment's details
    const randomCommentId = post.comments[Math.floor(Math.random() * post.comments.length)].id;
    query = `
      query {
        comment(id: "${randomCommentId}") {
          id
          content
          userId
          postId
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);

    // List chats
    query = `
      query {
        chats {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const chats = body.data.chats;
    const randomChatId = chats[Math.floor(Math.random() * chats.length)].id;
    sleep(ACTION_SLEEP);

    // Get chat participants
    query = `
      query {
        chat(id: "${randomChatId}") {
          participants {
            id
          }
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const participants = body.data.chat.participants;
    const randomParticipantId = participants[Math.floor(Math.random() * participants.length)].id;
    sleep(ACTION_SLEEP);

    // Get messages from chat
    query = `
      query {
        messagesByChat(chatId: "${randomChatId}") {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);

    // Get participants's messages from chat
    query = `
      query {
        messagesByChatAndUser(chatId: "${randomChatId}", userId: "${randomParticipantId}") {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);
  });
}

export function producer() {
  group('producer', () => {
    // Create a user
    let mutation = `
      mutation {
        createUser(input: { name: "New User", email: "${Date.now()}+${(Math.random() * 100).toFixed(0)}@mail.com" }) {
          id
        }
      }
    `;
    let response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    let body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const userId = body.data.createUser.id;
    sleep(ACTION_SLEEP);

    // Create a post
    mutation = `
      mutation {
        createPost(input: { title: "New Post", content: "Post content", userId: "${userId}" }) {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const postId = body.data.createPost.id;
    sleep(ACTION_SLEEP);

    // Create a comment
    mutation = `
      mutation {
        createComment(input: { content: "New Comment", userId: "${userId}", postId: "${postId}" }) {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);

    // Create a chat and add users
    mutation = `
      mutation {
        createChat(input: { participants: ["${userId}"] }) {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    const chatId = body.data.createChat.id;
    sleep(ACTION_SLEEP);

    // Create a message
    mutation = `
      mutation {
        createMessage(input: { content: "New Message", userId: "${userId}", chatId: "${chatId}" }) {
          id
        }
      }
    `;
    response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);
  });
}

export function updater() {
  group('updater', () => {
    const randomId = () => (Math.random() * 100).toFixed(0);

    // Update user, post, comment, and message
    let mutation = `
      mutation {
        updateUser(id: "${randomId()}", input: { email: "updated${Date.now()}@mail.com" }) {
          id
        }
        updatePost(id: "${randomId()}", input: { content: "Updated post content" }) {
          id
        }
        updateComment(id: "${randomId()}", input: { content: "Updated comment content" }) {
          id
        }
        updateMessage(id: "${randomId()}", input: { content: "Updated message content" }) {
          id
        }
      }
    `;
    const response = http.post(BASE_URL, JSON.stringify({ query: mutation }), { headers: { 'Content-Type': 'application/json' } });
    const body = JSON.parse(response.body);
    errorRate.add('errors' in body && body.errors.length > 0);
    if ('errors' in body && body.errors.length > 0) return;
    sleep(ACTION_SLEEP);
  });
}
