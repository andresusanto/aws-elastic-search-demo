resource "aws_ecs_cluster" "cluster" {
  name = "es-ingester-cluster"
}

module "ingester" {
  source = "umotif-public/ecs-fargate/aws"
  version = "~> 6.1.0"

  name_prefix = "es-ingester"
  
  vpc_id             = data.aws_vpc.default.id
  private_subnet_ids = data.aws_subnet_ids.all.ids
  cluster_id         = aws_ecs_cluster.cluster.id

  task_container_assign_public_ip = true
  wait_for_steady_state           = true

  platform_version = "1.4.0"

  task_container_image   = var.container_image
  task_definition_cpu    = var.container_cpu
  task_definition_memory = var.container_memory

  task_container_port         = 8080
  task_container_environment  = {
    PORT            = "8080"
    GIN_MODE        = "release"
    REGION          = var.region
    SIGN_ES_CLIENT  = "true"
    ES_ENDPOINT     = "https://${var.es_endpoint}"
  }

  target_groups = [
    {
      target_group_name = "ingester"
      container_port    = 8080
    }
  ]

  health_check = {
    port = "traffic-port"
    path = "/health"
  }

  task_stop_timeout = 90

  depends_on = [
    module.alb
  ]

  ### To use task credentials, below paramaters are required
  # create_repository_credentials_iam_policy = false
  # repository_credentials                   = aws_secretsmanager_secret.task_credentials.arn
}
