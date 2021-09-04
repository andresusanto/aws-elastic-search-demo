import json
import os
import logging
from elasticsearch import Elasticsearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth
import boto3

region = os.environ['REGION']
esHost = os.environ['ES_HOST']

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)
credentials = boto3.Session().get_credentials()
awsauth = AWS4Auth(
    credentials.access_key, 
    credentials.secret_key, 
    region, 
    'es', 
    session_token=credentials.token
)
es = Elasticsearch(
    hosts = [{'host': esHost, 'port': 443}],
    http_auth = awsauth,
    use_ssl = True,
    verify_certs = True,
    connection_class = RequestsHttpConnection
)

def query(event, context):
    logger.info('querying es...')
    res = es.search(index='events', body={"query": {"match_all": {}}})

    logger.info('result')
    logger.info(res)

    return res
