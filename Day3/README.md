# Day 3

## Info - Ansible Automation Platform
<pre>
- was also called as Ansible Tower
- it is developed on top the AWX (open source)
- AWX is developed on top Ansible core (open source)
- unlike Ansible core, Ansible automation platform supports webconsole, user management, etc.,
- this is a Red Hat Enterprise product with world-wide support
- functionally AWX and Ansible Automation Platform are one and same
- it is not possible develop Ansible Playbook within AWX or Ansible Automation Platform
- hence, we still need Ansible core to develop playbooks and test them before we push it to GitHub or any version control
- the existing Ansible Playbooks we can execute via Ansible Automation Platform
- can be installed as a stand-alone application on any Linux Distributions or we can install inside Kubernetes or Openshift
</pre>

## Lab - Starting Minikube to use Ansible Tower
```
minikube status
minikube start
minikube status
kubectl get nodes
```

Expected output
![image](https://github.com/user-attachments/assets/310342c7-b2e8-4edd-aa4e-3d5b15fb60ea)
![image](https://github.com/user-attachments/assets/befcda35-93a5-4a92-ab74-8773102bc9d5)

Launching AWX - opensource Ansible Automation Platform
```
minikube service awx-demo-service --url -n ansible-awx
```

You can launch the AWX webconsole using the url shown by the above command in the Ubuntu RPS cloud machine
<pre>
http://192.168.49.2:31225
</pre>
![image](https://github.com/user-attachments/assets/70d5fb5d-62fb-4321-b5c4-e5c2792b0c51)


To retrieve the AWX password ( save the password in a file to avoid typing this lengthy command each time )
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Login credentials for AWX is
<pre>
username - admin
password - 
</pre>
![image](https://github.com/user-attachments/assets/910b3144-e59c-459a-8c5b-26788553153c)
![image](https://github.com/user-attachments/assets/fb6b39b7-9752-4ce6-83b7-45ba185e7afa)
![image](https://github.com/user-attachments/assets/06607ea3-0fbf-4dfa-8cd4-0f4e78af0221)


## Lab - Creating a Project in Ansible Automation Platform
Navigate to your Ansible Automation Platform Dashboard page
![image](https://github.com/user-attachments/assets/23c4d27e-190c-49a5-9b82-29956b572b87)

Under Resources menu, select "Projects"
![image](https://github.com/user-attachments/assets/9025034d-15fe-4b68-9e4c-bfe5ffd0b3c7)
Click "Add" to create a new project
![image](https://github.com/user-attachments/assets/654db108-2a5e-4d39-a5fc-112fc58c5839)
Select "Git"
![image](https://github.com/user-attachments/assets/35e47d0e-05dc-43e0-a1c1-8c2d2c99c829)
![image](https://github.com/user-attachments/assets/25ce7682-a2f5-4424-96ac-8142a4929bb2)
![image](https://github.com/user-attachments/assets/8e51cd5c-3360-4fcd-8d41-941018d2cc2e)
Click "Save"
![image](https://github.com/user-attachments/assets/8d4c187b-8e21-4502-812b-1280be1cc3de)
![image](https://github.com/user-attachments/assets/104e7328-4198-4c22-9a9d-1eeb2b4e1b9b)
![image](https://github.com/user-attachments/assets/705c3006-183a-43dc-a455-1953a159a6fd)
![image](https://github.com/user-attachments/assets/31880195-ed4a-431a-adce-20d9b2460cae)
![image](https://github.com/user-attachments/assets/2f9f42f3-ff4c-455a-ba9c-4aa2c2bce3f3)
![image](https://github.com/user-attachments/assets/85e9b257-048c-4dea-ad77-302d6e9ecc8a)

## Lab - Creating an Inventory in Ansible Automation Platform
Navigate to your Ansible Automation Platform Dashboard page
![image](https://github.com/user-attachments/assets/23c4d27e-190c-49a5-9b82-29956b572b87)

Under Resources menu, select "Inventories"
![image](https://github.com/user-attachments/assets/b19a8acf-a079-4dd6-a387-d1a6c69589a5)
Click "Add" to create an Inventory
![image](https://github.com/user-attachments/assets/232c3bd3-617c-4908-96cf-538bd3cf81c2)
Select "Add Inventory"
![image](https://github.com/user-attachments/assets/b2361f4e-ccf9-438f-ba4e-235dbb5bee7c)
![image](https://github.com/user-attachments/assets/6c8c0acd-bdce-45c9-b9cb-a8654aaa72d3)
Give a name for your inventory
![image](https://github.com/user-attachments/assets/f51cddfa-d894-46ac-86ef-c8e9d8d65744)
Click "Save" button
![image](https://github.com/user-attachments/assets/d36f49d0-6e30-499e-855d-d11ec2ad4eec)

Click "Hosts" Tab within the inventory you created
![image](https://github.com/user-attachments/assets/5e2517ac-c132-4d6b-a650-3d81e594428e)

Find your Lab machine IP address
![image](https://github.com/user-attachments/assets/a4f246b8-52be-4dac-a782-1fb3d048cce6)

Click "Add" to add a Host(Ansible node connection details)
![image](https://github.com/user-attachments/assets/9c4e47b5-8f24-49b7-9340-817f54ddc133)
![image](https://github.com/user-attachments/assets/1f08e35a-24fc-47a1-a3c3-594d7a76f88b)
Click "Save"

Click "myinventory", "hosts" 
![image](https://github.com/user-attachments/assets/127a846a-fca6-46a1-8db8-47bdfc9f61e6)
![image](https://github.com/user-attachments/assets/a68e73e4-8023-4465-a5f8-0bccc03ae056)
![image](https://github.com/user-attachments/assets/9919d434-f54d-435b-9631-20500c712663)
![image](https://github.com/user-attachments/assets/ca5f6d2c-ba8f-46bf-8ae4-bc0a776ecfb2)
![image](https://github.com/user-attachments/assets/e8abf058-fa5a-4394-92b1-823609d81be8)
![image](https://github.com/user-attachments/assets/7a3f3d56-30d2-432c-b5ef-9bd6e88b1785)
![image](https://github.com/user-attachments/assets/b989c1ad-c146-4df4-80e1-53e2c240843a)
![image](https://github.com/user-attachments/assets/d7ba0b8d-5800-429b-acf1-eb13b5a283e5)
![image](https://github.com/user-attachments/assets/5493ce9a-35d6-420f-8452-7e5240ce7be0)


## Lab - Create Credentials
Navigate to your Ansible Automation Platform Dashboard page
![image](https://github.com/user-attachments/assets/23c4d27e-190c-49a5-9b82-29956b572b87)

Under "Resources" menu, select "Credentials"
![image](https://github.com/user-attachments/assets/9782831f-630c-4796-9402-b83be9cc87ab)
Click "Add" to create credentials
![image](https://github.com/user-attachments/assets/e66ba643-569f-4020-9937-1f056a00d5d6)
![image](https://github.com/user-attachments/assets/6d5e84af-db22-4352-9134-9135efee0051)
![image](https://github.com/user-attachments/assets/0de5c301-e28c-4453-8710-562cbce95d3a)
![image](https://github.com/user-attachments/assets/045f6044-589d-49eb-8c62-f2831d6e7085)
![image](https://github.com/user-attachments/assets/26d32cf0-60bd-403b-8681-bc1552f70796)
![image](https://github.com/user-attachments/assets/7727ee41-de92-42a0-91d2-69ad8aca42eb)
![image](https://github.com/user-attachments/assets/eb730661-8bf3-4bf9-80e8-fe8cc10a0445)
Click "Save"
![image](https://github.com/user-attachments/assets/e279adc4-fd37-4a9e-be84-52183cc1b949)
![image](https://github.com/user-attachments/assets/5e6d4122-da26-4756-a075-de4b58b3444f)



##

Click "Run command"
![image](https://github.com/user-attachments/assets/d2f1c793-0c2e-4a63-b70c-d7ee3ab26c0a)
![image](https://github.com/user-attachments/assets/e818aa76-00f7-41a0-970a-239d8d7ce54d)
Choose a module, select ping
![image](https://github.com/user-attachments/assets/2edbe0ba-e718-4320-b01d-d6da89c7320f)
![image](https://github.com/user-attachments/assets/8b0683ea-b355-4637-9515-f8b05be84a56)
Click "Next"
![image](https://github.com/user-attachments/assets/29f7c5e4-080e-4139-81b0-8f798fcd90c8)
Click "Next"
![image](https://github.com/user-attachments/assets/421b7806-319d-494f-be58-ab32164299aa)
![image](https://github.com/user-attachments/assets/a53735eb-a99d-4eaa-81ec-fdd953241e98)
![image](https://github.com/user-attachments/assets/d72dc32d-e518-4404-b1a3-ab94ac28e1ef)
Click "Launch"
![image](https://github.com/user-attachments/assets/35459970-67de-42c0-a049-5985fbaaab97)
![image](https://github.com/user-attachments/assets/5dcd6fa0-78fb-45cc-abbb-8c6f10f1c751)

## Lab - Creating a Job Template to invoke Ansible Playbook
Navigate to your Ansible Automation Platform Dashboard page
![image](https://github.com/user-attachments/assets/256e8757-4e9b-4cb2-b8ab-8343e2a73dd1)

On the Resources menu, click "Templates" to create a Job Template 
![image](https://github.com/user-attachments/assets/48e06b81-27ad-4121-bdb4-0b00a3dd3ccc)

Click "Add"
![image](https://github.com/user-attachments/assets/31f8b13f-c94c-4856-bcd2-caf79c79519d)
Select "Add Job Template"
![image](https://github.com/user-attachments/assets/ef9ee9da-02fb-4279-93f1-fdf50d9810f6)

![image](https://github.com/user-attachments/assets/b7ff8578-0ee8-4420-b5fc-2b7ef5474f52)
![image](https://github.com/user-attachments/assets/c2a7146f-6c29-4bb6-91d8-94f66c625559)
![image](https://github.com/user-attachments/assets/8852d272-b637-4cc9-89b5-d86cdc61dd21)
![image](https://github.com/user-attachments/assets/888939f8-d351-475f-a51a-f725d0f6a6aa)
![image](https://github.com/user-attachments/assets/8f5b0334-1bf5-4690-8577-b589904c9147)
![image](https://github.com/user-attachments/assets/977dda8c-1496-4a16-b89f-56d17e58130f)
![image](https://github.com/user-attachments/assets/f4438d57-2cc0-4d7b-bea3-cdde67b4359f)

![image](https://github.com/user-attachments/assets/d33159ef-1eb6-475e-9858-77a5ccc9a3ca)
Click "Save"
![image](https://github.com/user-attachments/assets/739330d0-717a-4ad6-9d08-a1281b3b69a2)
![image](https://github.com/user-attachments/assets/a7ef8b36-9a94-4156-8246-0773bce3d1ac)
Click "Launch" to run the playbook

![image](https://github.com/user-attachments/assets/b2793f94-4872-4184-8706-64f181711bf1)

![image](https://github.com/user-attachments/assets/098bf9d1-77b5-4724-bd7a-0aeddff6f977)

![image](https://github.com/user-attachments/assets/f940905f-5bb8-4f3d-868e-fe3d59373adb)


## Lab - Create a normal user
![image](https://github.com/user-attachments/assets/79f13f3a-12ed-4486-a525-1d84243b7d02)


## Lab - Installing tower cli
```
sudo apt install -y python3-pip
pip install ansible-tower-cli --break-system-packages
```
![image](https://github.com/user-attachments/assets/d18ad3b7-010a-4762-baed-42ffb4f40c2f)
![image](https://github.com/user-attachments/assets/1b3a19c8-3b76-47fd-90eb-d34e454c4574)
![image](https://github.com/user-attachments/assets/0b090ba2-dd13-4f14-b332-9d2bdac2da0c)


## Lab - Using tower-cli
```
export PATH=$PATH:/home/rps/.local/bin
tower-cli config host http://192.168.49.2:31225
tower-cli config username admin
tower-cli config password your-admin-password
tower-cli config verify_ssl false
tower-cli project list
tower-cli job list
tower-cli job launch --job-template=9
```

## Info - Go Programming Language Overview
<pre>
- is a programming language developed in C by Google
- later the go lang is recompiled/redeveloped in Go lang
- syntax wise, it is more or less similar to C Programming language
- also called as Golang
- golang design philosophy is to
  - keep it simple, hence it supports a total of just 25 keywords
  - loops - only for loop is supported
  - there are no class, only struct is supported to define user-defined data-type
  - supports pointers but manages memory using garbage collector
- golang supports the below basic data-types
  - bool
  - string
  - int, int8, int16, int32 and int64
  - uint, unit8, uint16, uint32, uintptr
  - byte ( alias of uint8 )
  - float32, float64( equivalent to double in most programming languages )
  - complex64, complex128
- it is a statically typed programming language
  - meaning, every variable must be declared before using it
- strongly typed
  - once a variable is declared as int, it will remain as integer throught the life-time of the application
- golang supports garbage collection
- some of the tools developed in golang
  - can be used to develop normal desktop applications, system utilities, compilers, interpreters, web application, AI 
  - Docker, Kubernetes, OpenShift, Terraform, DropBox, etc.,
- it is faster than most compiled languages
  - faster than .Net, Java, etc.,
  - equivalent to C/C++ in terms of performance
</pre>

## Lab - Checking if go is installed in your lab machine
```
go version
```

Expected output
![image](https://github.com/user-attachments/assets/638a308f-8af1-4697-994f-ecbc03ff46d7)

## Lab - Writing your first Hello World application in Golang

Create a file named hello.go with the below content

<pre>
package main

import "fmt"

func main() {
  fmt.Println ( "Hello World !" )
}
</pre>

Run the program
```
go run ./hello.go
```

Expected output
