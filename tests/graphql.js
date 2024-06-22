import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  scenarios: {
    consumer_scenario: {
      executor: 'ramping-vus',
      stages: [
        { duration: '15s', target: 1000 },
        { duration: '1m30s', target: 1000 },
        { duration: '15s', target: 0 },
      ],
      exec: 'consumer',
    }
  }
};

export function consumer() {
  const baseURL = 'http://localhost:3000/graphql';
  const ACTION_SLEEP = 20;
  // Get users and pick one 
  const headers = { 'Content-Type': 'application/json' };
  const users = http.post(baseURL, JSON.stringify({
    query: `query {
      users {
        id
        name
        email
        createdAt
        updatedAt
      }
    }`
  }), { headers }).json().data.users;
  const randomUserIndex = Math.floor(Math.random() * users.length);
  const randomUserId = users[randomUserIndex].id;
  http.post(baseURL, JSON.stringify({
    query: `query {
      user(id: ${randomUserId}) {
        id
        name
        email
        createdAt
        updatedAt
      }
    }`
  }), { headers });
  sleep(ACTION_SLEEP);
  // Get user's posts and pick one
  const posts = http.post(baseURL, JSON.stringify({
    query: `query {
      posts {
          id
          title
          content
          userId
          createdAt
          updatedAt
      }
    }`
  }), { headers }).json().data.posts;
  const randomPostIndex = Math.floor(Math.random() * posts.length);
  const randomPostId = posts[randomPostIndex].id;
  http.post(baseURL, JSON.stringify({
    query: `query {
      post(id: ${randomPostId}) {
          id
          title
          content
          userId
          createdAt
          updatedAt
      }
    }`
  }), { headers });
  sleep(ACTION_SLEEP);
  // Get post's comments and pick one
  const comments = http.post(baseURL, JSON.stringify({
    query: `query {
      comments {
          id
          content
          userId
          postId
          createdAt
          updatedAt
      }
    }`
  }), { headers }).json().data.comments;
  const randomCommentIndex = Math.floor(Math.random() * comments.length);
  const randomCommentId = comments[randomCommentIndex].id;
  http.post(baseURL, JSON.stringify({
    query: `query {
      comment(id: ${randomCommentId}) {
          id
          content
          userId
          postId
          createdAt
          updatedAt
      }
    }`
  }), { headers });
  sleep(ACTION_SLEEP);
  // Get chats and pick one
  const chats = http.post(baseURL, JSON.stringify({
    query: `query {
      chats {
          id
          createdAt
          updatedAt
      }
    }`
  }), { headers }).json().data.chats;
  const randomChatIndex = Math.floor(Math.random() * chats.length);
  const randomChatId = chats[randomChatIndex].id;
  const chat = http.post(baseURL, JSON.stringify({
    query: `query {
      chat(id: ${randomChatId}) {
          id
          createdAt
          updatedAt
          deletedAt
          messages {
              id
              content
              userId
              chatId
              createdAt
              updatedAt
          }
          participants {
              id
              name
              email
              createdAt
              updatedAt
          }
      }
    }`
  }), { headers }).json().data.chat;
  sleep(ACTION_SLEEP);
  // Get messages in chat and pick one
  const messages = http.post(baseURL, JSON.stringify({
    query: `query {
      chat(id: ${randomChatId}) {
          messages {
            id
            content
            userId
            chatId
            createdAt
            updatedAt
        }
      }
    }`
  }), { headers }).json().data.chat.messages;
  const randomMessageIndex = Math.floor(Math.random() * messages.length);
  const randomMessageId = messages[randomMessageIndex].id;
  http.post(baseURL, JSON.stringify({
    query: `query {
      message(id: ${randomMessageId}) {
          id
          content
          userId
          chatId
          createdAt
          updatedAt
      }
    }`
  }), { headers });
  sleep(ACTION_SLEEP);
  // Get messages from a specific user in chat
  const randomChatUserIndex = Math.floor(Math.random() * chat.participants.length);
  const randomChatUserId = chat.participants[randomChatUserIndex].id;
  http.post(baseURL, JSON.stringify({
    query: `query {
      messagesByChatAndUser(chatId: ${randomChatId}, userId: ${randomChatUserId}) {
          id
          content
          userId
          chatId
          createdAt
          updatedAt
      }
    }`
  }), { headers });
  sleep(ACTION_SLEEP);
}
