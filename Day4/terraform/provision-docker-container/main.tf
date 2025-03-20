data "docker_image" "ubuntu" {
  name = "tektutor/ubuntu-ansible-node:latest"
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


