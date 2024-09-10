import grpc from 'k6/net/grpc';
import { sleep, group } from 'k6';
import { Counter, Rate } from 'k6/metrics';

const client = new grpc.Client();
client.load(
  ['../protos'],
  'chats.proto',
  'comments.proto',
  'messages.proto',
  'posts.proto',
  'redis_record.proto',
  'users.proto'
);

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

const errorRate = new Rate('grpc_error_rate');
const requests = new Counter('grpc_requests');

export function consumer() {
  client.connect('localhost:50051', { plaintext: true });

  group('consumer', () => {
    // List users
    let response = client.invoke('UserService/ListUsers', {});
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const users = response.message.users;
    const randomUserIndex = Math.floor(Math.random() * users.length);
    const randomUserId = users[randomUserIndex].id;

    // Get user
    response = client.invoke('UserService/GetUser', { id: randomUserId });
    requests.add(1);
    sleep(ACTION_SLEEP);

    // List posts for the user
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    response = client.invoke('PostsService/ListPosts', { user_id: randomUserId });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const posts = response.message.posts;
    const randomPostIndex = Math.floor(Math.random() * posts.length);
    const randomPostId = posts[randomPostIndex].id;

    // Get post
    response = client.invoke('PostsService/GetPost', { id: randomPostId });
    requests.add(1);
    sleep(ACTION_SLEEP);

    // List comments for the post
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    response = client.invoke('PostsService/GetPostsComments', { post_id: randomPostId });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const comments = response.message.comments;
    const randomCommentIndex = Math.floor(Math.random() * comments.length);
    const randomCommentId = comments[randomCommentIndex].id;

    // Get comment
    response = client.invoke('CommentsService/GetComment', { id: randomCommentId });
    requests.add(1);
    sleep(ACTION_SLEEP);

    // List chats
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    response = client.invoke('ChatsService/ListChats', { eager: true });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const chats = response.message.chats;
    const randomChatIndex = Math.floor(Math.random() * chats.length);
    const randomChatId = chats[randomChatIndex].id;

    // Get chat
    response = client.invoke('ChatsService/GetChat', { id: randomChatId, eager: true });
    requests.add(1);
    sleep(ACTION_SLEEP);

    // List chat messages
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    response = client.invoke('ChatsService/GetChatMessages', { id: randomChatId });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const messages = response.message.messages;
    const randomMessageIndex = Math.floor(Math.random() * messages.length);
    const randomMessageId = messages[randomMessageIndex].id;

    // Get message
    response = client.invoke('MessagesService/GetMessage', { id: randomMessageId });
    requests.add(1);
    sleep(ACTION_SLEEP);

    // Get messages from a specific user in the chat
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const randomChatUserId = chats[randomChatIndex].participants[0].id;
    response = client.invoke('ChatsService/GetChatUserMessages', {
      chat_id: randomChatId,
      user_id: randomChatUserId,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);
  });

  client.close();
}

export function producer() {
  client.connect('localhost:50051', { plaintext: true });

  group('producer', () => {
    // Create a user
    let response = client.invoke('UserService/CreateUser', {
      name: `New User ${Date.now()}`,
      email: `${Date.now()}+${(Math.random() * 1000000).toFixed(0)}@mail.com`,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const user = response.message.user;
    sleep(ACTION_SLEEP);

    // Create a post
    response = client.invoke('PostsService/CreatePost', {
      title: `Post title ${Date.now()}`,
      content: 'Post content',
      user_id: user.id,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const post = response.message.post;
    sleep(ACTION_SLEEP);

    // Create a comment
    response = client.invoke('CommentsService/CreateComment', {
      content: 'Comment content',
      user_id: user.id,
      post_id: post.id,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);

    // Create a chat and add the user
    response = client.invoke('ChatsService/CreateChat', {});
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    const chat = response.message.chat;
    requests.add(1);
    response = client.invoke('ChatsService/AddUsersToChat', {
      chat_id: chat.id,
      user_ids: [user.id],
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);

    // Send a message
    response = client.invoke('MessagesService/CreateMessage', {
      content: 'Message content',
      user_id: user.id,
      chat_id: chat.id,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);
  });

  client.close();
}

export function updater() {
  client.connect('localhost:50051', { plaintext: true });

  group('updater', () => {
    const randomId = () => Math.floor(Math.random() * 100 + 1).toFixed(0);

    // Update user
    let response = client.invoke('UserService/UpdateUser', {
      id: randomId(),
      email: `updated${Date.now()}@mail.com`,
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);

    // Update post
    response = client.invoke('PostsService/UpdatePost', {
      id: randomId(),
      content: 'Post content updated',
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);

    // Update comment
    response = client.invoke('CommentsService/UpdateComment', {
      id: randomId(),
      content: 'Comment content updated',
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);

    // Update message
    response = client.invoke('MessagesService/UpdateMessage', {
      id: randomId(),
      content: 'Message content updated',
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    if (response.status !== grpc.StatusOK) return;
    sleep(ACTION_SLEEP);
  });

  client.close();
}
