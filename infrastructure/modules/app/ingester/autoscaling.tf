resource "aws_appautoscaling_target" "ingester" {
  max_capacity       = 4
  min_capacity       = 1
  resource_id        = "service/es-ingester-cluster/${module.ingester.service_name}"
  scalable_dimension = "ecs:service:DesiredCount"
  service_namespace  = "ecs"
}

resource "aws_appautoscaling_policy" "ecs_policy" {
  name               = "scale-down"
  policy_type        = "StepScaling"
  resource_id        = aws_appautoscaling_target.ingester.resource_id
  scalable_dimension = aws_appautoscaling_target.ingester.scalable_dimension
  service_namespace  = aws_appautoscaling_target.ingester.service_namespace

  step_scaling_policy_configuration {
    adjustment_type         = "ChangeInCapacity"
    cooldown                = 60
    metric_aggregation_type = "Maximum"

    step_adjustment {
      metric_interval_upper_bound = 0
      scaling_adjustment          = -1
    }
  }
}
