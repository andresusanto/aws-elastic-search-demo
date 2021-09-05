resource "aws_ssm_parameter" "endpoint" {
  name  = "/app/ingester/endpoint"
  type  = "String"
  value = "https://${aws_route53_record.ingester.fqdn}"
}