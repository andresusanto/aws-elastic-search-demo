resource "aws_route53_record" "ingester" {
  zone_id = var.dns_zone
  name    = "ingester"
  type    = "CNAME"
  ttl     = "60"
  records = [module.alb.dns_name]
}
