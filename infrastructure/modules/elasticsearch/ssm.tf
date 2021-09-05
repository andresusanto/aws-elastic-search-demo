resource "aws_ssm_parameter" "endpoint" {
  name  = "/es/events/host"
  type  = "String"
  value = module.es.endpoint
}

resource "aws_ssm_parameter" "arn" {
  name  = "/es/events/arn"
  type  = "String"
  value = module.es.arn
}
