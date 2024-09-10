import json
import glob
from collections import defaultdict
import statistics
from datetime import datetime
import re

METRICS_OF_INTEREST = {
    'http_reqs',
    'http_req_duration',
    'http_req_failed',
    'graphql_error_rate',
    'data_sent',
    'data_received',
    'grpc_req_duration',
    'grpc_requests',
    'grpc_error_rate'
}

def parse_time(time_str):
    # Remove the last 3 digits of the nanosecond part
    modified_time_str = re.sub(r'(\.\d{6})\d+', r'\1', time_str)
    return datetime.strptime(modified_time_str, "%Y-%m-%dT%H:%M:%S.%f%z")

def load_json_file(filename):
    metrics = defaultdict(list)
    start_time = None
    end_time = None
    with open(filename, 'r') as file:
        for line in file:
            try:
                data = json.loads(line)
                if data['type'] == 'Point' and data['metric'] in METRICS_OF_INTEREST:
                    metrics[data['metric']].append(data['data']['value'])
                    time = parse_time(data['data']['time'])
                    if start_time is None or time < start_time:
                        start_time = time
                    if end_time is None or time > end_time:
                        end_time = time
            except json.JSONDecodeError:
                print(f"Error parsing line in {filename}: {line}")
            except ValueError as e:
                print(f"Error parsing time in {filename}: {e}")
    return metrics, start_time, end_time

def analyze_metrics(metrics, start_time, end_time):
    analysis = {}
    duration = (end_time - start_time).total_seconds()
    
    for metric, values in metrics.items():
        if metric in ['http_reqs', 'grpc_requests']:
            total_requests = len(values)
            analysis[metric] = {
                'total': total_requests,
                'per_second': total_requests / duration if duration > 0 else 0
            }
        elif metric in ['http_req_failed', 'graphql_error_rate', 'grpc_error_rate']:
            total_errors = sum(values)
            total_requests = len(metrics.get('http_reqs', [])) or len(metrics.get('grpc_requests', []))
            error_percentage = (total_errors / total_requests * 100) if total_requests > 0 else 0
            analysis[metric] = {
                'total': total_errors,
                'percentage': error_percentage
            }
        elif metric in ['data_sent', 'data_received']:
            analysis[metric] = {
                'min': min(values),
                'max': max(values),
                'avg': statistics.mean(values),
                'median': statistics.median(values),
                'total': sum(values)
            }
        else:
            analysis[metric] = {
                'min': min(values),
                'max': max(values),
                'avg': statistics.mean(values),
                'median': statistics.median(values)
            }
    return analysis

def compare_apis(results):
    comparisons = {
        'http_requests': compare_metric(results, 'http_reqs', 'grpc_requests', key='total'),
        'request_duration': compare_metric(results, 'http_req_duration', 'grpc_req_duration', key='avg'),
        'error_rate': compare_metric(results, 'http_req_failed', 'graphql_error_rate', 'grpc_error_rate', key='percentage'),
        'data_sent': compare_metric(results, 'data_sent', key='total'),
        'data_received': compare_metric(results, 'data_received', key='total')
    }
    return comparisons

def compare_metric(results, *metric_names, key='avg'):
    comparison = {}
    for api, data in results.items():
        for metric in metric_names:
            if metric in data:
                comparison[api] = data[metric][key]
                break
    return comparison

def main():
    results = {}
    for filename in glob.glob('*.json'):
        api_name = filename.split('.')[0]
        metrics, start_time, end_time = load_json_file(filename)
        if start_time and end_time:
            results[api_name] = analyze_metrics(metrics, start_time, end_time)
        else:
            print(f"Warning: Could not determine start and end times for {filename}")

    comparisons = compare_apis(results)

    print("API Performance Comparison:")
    for metric, values in comparisons.items():
        print(f"\n{metric.capitalize()}:")
        for api, value in values.items():
            print(f"  {api}: {value}")

    print("\nDetailed Results:")
    for api, metrics in results.items():
        print(f"\n{api}:")
        for metric, values in metrics.items():
            print(f"  {metric}:")
            for key, value in values.items():
                print(f"    {key}: {value}")

if __name__ == "__main__":
    main()