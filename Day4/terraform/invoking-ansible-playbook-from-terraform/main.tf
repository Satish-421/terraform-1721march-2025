data "docker_image" "tektutor_image" {
   name = var.image_name 
}

resource "docker_container" "container" {
   name = var.container_name 
   hostname = var.container_name
   image = data.docker_image.tektutor_image.name
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

resource "local_file" "invoke_ansible_playbook" {
   content = "test file"
   filename = "./test.txt"

   provisioner "local-exec" {
      command = "ansible-playbook -i ./hosts install-nginx-playbook.yml"
   }

}
