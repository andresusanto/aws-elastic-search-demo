import datetime
import os
import logging
from elasticsearch import Elasticsearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth
import boto3

region = os.environ['REGION']
esHost = os.environ['ES_HOST']

logger = logging.getLogger()
logger.setLevel(logging.INFO)

credentials = boto3.Session().get_credentials()
awsauth = AWS4Auth(
    credentials.access_key,
    credentials.secret_key,
    region,
    'es',
    session_token=credentials.token
)
es = Elasticsearch(
    hosts=[{'host': esHost, 'port': 443}],
    http_auth=awsauth,
    use_ssl=True,
    verify_certs=True,
    connection_class=RequestsHttpConnection
)


def querier(event, context):
    end = datetime.datetime.utcnow().replace(second=0, microsecond=0)
    start = end - datetime.timedelta(minutes=3)

    res = es.search(
        index='events',
        body={
            "size": 0,
            "query": {
                "range": {
                    "created_at": {
                        "gte": start.isoformat(),
                        "lt": end.isoformat(),
                    }
                }
            },
            "aggs": {
                "num_users": {"cardinality": {"field": "user_id"}},
                "num_events": {"value_count": {"field": "user_id"}}
            }
        }
    )

    num_users = res["aggregations"]["num_users"]["value"]
    num_events = res["aggregations"]["num_events"]["value"]

    logger.info(f'from {start.isoformat()} to {end.isoformat()} we have {num_events} events from {num_users} users')
