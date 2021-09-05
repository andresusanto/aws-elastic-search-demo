# Data Producer Script

## Requirements

1. Python 3.8+

Prior to running the script for the first time:

```bash
# Install required dependencies
$ pip install -r requirements.txt


# Add exec permission
$ chmod +x gen-events.py
```

## Usage

```bash
# Help message
$ ./gen-events.py -h
usage: gen-events.py [-h] -n NUM_EVENTS [-u NUM_USERS] [-c CONCURRENCY] [-s SLEEP] [--endpoint ENDPOINT]

optional arguments:
  -h, --help            show this help message and exit
  -n NUM_EVENTS, --num-events NUM_EVENTS
                        number of events that will be generated
  -u NUM_USERS, --num-users NUM_USERS
                        number of unique users that will be used as userId
  -c CONCURRENCY, --concurrency CONCURRENCY
                        number of concurrent requests (default: 5)
  -s SLEEP, --sleep SLEEP
                        sleep duration between each concurrent event in seconds (default: 0.1)
  --endpoint ENDPOINT   event ingester endpoint


# Simulating 10000 events from 200 users, with concurrency of 10 and 3s between each concurrent request:
$ ./gen-events.py -n 10000 -u 200 -c 10 -s 3
```
