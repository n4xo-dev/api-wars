import json
import glob
from collections import defaultdict
import statistics
from datetime import datetime, timedelta
import re
import matplotlib.pyplot as plt
import numpy as np
from matplotlib.dates import DateFormatter

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
    modified_time_str = re.sub(r'(\.\d{6})\d+', r'\1', time_str)
    return datetime.strptime(modified_time_str, "%Y-%m-%dT%H:%M:%S.%f%z")

def load_json_file(filename):
    metrics = defaultdict(list)
    timestamps = []
    start_time = None
    end_time = None
    with open(filename, 'r') as file:
        for line in file:
            try:
                data = json.loads(line)
                if data['type'] == 'Point' and data['metric'] in METRICS_OF_INTEREST:
                    time = parse_time(data['data']['time'])
                    timestamps.append(time)
                    metrics[data['metric']].append((time, data['data']['value']))
                    if start_time is None or time < start_time:
                        start_time = time
                    if end_time is None or time > end_time:
                        end_time = time
            except json.JSONDecodeError:
                print(f"Error parsing line in {filename}: {line}")
            except ValueError as e:
                print(f"Error parsing time in {filename}: {e}")
    return metrics, start_time, end_time, timestamps

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
            total_errors = sum(value for _, value in values)
            total_requests = len(metrics.get('http_reqs', [])) or len(metrics.get('grpc_requests', []))
            error_percentage = (total_errors / total_requests * 100) if total_requests > 0 else 0
            analysis[metric] = {
                'total': total_errors,
                'percentage': error_percentage
            }
        elif metric in ['data_sent', 'data_received']:
            values_only = [value for _, value in values]
            analysis[metric] = {
                'min': min(values_only),
                'max': max(values_only),
                'avg': statistics.mean(values_only),
                'median': statistics.median(values_only),
                'total': sum(values_only)
            }
        else:
            values_only = [value for _, value in values]
            analysis[metric] = {
                'min': min(values_only),
                'max': max(values_only),
                'avg': statistics.mean(values_only),
                'median': statistics.median(values_only)
            }
    
    # Calculate average request size
    total_data_sent = analysis.get('data_sent', {}).get('total', 0)
    total_requests = analysis.get('http_reqs', {}).get('total') or analysis.get('grpc_requests', {}).get('total', 0)
    if total_requests > 0:
        analysis['avg_request_size'] = total_data_sent / total_requests
    else:
        analysis['avg_request_size'] = 0

    return analysis

