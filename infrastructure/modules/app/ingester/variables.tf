variable "region" {
    type = string
}

variable "dns_zone" {
    type = string
}

variable "container_image" {
    type = string
}

variable "es_endpoint" {
    type = string
}

variable "es_arn" {
    type = string
}

variable "cert_arn" {
    type = string
}

variable "container_cpu" {
    type    = number
    default = 256
}

variable "container_memory" {
    type    = number
    default = 512
}
