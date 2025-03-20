# Day 4

## Lab - Golang modules

Let's create 3 modules in your home directory
```
cd ~
mkdir addition subtraction main
```

Let's create the first golang module called addition
```
cd ~/addition
go mod init addition
cat go.mod
touch addition.go
```

Update the addition.go file with the below content
<pre>
package addition

func Add( x float32, y float32 ) float64 {
  return float64( x + y )
}
</pre>

Let's create the second golang module called subtraction
```
cd ~/subtraction
go mod init subtraction
cat go.mod
touch subtraction.go
```

Update the subtraction.go file with the below content
<pre>
package subtraction

func Subtract( x float32, y float32 ) float64 {
  return float64(x-y)
}
</pre>

Let's create the main module
```
cd ~/main
go mod init main
cat go.mod
touch main.go
```

Update the main.go with the below content
<pre>
package main

import (
   "fmt"
   "addition"
   "subtraction"
)
  
func main( ) {
   x := float32(500.7)
   y := float32(200.5)
   result1 := addition.Add( x, y )
   result2 := subtraction.Subtract( x, y )

   fmt.Println ( "The sum of ", x, " and ", y, " is ", result1 )
   fmt.Println ( "The difference of ", x, " and ", y, " is ", result2 )
}
</pre>

Running the above application, the go mod edit command below helps golang locate the addition and subtraction modules. Otherwise go mod tidy will attempt to download the package from internet and it fail as we haven't published our module.
```
cd ~/main
go mod edit --replace addition=../addition
go mod edit --replace subtraction=../subtraction
go mod tidy
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/783e23c4-cc62-4ef4-9238-0f416ac56206)

## Lab - Golang concurrency to run functions in different threads

Create a file named main.go with the below content
```
package main

import (
  "fmt"
  "time"
)

func firstFunction( count int ) {
   for i := 0; i<count; i++ {
	fmt.Println ("First function", i )
	time.Sleep( time.Millisecond * 5 )
   }
}

func secondFunction( ) {
   for i := 0; i<1000; i++ {
	fmt.Println ("Second function", i )
	time.Sleep( time.Millisecond * 5 )
   }
}

