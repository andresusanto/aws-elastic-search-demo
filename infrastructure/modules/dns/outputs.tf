output "zone_id" {
  value = aws_route53_zone.main.zone_id
}

output "cert_arn" {
  value = aws_acm_certificate.cert.arn
}
