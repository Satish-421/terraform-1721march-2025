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
