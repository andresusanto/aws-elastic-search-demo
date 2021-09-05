# AWS Elastic Search Demo &middot; [![License](https://img.shields.io/github/license/andresusanto/aws-elastic-search-demo.svg)](https://github.com/andresusanto/aws-elastic-search-demo/blob/main/LICENSE)

Welcome! This project is a demonstration of event ingestion and processing using ElasticSearch on AWS.

[![image](https://user-images.githubusercontent.com/7076809/132135012-f6119ec7-0afa-4858-970f-6b51a323640e.png)](./infrastructure/modules/elasticsearch/README.md)


### Data Ingestion &middot; [![Build](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/build-ingester.yml/badge.svg)](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/build-ingester.yml)

**Ingester** is a microservice that ingest user-click events via HTTP requests into ElasticSearch. It is written in Go and deployed on AWS ECS Fargate.

[More about Ingester](./ingester/README.md)

### Data Querying &middot; [![Build](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/deploy-querier.yml/badge.svg)](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/deploy-querier.yml)

**Querier** is a AWS Lambda function that performs a query to ElasticSearch periodically. It is written in Python and managed using Serverless Framework.

[More about Querier](./querier/README.md)

### Infrastructure &middot; [![Build](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/apply-infra.yml/badge.svg)](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/apply-infra.yml)

All the infrastructure used by this project is managed by Terraform.

[More about infrastructure](./querier/README.md)

### Data Producing

A script in `/scripts` folder is available to generate fake events.

[More about the script](./querier/README.md)

### Index Strategy

Note about the elastic search index strategy can be found [here](./infrastructure/modules/elasticsearch/README.md).
