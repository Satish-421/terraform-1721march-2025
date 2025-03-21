terraform {
   required_providers {
      docker = {
         source = "tektutor/docker"
      }
   }
}

resource "docker_container" "container" {
   container_name = "c1"
   host_name = "c1"
   image_name = "tektutor/ubuntu-ansible-node:latest"
}