def compare_apis(results):
    comparisons = {
        'http_requests': compare_metric(results, 'http_reqs', 'grpc_requests', key='total'),
        'request_duration': compare_metric(results, 'http_req_duration', 'grpc_req_duration', key='avg'),
        'error_rate': compare_error_rates(results),
        'data_sent': compare_metric(results, 'data_sent', key='total'),
        'data_received': compare_metric(results, 'data_received', key='total'),
        'avg_request_size': compare_metric(results, 'avg_request_size')
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

def compare_error_rates(results):
    comparison = {}
    for api, data in results.items():
        if 'graphql_error_rate' in data:
            comparison[api] = data['graphql_error_rate']['percentage']
        elif 'grpc_error_rate' in data:
            comparison[api] = data['grpc_error_rate']['percentage']
        elif 'http_req_failed' in data:
            comparison[api] = data['http_req_failed']['percentage']
    return comparison

def plot_bar_chart(data, title, ylabel, filename):
    plt.figure(figsize=(10, 6))
    plt.bar(data.keys(), data.values())
    plt.title(title)
    plt.ylabel(ylabel)
    plt.xlabel('API')
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()

def plot_stacked_bar_chart(data, title, ylabel, filename):
    apis = list(data.keys())
    metrics = list(data[apis[0]].keys())
    
    fig, ax = plt.subplots(figsize=(12, 6))
    bottom = np.zeros(len(apis))
    
    for metric in metrics:
        values = [data[api][metric] for api in apis]
        ax.bar(apis, values, label=metric, bottom=bottom)
        bottom += np.array(values)
    
    ax.set_title(title)
    ax.set_ylabel(ylabel)
    ax.set_xlabel('API')
    ax.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()

def plot_detailed_time_series(api_name, rps_data, latency_data, errors_per_second_data, data_sent, data_received, filename):
    fig, (ax1, ax2, ax3, ax4, ax5) = plt.subplots(5, 1, figsize=(12, 25), sharex=True)
    fig.suptitle(f'Time Series Analysis for {api_name}', fontsize=16)

    # Requests per Second
    times_rps, values_rps = zip(*rps_data)
    ax1.plot(times_rps, values_rps, label='RPS', color='blue')
    ax1.set_ylabel('Requests per Second')
    ax1.legend()
    ax1.grid(True)

    # Latency
    times_latency, values_latency = zip(*latency_data)
    ax2.plot(times_latency, values_latency, label='Latency', color='green')
    ax2.set_ylabel('Latency (ms)')
    ax2.legend()
    ax2.grid(True)

    # Errors per Second
    if errors_per_second_data:
        times_errors, values_errors = zip(*errors_per_second_data)
        ax3.plot(times_errors, values_errors, label='Errors per Second', color='red')
    else:
        ax3.text(0.5, 0.5, 'No errors recorded', horizontalalignment='center', verticalalignment='center', transform=ax3.transAxes)
    ax3.set_ylabel('Errors per Second')
    ax3.legend()
    ax3.grid(True)

    # Data Sent
    times_sent, values_sent = zip(*data_sent)
    ax4.plot(times_sent, values_sent, label='Data Sent', color='purple')
    ax4.set_ylabel('Data Sent (bytes)')
    ax4.legend()
    ax4.grid(True)

    # Data Received
    times_received, values_received = zip(*data_received)
    ax5.plot(times_received, values_received, label='Data Received', color='orange')
    ax5.set_ylabel('Data Received (bytes)')
    ax5.legend()
    ax5.grid(True)


    for ax in (ax1, ax2, ax3, ax4, ax5):
        ax.xaxis.set_major_formatter(DateFormatter('%H:%M:%S'))
        
    plt.xlabel('Time')
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()

def calculate_data_transfer_per_second(data_points):
    transfer = defaultdict(int)
    for time, value in data_points:
        transfer[time.replace(microsecond=0)] += value
    return sorted((time, count) for time, count in transfer.items())

def calculate_requests_per_second(requests):
    rps = defaultdict(int)
    for time, _ in requests:
        rps[time.replace(microsecond=0)] += 1
    return sorted((time, count) for time, count in rps.items())

def calculate_errors_per_second(errors):
    eps = defaultdict(int)
    for time, value in errors:
        if value == 1:  # Count only when the error metric is 1
            eps[time.replace(microsecond=0)] += 1
    return sorted((time, count) for time, count in eps.items())

def main():
    results = {}
    time_series_data = defaultdict(lambda: defaultdict(dict))
    
    for filename in glob.glob('*.json'):
        api_name = filename.split('.')[0]
        metrics, start_time, end_time, timestamps = load_json_file(filename)
        
        if start_time and end_time:
            results[api_name] = analyze_metrics(metrics, start_time, end_time)
            

            if 'http_reqs' in metrics:
                time_series_data[api_name]['requests_per_second'] = calculate_requests_per_second(metrics['http_reqs'])
                if 'graphql' in api_name.lower():

                    time_series_data[api_name]['errors_per_second'] = calculate_errors_per_second(metrics.get('graphql_error_rate', []))
                else:

                    time_series_data[api_name]['errors_per_second'] = calculate_errors_per_second(metrics.get('http_req_failed', []))
                time_series_data[api_name]['latency'] = metrics['http_req_duration']
            elif 'grpc_requests' in metrics:
                time_series_data[api_name]['requests_per_second'] = calculate_requests_per_second(metrics['grpc_requests'])
                time_series_data[api_name]['errors_per_second'] = calculate_errors_per_second(metrics.get('grpc_error_rate', []))
                time_series_data[api_name]['latency'] = metrics['grpc_req_duration']
            

            time_series_data[api_name]['data_sent'] = calculate_data_transfer_per_second(metrics['data_sent'])
            time_series_data[api_name]['data_received'] = calculate_data_transfer_per_second(metrics['data_received'])
        else:
            print(f"Warning: Could not determine start and end times for {filename}")

    comparisons = compare_apis(results)

    # Generate bar charts
    plot_bar_chart(comparisons['http_requests'], 'Total Requests by API', 'Number of Requests', 'total_requests.png')
    plot_bar_chart(comparisons['request_duration'], 'Average Request Duration by API', 'Duration (ms)', 'avg_duration.png')
    plot_bar_chart(comparisons['error_rate'], 'Error Rate by API', 'Error Rate (%)', 'error_rate.png')
    plot_bar_chart(comparisons['data_sent'], 'Total Data Sent by API', 'Data Sent (bytes)', 'data_sent.png')
    plot_bar_chart(comparisons['data_received'], 'Total Data Received by API', 'Data Received (bytes)', 'data_received.png')
    plot_bar_chart(comparisons['avg_request_size'], 'Average Request Size by API', 'Size (bytes)', 'avg_request_size.png')

    data_transfer = {api: {'Sent': results[api]['data_sent']['total'], 'Received': results[api]['data_received']['total']} for api in results}
    plot_stacked_bar_chart(data_transfer, 'Data Transfer by API', 'Data Transfer (bytes)', 'data_transfer.png')

    for api_name, data in time_series_data.items():
        plot_detailed_time_series(
            api_name,
            data['requests_per_second'],
            data['latency'],
            data['errors_per_second'],
            data['data_sent'],
            data['data_received'],
            f'{api_name}_time_series.png'
        )


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

    print("\nGraphs have been saved as PNG files in the current directory.")

if __name__ == "__main__":
    main()