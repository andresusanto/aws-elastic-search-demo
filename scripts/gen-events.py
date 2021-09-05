#!/usr/bin/env python
import time
import argparse
import requests
from uuid import uuid4
from concurrent.futures import ThreadPoolExecutor, as_completed


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-n', '--num-events', type=int, required=True,
                        help='number of events that will be generated')
    parser.add_argument('-u', '--num-users', type=int, default=3,
                        help='number of unique users that will be used as userId')
    parser.add_argument('-c', '--concurrency', type=int, default=5,
                        help='number of concurrent requests (default: 5)')
    parser.add_argument('-s', '--sleep', type=float, default=0.1,
                        help='sleep duration between each concurrent event in seconds (default: 0.1)')
    parser.add_argument('--endpoint', type=str,
                        default='https://ingester.es-demo.susanto.link', help='event ingester endpoint')
    args = parser.parse_args()

    if args.num_events < 1:
        print('Error: number of events must be at least 1')
        exit(1)

    if args.num_users < 1:
        print('Error: number of users must be at least 1')
        exit(1)

    if args.concurrency < 1:
        print('Error: concurrency must be at least 1')
        exit(1)

    print(
        f'\n\nSimulating {args.num_events} events from {args.num_users} users (every {args.sleep}s), with concurency of {args.concurrency}.\n')

    def _req(user_id):
        res = requests.post(f'{args.endpoint}/v1/event', json={
            'user_id': user_id,
            'type': 'CLICK'
        })
        time.sleep(args.sleep)
        return res.status_code

    num_success, num_error, total = 0, 0, 0

    with ThreadPoolExecutor() as executor:
        users = [str(uuid4()) for _ in range(args.num_users)]
        futures = [executor.submit(_req, users[i % len(users)])
                   for i in range(args.num_events)]

        for future in as_completed(futures):
            status_code = future.result()
            if status_code == 201:
                num_success += 1
            else:
                num_error += 1

            total = num_success + num_error
            if total % 10 == 0:
                print(
                    f'Completed {total} requests, with {num_success} success and {num_error} errors.')

    print(
        f'\nDone. In total {total} requests were made. {num_success} were successful and {num_error} were failed.')
