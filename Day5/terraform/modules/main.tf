module "create-docker-containers" {
   source = "./provision_docker_containers/"
   container_count = var.container_count
}
