module "es" {
  source = "lgallard/elasticsearch/aws"

  domain_name           = "es-events"
  elasticsearch_version = "7.1"

  cluster_config = {
    dedicated_master_enabled = false
    instance_count           = var.instance_count
    instance_type            = var.instance_type
  }

  ebs_options = {
    ebs_enabled = "true"
    volume_size = "25"
  }

  node_to_node_encryption_enabled                = true
  snapshot_options_automated_snapshot_start_hour = 23
}
