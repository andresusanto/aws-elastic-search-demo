import os
import datetime
import boto3
import logging
from pythonjsonlogger import jsonlogger
from elasticsearch import Elasticsearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth


region = os.environ['REGION']
es_host = os.environ['ES_HOST']

logger = logging.getLogger()
logger.setLevel(logging.INFO)
json_handler = logging.StreamHandler()
formatter = jsonlogger.JsonFormatter(
    fmt='%(asctime)s %(levelname)s %(name)s %(message)s')
json_handler.setFormatter(formatter)
logger.addHandler(json_handler)
logger.removeHandler(logger.handlers[0])

credentials = boto3.Session().get_credentials()
awsauth = AWS4Auth(
    credentials.access_key,
    credentials.secret_key,
    region,
    'es',
    session_token=credentials.token) if credentials else None
es = Elasticsearch(
    hosts=[{'host': es_host, 'port': 443}],
    http_auth=awsauth,
    use_ssl=True,
    verify_certs=True,
    connection_class=RequestsHttpConnection)


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

    logger.info('querier metric', extra={
        "start": start.isoformat(),
        "end": end.isoformat(),
        "num_events": num_events,
        "num_users": num_users})
