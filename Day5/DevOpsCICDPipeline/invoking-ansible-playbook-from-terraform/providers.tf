terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
      version = "2.0.0"
    }
    file = {
      source = "tektutor/file"
    }
  }
}

provider "docker" {
  # Configuration options
}

provider "file" {
  # Configuration options
}