func main() {
  fmt.Println ( "Press any key to exit ..." )

  go firstFunction( 1000 )
  go secondFunction( )

  var tmp string
  fmt.Scanln(&tmp )
}
```

Run it
```
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/549a9a15-6063-41cb-b63e-765dc3352db6)
![image](https://github.com/user-attachments/assets/53d747ba-328f-42d8-b871-f5f0f73fe90f)
![image](https://github.com/user-attachments/assets/298e1eab-77b3-4b87-999e-d993baf3b4cb)

## Info - Terraform
<pre>
- is a Infrastructure as a code tool
- it is provisioning tool that helps automating all provisioning activities via decalarative code
- it is helpful to provision resources in local environment, private cloud, public cloud, hybrid cloud, etc.,
- it can help automate managing containers/images locally using docker
- it can help provisioning virtual machines using virtualbox, vmware locally
- it can help provisioning ec2 instances, storage in AWS public cloud
- it can help provision virtual machine, storage in azure public cloud
- it is a provisioning tool
- Alternates to Terraform
  - AWS Cloudformation
- is developed by a company called Hashicorp in Go lang
- Terraform uses a DSL called HCL(Hashicorp Configuration Language - for automation )
- is a cloud-newtral provisioning tool, it works locally as well on all the clouds
- it comes in two flavours
  1. Terraform Core - opensource and supports only command-line interface
  2. Terraform Enterprise - requires license, supports web interface
</pre>

## Info - Terraform vs Ansible
<pre>
- Creating a VM on the local environment, public/private/hybrid cloud with some Operating System can be done via Terraform
- To install softwares on those VMs, we can Ansible or any configuration management tool
- Terraform can also invoke ansible playbooks once the provisioning is completed by Terraform
- Terraform can invoke any configuration management tool after provisioning, for example it can invoke chef/puppet/saltstack/ansible	
- In general Terraform and Ansible are used in combination
- Terraform 
  - its strength is Provisioning
  - mainly used for Provisioning Infrastructure i.e Creating a Virtual Machine and installing an Operating into the Virtual Machine
  - also supports basic configuration management using shell scripts/powershell which is imperative, hence this approach is not recommended for complex configuration management
- Ansible
  - its strength is Configuration Management
  - mainly used to install/uninstall/upgrade softwares on an existing virtual machine/base machine with some OS pre-installed in it
  - also supports basic provisioning
</pre>

## Info - Terraform High Level Architecture
![terraform](terraform-architecture-diagram.png)

## Info - Terraform Providers
<pre>
- Terraform depends on providers to provision resources
- Terraform providers are implemented in Golang using Terraform Provider plugin framework
- Terraform supports two types of resources
  - Resources
    - these are resources managed by Terraform
  - Datasources 
    - these are existing resources that were either created manually or without using Terraform
    - these resources will not be deleted by Terraform
- For instance
  - to provision resources in AWS, there is an aws provider that supports many aws resources and datasources
</pre>

## Lab - Checking Terraform version
```
terraform version
```

Expected output
![image](https://github.com/user-attachments/assets/8a6d5153-42d6-45ae-885f-e02d8dac6e25)


## Demo - Writing your first Terraform manifest file ( automation script )
We will be using docker provider from registry.terraform.io website
![image](https://github.com/user-attachments/assets/08938fce-41f7-47d5-a5ac-198557781ca4)

```
cd ~/terraform-1721march-2025
git pull
cd Day4/terraform/pull-docker-image
cat main.tf
terraform init
tree .terraform
terraform plan
terraform apply 
docker images
terraform destroy
```

Expected output
![image](https://github.com/user-attachments/assets/24002e2b-3731-42c2-9608-c655a182c243)
![image](https://github.com/user-attachments/assets/1b699998-5658-4f0f-ba29-c716ebc230f2)
![image](https://github.com/user-attachments/assets/709da4de-9fa8-45b4-8c75-157d8c69055f)
![image](https://github.com/user-attachments/assets/dff79bf1-cddc-4da4-a4dc-27b240271dff)
![image](https://github.com/user-attachments/assets/eec98ed7-fee6-4cc3-9b44-6d1f2fe2c02a)
![image](https://github.com/user-attachments/assets/c53ee31a-f3d8-40c0-b018-41c414aaf0c7)
![image](https://github.com/user-attachments/assets/0fd3f82e-620c-4a05-bf81-6af0c59b9a96)

## Lab - Let's use the ubuntu:latest image that is already there in our lab machine
```
terraform init
ls -lha
tree .terraform
terraform plan
docker ps
docker ps -a
```

Expected output
![image](https://github.com/user-attachments/assets/2ae2b464-4aeb-4088-86fd-7bc2e32342fb)
![image](https://github.com/user-attachments/assets/7b32b05f-57b4-4402-8fae-117a8ef07e19)
![image](https://github.com/user-attachments/assets/274dbc06-735d-4830-b5b8-429da00f6230)
![image](https://github.com/user-attachments/assets/0c7b6e8c-9afe-4a21-9d74-3b49e6b1a01e)
![image](https://github.com/user-attachments/assets/4d6c2941-9bfe-49e4-a0ac-c3391788fadf)
![image](https://github.com/user-attachments/assets/ba731016-cda3-47d9-b6bb-d495196e23dd)
![image](https://github.com/user-attachments/assets/8fc5b1d2-df6e-4ec8-9381-9ff155647dbf)
![image](https://github.com/user-attachments/assets/d58e394e-e434-47b9-b90e-5174c8626153)
![image](https://github.com/user-attachments/assets/4b203b35-6bc4-4bb2-a177-461b5781f922)
![image](https://github.com/user-attachments/assets/306d8ed9-1c08-49b2-acae-9c7c455f66f4)

## Info - Terraform Lifecycle commands
To get all possible terraform command options
```
terraform
```
Expected output
![image](https://github.com/user-attachments/assets/56ae9662-0b22-4dd8-af9f-5a3d4796a591)

```
terraform init
terraform plan
terraform apply
terraform refresh
terraform destroy
```

## Lab - Invoking ansible playbook from Terraform
```
cd ~/terraform1721-march-2025
git pull
cd Day4/terraform/invoking-ansible-playbook-from-terraform
terraform init
terraform apply --auto-approve
docker ps
curl http://localhost:8001
```

Expected output
![image](https://github.com/user-attachments/assets/3658a754-1766-4066-974a-aeefe7d72ed7)
![image](https://github.com/user-attachments/assets/ce27ed77-6430-4ac7-aaf6-464d6fd11072)
![image](https://github.com/user-attachments/assets/fd007e89-0bbb-4427-9127-787cf36c6464)
![image](https://github.com/user-attachments/assets/1c701b39-82f9-4f2a-826b-2c5902150fb8)
![image](https://github.com/user-attachments/assets/8332e84a-26a1-4b5a-8dbf-a5aae3713633)
![image](https://github.com/user-attachments/assets/5affeccc-5355-4670-8d35-04015d1f3307)

## Lab - Using Terraform count and remote-exec
```
cd ~/terraform1721-march-2025
git pull
cd Day4/terraform/remote-exec
terraform init
terraform apply --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/112b74d7-d17e-47e5-b344-02bd98e53b23)
![image](https://github.com/user-attachments/assets/1994f620-714b-4a3f-a122-34f7ad80035f)
![image](https://github.com/user-attachments/assets/9fb555c2-aaf5-4cbf-8826-5cb649866e3a)
![image](https://github.com/user-attachments/assets/fb1c6a98-d832-4735-999d-3a1fdf70a9fd)
![image](https://github.com/user-attachments/assets/16e6196b-0654-452d-975d-010434dbe385)
![image](https://github.com/user-attachments/assets/1dffc92d-09f0-4ee5-b4c6-7caf8dfe8f33)
![image](https://github.com/user-attachments/assets/10d534df-56c2-41e2-a24a-ef76d705ee8e)
![image](https://github.com/user-attachments/assets/7c50f083-b30e-4ff4-9f4c-db4a739a386a)
![image](https://github.com/user-attachments/assets/15951355-b67e-47bb-8efe-383500751fdb)
![image](https://github.com/user-attachments/assets/a52de9f3-c84f-48ad-a86e-6d0a6f6255f4)
![image](https://github.com/user-attachments/assets/81b12c0f-dd42-48cb-b22a-57a676cac5e2)
![image](https://github.com/user-attachments/assets/420912b1-b5c9-4bdf-b05f-bf2662ffce3e)
![image](https://github.com/user-attachments/assets/6c5f9095-7233-41c2-90f5-ffa2358566de)
![image](https://github.com/user-attachments/assets/2c67148b-5e8d-4ca1-ab8c-97a9b6499665)
![image](https://github.com/user-attachments/assets/d2826bdd-23cb-4571-b0e1-95cdad69aaa1)
![image](https://github.com/user-attachments/assets/cc37e612-6738-4cbb-847c-44a9c792603c)
![image](https://github.com/user-attachments/assets/dccc6ca7-8351-47d1-a3b7-588db23d727d)
