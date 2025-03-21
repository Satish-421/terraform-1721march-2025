resource "docker_container" "container" {
  container_name  = var.container_name
  host_name = var.container_name 
  image_name = var.image_name
}

resource "file_localfile" "myfile" {
  file_content = "some ip"
  file_name = "./ip.txt"
}
