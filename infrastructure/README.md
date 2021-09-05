# Infrastructure &middot; [![Build](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/apply-infra.yml/badge.svg)](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/apply-infra.yml)

### app/ingester

ECS Fargate deployment. Features:

1. Autoscaling using AWS App Autoscaling
2. Loadbalancer with custom domain name (https) using Route53 and AWS ACM
3. IAM Role and Policy to access AWS ElasticSearch Service
4. SSM Parameter export for integration with Serverless Framework

### dns

Custom domain name management. Features:

1. Route53 hosted-zone management
2. ACM Certificate request
3. Automatic ACM certificate validation by using Route53 DNS records.

### elasticsearch

ElasticSearch instance management. Feature:

1. ElasticSearch instance provisioning.
2. Index Template creation/migration.
3. SSM Parameter export for integration with Serverless Framework.

### monitoring

<img width="480" alt="image" src="https://user-images.githubusercontent.com/7076809/132136181-fa982f5e-ddd5-4dae-8b3d-dd1a648be6f2.png">

AWS CloudWatch management. Feature:

1. CloudWatch log metric filter management.
2. CloudWarch monitoring dashboard.
