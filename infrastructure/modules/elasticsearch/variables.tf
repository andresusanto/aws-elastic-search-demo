variable "instance_type" {
    type = string
}

variable "instance_count" {
    type    = number
    default = 1
}

variable "index_shard_count" {
    type    = number
    default = 1
}
