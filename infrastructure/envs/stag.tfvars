stage             = "stag"
fqdn              = "es-demo-staging.susanto.link"
region            = "ap-southeast-2"
es_instance_type  = "t3.small.elasticsearch"
create_monitoring = false

ingester_container_name = "ghcr.io/andresusanto/es-event-ingester"
ingester_container_tag  = "sha-fc0ea56"
