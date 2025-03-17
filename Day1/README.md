![image](https://github.com/user-attachments/assets/93fe1980-87d6-414c-b115-2501b4c89356)# Day 1

## Info - Dual booting or multi-booting
<pre>
- Boot loader system software utilities that is installed in Hard Disk - Master Boot Record(MBR)
- Master Boot Record is Sector 0, Byte 0 in the hard disk which is just 512 bytes
- Boot Loader software normally has to fit within the 512 bytes, it gets installed in MBR
- When the system is booted, after the BIOS POSt is completed, the BIOS instructs the CPU to run the Boot Loader
- The Boot Loader utility searches the hard disk looking for Operating System installed in it, if the number of OS is more than one, it prepares a menu and presents the menu for us to choose the OS we wish to boot into
- Only one OS can be active at any point of time
- Examples
  - LILO ( Linux Loader - legacy boot loader that was used in older Linux distributions )
  - GRUB - almost all Linux distributions
  - BootCamp ( to support dual boot with Mac OS-X or Windows )
</pre>  

## Info - Hypervisor
<pre>
- virtualization technology
- with virtualization, we can run many OS side by side on the same machine ( laptop, desktop, workstation or Server )
- many OS can be actively run at the same time on the same machine
- two types of Hypervisors
  - Type 1 
    - Bare Metal Hypervisors
    - Doesn't require a Host OS to install the Hypervisor software
    - Which is used in Servers & Workstations
    - examples
      - VMWare vSphere/vCenter
      - KVM 
  - Type 2
    - is used in laptops/desktops/workstations
    - requires Host OS ( Windows, Mac, Linux )
    - examples
      - VMWare Workstation ( Windows & Linux )
      - VMWare Fusion ( Mac OS-X )
      - Parallels ( Mac OS-X )
      - Oracle VirtualBox ( Windows, Linux & Mac )
      - Microsoft Hyper-V ( Windows )
- each Guest OS is installed in a Virtual Machine (VM)
- each Guest OS is a fully functional Operating System with their own hardware and dedicated OS Kernel, etc
- for each VM, we must allocated dedicated hardware resources, hence it is considered a heavy weight virtualization technology
  - CPU Cores
  - RAM
  - Storage
  - Virtual Network Card(s)
  - Virtual Graphics Cards(s)
</pre>

## Info - Hypervisor - High Level Architecture
![Hypevisor](HypervisorHighLevelArchitecture.png)

## Container Technology
<pre>
- Containerization - is an application virtualization technology
- light-weight virtualization technology
- each container, represents one application or application component
- containers are not Operating System
- containers are just application process not a OS
- containers don't even has their own dedicated hardwares or Operating System kernel
- all the containers that runs on a system shares the underlying HOST OS Kernel and Hardware resources
- contaners have some similartity to a virtual machine
  - just like VMs, containers are also allocated with Ip address
  - just like VMs, containers has its own file systems
  - just like VMs, containers has their own network stack ( 7 OSI Layers )
  - just like VMs, containers has their own port range ( 0 - 65535 )
- Linux Kernel container related features
  - Namespace and
    - helps isolating one container from other containers
  - Control Groups (CGroups)
    - helps applying resource quota restrictions for each container
      - how many cpu one container can use at the max
      - how much RAM one container can use at the max
</pre>

## Info - Container Runtime
<pre>
- is a low-level software that manages container images and containers
- is not so user-friendly, hence normally no end-users use it directly
- examples
  - runC
  - CRI-O 
</pre>

## Info - Container Engine
<pre>
- is a high-level software that manages container images and containers
- provides user-friendly command to manage images and containers, even without low-level linux kernel knowledge one can manages manages and container with the help of container engine
- container engines internally makes use of Container Runtime to manage images and containers
- examples
  - Docker is a Container Engine that depends on containerd which internally depends on runC Container Runtime
  - Podman is a Container Engine that depends on CRI-O container runtime
</pre>

## Info - Container Images
<pre>
- is a blueprint of a container
- it is similar to Windows.iso OS images, with single Windows12.iso we can install Windows 12 on multiple machines, same way with Container image we can creates multiple container instances
- Container Images does'nt represent OS
- container Images normally has
  - some application and its dependent libraries/frameworks, etc.,
  - supports file systems
  - some shell bash or sh
  - optional package managerpu
    - apt/apt-get package manager in case of ubuntu:latest docker image
    - yum/dnf/rpm package managers are supported in case of rhel:latest or rocky-linux:latest
</pre>

## Info - Containers
<pre>
- is an instance of a container image
- container represents an application 
- all the tools and sofwares installed in the container image are available in every container that is created out of the container image
</pre>

## Info - How Linux containerized applications run in Windows/Mac?
<pre>
- When we install Docker for Windows on any Windows OS, it installs a thin-linux layer that has the Linux kernel
- When a linux containerized application runs on Windows, it is still running on top of Linux kernel on a Windows machine
- When we install Docker for Mac on any Mac machine, it installs a thin-linux layer that has the Linux kernel
- When a linux containerized application runs on Mac, it is still running on top of Linux kernel on a Mac machine
</pre>

## Info - Image Registry
<pre>
- is a server that manages/maintains multiple container images
- there are 3 types of Image Registries
  1. Local Image Registry
     - it is folder under the home directory of a user
     - has a collection of many container images
  2. Private Image Registry
     - is a server
     - has a collection of many container images
     - can be setup using JFrog Artifactory or Sonatype Nexus
  3. Remote Registry
     - is a website powered by a Server ( JFrog Artifactory or Sonatype Nexus or something similar )
     - has a collection of many container images
</pre>

## Info - Docker Container Engine
<pre>
- is developed in Go language by Docker Inc company
- comes in 2 flavours
  1. Docker Community Edition - Docker CE ( opensource )
  2. Docker Enterprise Edition - Docker EE ( paid commercial tool )
- follows client/server architecture
  - docker client is a utility name docker
  - docker server runs a linux/windows service ( dockerd - daemon )
</pre>

## Info - Docker High Level Architecture
![Docker](DockerHighLevelArchitecture.png)


## Info - Configuration Management Tools
<pre>
- helps us automate system administrative activities
- assume you have some server with certain OS pre-installed in it ( already provisioned machine )
  - OS could be Windows/Mac/Linux/Unix/Network Switch/Router
  - ansible or any configuration management tool can help you in automating the configuration management activities
  - can install/uninstall/upgrade/downgrade softwares
  - can manage users on an existing Windows/Linux/Unix machines
  - reboot a machine
  - install software patches or service packs in Windows server
- examples
  - Puppet
  - Chef
  - SaltStack or Salt
  - Ansible
</pre>


## Info - Puppet Configuration Management Tool Overview
<pre>
- is developed and maintained by a company called Perforce
- comes in 2 flavours
  - Puppet Community edition ( opensource )
  - Puppet Enterprise ( Paid version for commercial use )
- follows Client/Server architecture
- follows PULL based architecture
- each machine with OS that must be managed by Puppet should have a proprietary tool running called Puppet Agent
- Domain Specific Language (DSL) is the language used to automate the configuration management activities in declarative style
- DSL used by Puppet is Puppet Configuration Language ( Proprietary Language )
- learning curve is steep
- installation is complex
</pre>
https://blog.nashtechglobal.com/what-is-puppet-and-its-architecture/

## Info - Chef Configuration Management Tool Overview
<pre>
- developed and mantained by a company called ProgressChef
- follows client/architecture architecutre
- the chefs nodes ( servers managed by Chef ) must have a proprietary agent installed in them called Chef Agent
- DSL used is Ruby ( scripting language )
- learning curve is very steep as it comes more than 15 or more tools
- installation is complex
</pre>

## Info - Ansible Configuration Management tool overview
<pre>
- Ansible is developed in Python by Michael Deehan ( former employee of Red Hat )
- Michael Deehan was an architect at Red Hat, he had a configuration management tool project idea, for some reason Red Hat had put that idea on shelf
- Michael Deehan had quit Red Hat and started a company called Ansible Inc and developed Ansible core as an opensource project
- When Ansible core came into the industry, already many companies were either using Puppet/Chef
- follows a very simple agent-less architecture, not a client/server architecture, stand-alone tool
- follows a PUSH based architecture
- DSL used is YAML
- Ansible nodes
  - the machines managed by Ansible are called Ansible nodes
  - this could a physical server with some OS
  - this could a virtual machine running in your private datacenter
  - this could a virtual machine running on some public cloud like AWS/Azure/GCP/Digital Ocean, etc.,
  - could be Network switch/router
- Ansible Controller Machine(ACM)
  - is the machine where Ansible is installed ( this could be your laptop/desktop - your work machine )
  - it must be Linux OS ( Unix or Mac - officially only Linux machine can be an ACM )
- Ansible Nodes software dependencies
  - Unix/Linux/Mac
    - Python and SSH Server must be pre-installed/enabled/configured
    - any unix/linux/mac server OS by default comes with Python and SSH Server pre-installed
  - Windows Server
    - PowerShell and WinRm must be pre-installed/enabled/configured
    - any Windows Server, by default comes with Powershell and WinRM pre-installed
  - Playbooks - declarative automation script is written in YAML format, hence python knowledge is not required to automate
  - installing ansible and learning ansible is easier compared to other configuration management tools
- comes in 3 flavours

  1. Ansible Core 
     - open source and supports only Linux Command Line Interface
     - no support available
     - no role-based access ( lacks user management )
     - no GUI

  2. Ansible AWX 
    -  developed on top of Ansible Core, supports Webconsole, opensource 
    - no support available
  
  3. Red Hat Ansible Automation Platform ( a.k.a in the past as Ansible Tower - Enterprise product requires license )
     - developed on top of opensource AWX
     - supports only webconsole
     - Red Hat gives world-wide support
</pre>


## Info - Ansible Modules
<pre>
- When we install ansible, it installs many Ansible modules 
- ansible modules are
  - python scripts that helps automating one thing in unix/linux/mac
  - powershell scripts that helps automating one thing in Windows
- Examples
  - If you wish to copy a file from local machine to ansible node, there is a ready-made module called copy 
  - If you wish to create a folder in the anisble node, there is a ready-made module called file
  - If you wish to install softwares in Debian Linux, you have an apt module 
  - If you wish to install softwares in Red Hat Linux family, you have a module called yum 
- normally ansibles comes with almost 8300+ modules, if you find a scenario where those 8300 existing ansible modules can't automate something, that is the only scenario you have to develop your own custom ansible module, in that scenario one is expected to have python or powershell knowledge
</pre>

## Info - Ansible Playbook
<pre>
- is the declarative automation code we writing in Ansible to perform configuration management
- is an YAML file ( superset of JSON - JavaScript Object Notation )
- it invokes one to many ansible modules to perform the automation
</pre>

## Info - Ansible High-Level Architecture
![Ansible](AnsibleHighLevelArchitecture.png)

## Demo - Installing Ansible Core in Ubuntu Linux
```
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

## Lab - Checking docker version
```
docker --version
docker info
```

Expected output
![image](https://github.com/user-attachments/assets/944c7d29-3217-4951-863f-46ff62f0232a)
![image](https://github.com/user-attachments/assets/9145954f-0fb2-439c-9b96-06b865129404)

## Lab - Checking ansible core version
```
ansible --version
```

Expected output
![image](https://github.com/user-attachments/assets/f9a42da0-810d-4cd3-9878-f58a11d491c4)

## Info - Ansible inventory
<pre>
- has connection details to ansible nodes
- it is a text file
- winrm connection details in case the nodes happens to be a Windows server
- ssh connnection details in case the nodes happens to be a Unix/Linux/Mac Server
- there are two types
  1. Static inventory ( text file with usually no extenstion - recommended names are hosts or inventory )
  2. Dynamic Inventory ( Python script )
</pre>

## Lab - Cloning Terraform Training GitHub Repository ( Ubuntu - RPS Cloud machine )
```
cd ~/
git clone https://github.com/tektutor/terraform-1721march-2025
cd terraform-1721march-2025
```

## Lab - Installing SSH Server in RPS Ubuntu Lab machine
```
sudo apt install -y net-tools openssh-server tree curl iputils-ping
```

## Lab - Building an Ubuntu Custom Docker Image to use it as an ansible node
```
cd ~/terraform-1721march-2025
git pull
cd Day1/CustomDockerImages/ubuntu
cat Dockerfile
ssh-keygen
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible-node:latest .
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/2abba6ec-d6e6-49a4-b2a4-2e1adc28e014)
![image](https://github.com/user-attachments/assets/d17e0e19-133d-46f5-803d-3db4530e8cdf)
![image](https://github.com/user-attachments/assets/62f8ef81-dd05-4b65-99af-c61b1764132c)
![image](https://github.com/user-attachments/assets/eafb7656-bcc0-4da0-ba5a-7727e34338f5)

Troubleshooting, the docker rate pull limit reached error. Those who are getting this error, alone can follow the below instructions
Run this on the terminal
```
docker login
```

Then the docker login url will be displayed, copy/paste the url in some web browser on the lab machine and paste the code shown in the terminal. You need to login with the below credentials
<pre>
username - alchemysolutions003@gmail.com
password - Alchemy@321
</pre>

Once you successfully, logged in it show, successfully logged in as shown belo
![image](https://github.com/user-attachments/assets/9d6db808-d34a-4a37-b198-881e577cd5f3)
![image](https://github.com/user-attachments/assets/488c6481-8df7-4cff-a765-0418662f8558)

## Lab - Creating couple of containers using our custom docker ansible node image

Let's create two container using our custom ubuntu ansible node docker image
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```

Listing the running containers
```
docker ps
```
Expected output
![image](https://github.com/user-attachments/assets/04518456-09b1-4583-9b5b-7c5e0b8a022e)


## Lab - Check if you are able to SSH into the ubuntu1 and ubuntu2 containers
```
ssh -p 2001 root@localhost
exit

ssh -p 2002 root@localhost
exit
```

![image](https://github.com/user-attachments/assets/262e1b41-a604-4f6a-ad68-33e644f8de78)
![image](https://github.com/user-attachments/assets/67873d05-add3-4f0f-96a1-99a60fbe27eb)

Troubleshooting, permission denied error. You need to delete the image
```
docker rm -f ubuntu1 ubuntu2
docker rmi tektutor/ubuntu-ansible-node:latest
cd ~/terraform-1721march-2025
git pull
cd Day1/CustomDockerImages/ubuntu
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible-node:latest .
docker images

docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
docker ps

ssh -p 2001 root@localhost
exit

ssh -p 2002 root@localhost
exit
```


## Lab - Writing an Ansible static inventory file
In the below command, 
<pre>
i - inventory flag
inventory - inventory file name
all - all the machines listed under all group in the inventory file
m - module flag
ping - is the module name (ping.py script - ansible module )
</pre>

```
cd ~/terraform-1721march-2025
git pull
cd Day1/ansible
cat inventory
ansible -i inventory all -m ping
ansible -i inventory ubuntu1 -m ping
ansible -i inventory ubuntu2 -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/b11114ec-7ac5-4488-8d77-948459100591)
![image](https://github.com/user-attachments/assets/66d5a2ae-2258-4d66-8363-157f074338d1)

## Lab - Collecting ansible facts using ansible module setup
```
cd ~/terraform-1721march-2025
git pull
cd Day1/ansible
cat inventory
ansible -i inventory ubuntu1 -m setup
```
Expected output
![image](https://github.com/user-attachments/assets/6bbae6e5-ba04-4987-95b1-ce2c735fea4a)
![image](https://github.com/user-attachments/assets/52b59a3e-ecde-4222-ab60-d37d1efa7beb)
![image](https://github.com/user-attachments/assets/d749ed4f-0d49-404c-a3f3-2deeb206bb35)

## Lab - Finding the list of all ansible modules supported by your ansible installed on the lab machine
```
ansible-doc -l
```

Expected output
![image](https://github.com/user-attachments/assets/26024970-7950-4f95-a4ed-fbcc5a7ac396)
![image](https://github.com/user-attachments/assets/16eb9a18-4429-4dee-a22c-562b048cffd3)


Finding number of modules supported by ansible
```
ansible-doc -l | wc 
```

## Info - Imperative vs Declarative language
<pre>
Imperative Programming Language
- They are powerful languages
- They are ideal for application development, not ideal for automation
- Assume I wanted to automate installing Weblogic latest version in all my Linux servers
- You need to write code for
  - What - I wanted to automate installing Weblogic latest version in all my Linux servers
  - How - We need to write code in Python to automate our requirement
- Examples
  - C/C++, Java, Dot Net, Python, Ruby 
Declarative Language
- Ansible Playbook is a declarative code
- You just need to mention what you wish to do
- the language will take care of how to do it
</pre>

## Lab - Writing our first ansible playbook
```
cd ~/terraform1721march-2025
git pull
cd Day1/ansible
cat inventory
cat ping-playbook.yml
ansible-playbook -i inventory ping-playbook.yml
ansible-playbook -i inventory ping-playbook.yml -vvvv
```

Expected output
![image](https://github.com/user-attachments/assets/3d2dff23-3b0e-46d2-afc9-a21f3b02e7e0)
![image](https://github.com/user-attachments/assets/089a8026-4b88-4fff-a3c0-aada3c61fba1)
![image](https://github.com/user-attachments/assets/6b182fed-a3a8-481e-a2bb-bf99cbf9378b)
