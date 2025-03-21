terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
      version = "1.0"
    }
    file = {
      source = "tektutor/file"
      version = "1.0"
    }
  }
}

provider "docker" {
  # Configuration options
}

provider "file" {
  # Configuration options
}
