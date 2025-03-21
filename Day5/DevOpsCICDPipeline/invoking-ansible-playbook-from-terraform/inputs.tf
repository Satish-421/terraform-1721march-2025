variable "image_name" {
  description = "Name of the docker image"
  type        = string
  default     = "tektutor/ubuntu-ansible-node:latest"
}

variable "container_name" {
  description = "Name of the container"
  type        = string
  default     = "container-1"
}
