resource "aws_iam_policy" "policy" {
  name        = "ingester-allow-es"
  description = "Allow Elastic Search Access"

  policy = templatefile("${path.module}/iam-policy.json.tmpl", {
      arn = var.es_arn
  })
}

resource "aws_iam_role_policy_attachment" "attachment" {
  role       = module.ingester.task_role_name
  policy_arn = aws_iam_policy.policy.arn
}
