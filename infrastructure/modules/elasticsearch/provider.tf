terraform {
  required_providers {
    elasticsearch = {
      source = "phillbaker/elasticsearch"
      version = "2.0.0-beta.1"
    }
  }
}

provider "elasticsearch" {
  url         = "https://${module.es.endpoint}"
  sniff       = false
  healthcheck = false
}
