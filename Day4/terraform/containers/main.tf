data "docker_image" "ubuntu" {
  name = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "ubuntu1" {
  name = var.container_name1 
  hostname = var.container_name1 
  image = data.docker_image.ubuntu.name

  must_run = true

  ports {
	internal = "22"
	external = "2001"
  }
  ports {
	internal = "80"
	external = "8001"
  }
}

resource "docker_container" "ubuntu2" {
  name = var.container_name2 
  hostname = var.container_name2 
  image = data.docker_image.ubuntu.name

  must_run = true
  ports {
	internal = "22"
	external = "2002"
  }
  ports {
	internal = "80"
	external = "8002"
  }
}

resource "docker_container" "ubuntu3" {
  name = var.container_name3 
  hostname = var.container_name3 
  image = data.docker_image.ubuntu.name

  must_run = true
  ports {
	internal = "22"
	external = "2003"
  }
  ports {
	internal = "80"
	external = "8003"
  }
}




