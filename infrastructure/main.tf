terraform {
  backend "s3" {
    key = "tf/aws-es-demo"
  }
  required_version = ">= 0.14.9"
}

module "dns" {
  source = "./modules/dns"
  fqdn   = var.fqdn
}

module "es" {
  source        = "./modules/elasticsearch"
  instance_type = var.es_instance_type
}

module "app_ingester" {
  source = "./modules/app/ingester"

  region          = var.region
  dns_zone        = module.dns.zone_id
  cert_arn        = module.dns.cert_arn
  es_arn          = module.es.arn
  es_endpoint     = module.es.endpoint
  container_image = "${var.ingester_container_name}:${var.ingester_container_tag}"
}

module "monitoring" {
  source = "./modules/monitoring"
  stage  = var.stage
  region = var.region
}