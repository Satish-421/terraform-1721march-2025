variable "image_name" {
   description = "Docker Image Name"
   type        = string
   default     = "tektutor/ubuntu-ansible-node:latest"
}

variable "container_count" {
   description = "Number of containers"
   type = number 
}
