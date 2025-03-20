terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {
  # Configuration options
}

data "docker_image" "ubuntu" {
  name = "ubuntu:16.04"
}

resource "docker_container" "ubuntu1" {
  name = "ubuntu1_container"
  hostname = "ubuntu1_container"
  image = data.docker_image.ubuntu.name

  must_run = true
  command = [
    "tail",
    "-f",
    "/dev/null"
  ]
}

resource "docker_container" "ubuntu2" {
  name = "ubuntu2_container"
  hostname = "ubuntu2_container"
  image = data.docker_image.ubuntu.name

  must_run = true
  command = [
    "tail",
    "-f",
    "/dev/null"
  ]
}


