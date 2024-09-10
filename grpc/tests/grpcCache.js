import grpc from 'k6/net/grpc';
import { group } from 'k6';
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

const errorRate = new Rate('grpc_error_rate');
const requests = new Counter('grpc_requests');

export function cache() {
  client.connect('localhost:50051', { plaintext: true });

  group('cache', () => {
    // Ping Redis
    let response = client.invoke('RedisService/Ping', {});
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    // Set a value
    const key = Math.floor(Math.random() * 1000000).toString();
    response = client.invoke('RedisService/Set', {
      record: { key: key, value: (Math.random() * 1000000000).toString(36) },
    });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
    // Get the value
    response = client.invoke('RedisService/Get', { key: key });
    requests.add(1);
    errorRate.add(response.status !== grpc.StatusOK);
  });

  client.close();
}
