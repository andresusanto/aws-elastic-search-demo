resource "aws_cloudwatch_log_metric_filter" "user_count" {
  name           = "ClickEventUserCounts"
  pattern        = "{ $.message = \"querier metric\" }"
  log_group_name = "/aws/lambda/querier-${var.stage}-fn"

  metric_transformation {
    name      = "NumUsers"
    namespace = "ClickEvents"
    value     = "$.num_users"
    unit      = "Count"
  }
}

resource "aws_cloudwatch_log_metric_filter" "event_count" {
  name           = "ClickEventCounts"
  pattern        = "{ $.message = \"querier metric\" }"
  log_group_name = "/aws/lambda/querier-${var.stage}-fn"

  metric_transformation {
    name      = "NumEvents"
    namespace = "ClickEvents"
    value     = "$.num_events"
    unit      = "Count"
  }
}

resource "aws_cloudwatch_dashboard" "dashboard" {
  dashboard_name = "click-event-insight"
  dashboard_body = templatefile("${path.module}/dashboard.json.tmpl", {
    region = var.region
  })
}
