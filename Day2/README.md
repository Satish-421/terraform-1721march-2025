# Day 2

## Info - Linux Distributions
<pre>
- Linux Operating System is distributed by many organizations
- Each Linux OS belongs to different Linux OS Family
- Each Linux Distribution
- Examples
  - Ubuntu - is a Linux distribution that belongs to Debian Linux Family
  - Fedora - is a Linux distribution that belongs to Red Hat Linux Family
  - Rocky Linux - is a Linux distribution that belongs to Red Hat Linux Family
  - Kali Linux - is a Linux distribution that belongs to Debian Linux Family
  - Red Hat Enterprise Linux(RHEL) - is a Linux distribution that belongs to to Red Hat Linux Family
- Each Linux Family supports different Package Managers
- Package Manager
  - is a tool that helps us install/uninstall/upgrade softwares in Linux
  - Ubuntu supports a package manager called apt/apt-get
  - Fedora supports package managers called dnf,yum,rpm
  - RHEL supports package managers dnf,yum,rpm, etc.,
- Linux Repository Servers
  - Each Linux Distribution company maintains a Server that has test,stable softwares for every version of Linux distribution
  - We need to update the repository url, and then we can install any linux software from the repository
  - Package Managers will be using the repository configuration files to find the repository url
</pre>  

## Lab - Understanding host and group variables in ansible inventory file
<pre>
- whatever variables are unique to an ansible node are called as host variable
- for example, the ansible_port variable is called a host variable as the ssh port on the local machine varies for each container
- the common variables for ubuntu1 and ubuntu2 are defined in the group level as they are common to ubuntu1 and ubuntu2. These variables are called group variables
- ansible_host, ansible_private_key_file, ansible_user are called group variables in this inventory
</pre>

Make sure the ubuntu1 and ubuntu2 containers are running
```
docker ps -a
docker start ubuntu1 ubuntu2
docker ps
```

Now you may proceed
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible
cat hosts
ansible -i hosts all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/9ff10406-dc29-45d4-a696-84f4bac08854)

## Lab - Installing nginx web server using ansible playbook
```
cd ~/terraform1721-march-2025
git pull
cd Day2/ansible
cat hosts
cat install-nginx-playbook.yml
ansible-playbook -i hosts install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/3e6ca9c5-81e5-4a72-bb68-dfd5791c8930)

Let's find the IP address of ubuntu1 and ubuntu2 ansible node containers
```
docker inspect -f {{.NetworkSettings.IPAddress}} ubuntu1
docker inspect ubuntu2 | grep IPA
curl http://172.17.0.2:80
curl http://172.17.0.3:80
```
Let's check if the nginx server is up and running using ansible ad-hoc command
```
ansible -i hosts all -m shell -a "service nginx status"
```
![image](https://github.com/user-attachments/assets/eb06b7d6-3c48-49f9-8b33-6f4e6c230e0f)
![image](https://github.com/user-attachments/assets/3e42472f-b836-4522-8a55-1c0bf086d9e9)
![image](https://github.com/user-attachments/assets/7b5b25aa-4a88-4d6f-ba09-e78fdafec832)

Let's test the web page response
```
curl http://localhost:8001
curl http://localhost:8002
```

Expected output
![image](https://github.com/user-attachments/assets/a140c2a6-e40a-4fa9-8f7a-66b66b7af93c)
![image](https://github.com/user-attachments/assets/39ef2320-95c9-4ac5-863c-b2d28adf6c3e)
![image](https://github.com/user-attachments/assets/f24f025d-0205-4d8d-b2b4-4dd6218a5883)
![image](https://github.com/user-attachments/assets/f62d7738-9971-4422-8bee-7fedd01e8c82)
![image](https://github.com/user-attachments/assets/f5801875-b1a6-43fa-956b-54e907b9a5ba)
![image](https://github.com/user-attachments/assets/8897d511-0caf-4d21-9916-27af923ce303)
![image](https://github.com/user-attachments/assets/695316de-062b-4c29-b14e-a1435a85755c)
![image](https://github.com/user-attachments/assets/6302e29d-e903-4665-b7ec-6d2311d2602e)
![image](https://github.com/user-attachments/assets/5c4045d7-860d-467b-8675-1da2dda02d41)
![image](https://github.com/user-attachments/assets/fa6fcd31-d9da-448b-a756-953f9c1b0bab)

## Lab - Building a Custom Rocky Linux Ansible Node Image
```
cd ~/terraform-1721march-2025
git pull
cd Day2/CustomDockerImages/rocky
cp ~/.ssh/id_ed25519.pub authorized_keys
ls -l
docker build -t tektutor/rocky-ansible-node:latest .
docker images
rm authorized_keys
```

Expected output
![image](https://github.com/user-attachments/assets/469e0f66-1d3d-4d2b-823b-8161b5ba9778)
![image](https://github.com/user-attachments/assets/7392fe27-a150-48a9-8b46-255537457c78)
![image](https://github.com/user-attachments/assets/f80a0db5-0489-4e48-955d-6722327bf976)

## Lab - Let's create couple of rocky linux ansible node containers
```
docker run -d --name rocky1 --hostname rocky1 -p 2003:22 -p 8003:80 tektutor/rocky-ansible-node:latest
docker run -d --name rocky2 --hostname rocky2 -p 2004:22 -p 8004:80 tektutor/rocky-ansible-node:latest
docker ps
```

Let's SSH into rocky1,rocky2 and see if it is allowing to login without asking for password
```
ssh -p 2003 root@localhost
exit

ssh -p 2004 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/c54e3fa7-0015-4e79-836a-24f31937abbb)
![image](https://github.com/user-attachments/assets/3fdbf589-0d6e-492b-9414-fd1d8bc8bed9)

## Info - Ansible configuration file
<pre>
- we can export the ANSIBLE_CONFIG environment variable to point the location and file name of ansible configuration file
- When the ANSIBLE_CONFIG environment is undefined, ansible or ansible-playbook will search for the ansible.cfg in the current directory, if it finds it will stop searching and it will use the ansible.cfg from the current directory
- in case, the ansible.cfg is not present in the current, ansible will try to locate a file name .ansible.cfg from home directory of the user, if it finds it will stop searching elsewhere and starts using the .ansible.cfg file
- if the home directory of the user doesn't have the file, ansible will finally try to locate at /etc/ansible/ansible.cfg file, it that is also missing, it will give up

- whichever place ansible.cfg file is located first, it will use it and stops searching for ansible.cfg elsewhere
</pre>

## Lab - Adding rocky container details in the hosts file
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible
cat ansible.cfg
cat hosts
ansible all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/f90a9761-51f5-40c8-9c2c-bde3ecdc51ab)


## Lab - Installing nginx via playbook on Ubuntu and Rocky linux servers
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/31aaa6a5-0daf-4d59-a23e-9ccedb1390b5)
![image](https://github.com/user-attachments/assets/cc5904f5-ff82-4536-b0a5-d17038e44ad6)
![image](https://github.com/user-attachments/assets/f58e843e-6855-40c6-9497-45a6eef94e1a)
![image](https://github.com/user-attachments/assets/70a7cbdc-b544-4d85-b0ae-c9185ad7800f)

## Info - Ansible facts
<pre>
- Ansible setup module is implicitly called as the very first task in every play in a playbook
- setup module collects loads of facts(meta-data) about the ansible node like 
  - what is the OS distribution, OS distribution version, python version installed, hardware details, etc., in case of unix/linux/mac
- what is the Windows OS, Window OS version, powershell version installed, hardware details, etc., in case of windows servers  
- we can use the facts to perform conditional installation via playbook
</pre>


## Info - Clean code practices and design principles
<pre>
- SOLID Design Principles
- S - Single Responsibility Principle (SRP)
- O - Open Closed Principle (OCP)
- L - Liskov Substituion Principle (LSP)
- I - Interface Seggregation 
- D - Dependency Injection or Dependency Inversion or Inversion of Control (IOC)
</pre>

As per Open Closed Principle, a good design will be open for extension without modifying existing code.

For example, let's say we wanted to add support for Fedora linux in the install-nginx-playbook.yml. Is it possible to add support to Fedora linux without modifying existing install-nginx-playbook.yml?

## Lab - Running the refactored install nginx ansible playbook
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible/after-refactoring
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/f8c9a7a6-3fe2-4a46-a433-9311e220803d)
![image](https://github.com/user-attachments/assets/90a38089-dbb6-492f-ba37-59b9f14c70aa)
![image](https://github.com/user-attachments/assets/ff6a36d4-ba89-47f9-abf6-e15954e9bd09)

## Info - Ansible Role
<pre>
- is a reusable code that can invoked from any ansible playbooks
- ansible role follows certain folder structure and best practices
- ansible roles can't executed independently
- ansible roles can only be accessed from ansible playbooks
- imagine ansible role like Dynamic Link Library, just like functions in DLL can't be executed directly ansible roles can't be executed directlry
- just like DLL functions can be invoked from application after loading the dll, from ansible playbook we can invoke ansible roles
- each ansible role will focus on one software
- For example
  - for installing nginx in ubuntu, rocky, fedora, rhel, suse we could write one role
  - for installing and configuring weblogic in multiple OS we could develop an ansible role
  - these ansible roles can later be accessed from any ansible playbooks
</pre>  

## Info - Ansible utilities
<pre>
ansible-doc - used to get help about any ansible module
ansible - to run ansible ad-hoc command
ansible-playbook - used to execute ansible playbooks
ansible-galaxy 
  - used to download,install reusable roles from galaxy.ansible.com website
  - also can be to develop a custom ansible role
</pre>

## Lab - Ansible role to install nginx in Ubuntu and Rocky Linux distributions
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible/nginx-role
ansible-galaxy init nginx
tree nginx
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/5e6ae941-4249-42ab-95f9-80d94e979c9f)
![image](https://github.com/user-attachments/assets/3dd93560-f2b6-41fb-8500-7d92ba3455d2)
![image](https://github.com/user-attachments/assets/e0ff355d-80c8-43a7-814a-e2fab1c33acf)
![image](https://github.com/user-attachments/assets/c33d2955-6eb9-4f2a-8301-5d48f8ae2008)

## Lab - Ansible Dynamic Inventory
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible/dynamic-inventory
docker ps
./inventory.py
cat ansible.cfg
ansible all -m ping
docker stop ubuntu2
ansible all -m ping
docker start ubuntu2
ansible all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/3cef1ea8-389a-4aac-9a51-1552b0ca1d76)
![image](https://github.com/user-attachments/assets/b7ac50ae-5c90-403e-bfff-484acb25c11e)
![image](https://github.com/user-attachments/assets/4945770c-8e53-4104-bfeb-e43a158715d5)
![image](https://github.com/user-attachments/assets/d485f268-e93d-4dc2-9a10-e45f140605d9)
![image](https://github.com/user-attachments/assets/462a571c-36c7-4756-abcf-df0893fde2f2)

## Info - Ansible variable precedence
<pre>
https://docs.ansible.com/ansible/latest/reference_appendices/general_precedence.html  
</pre>

## Lab - Suppressing ansible warnings
```
export ANSIBLE_PYTHON_INTERPRETER=auto_silent

cd ~/terraform-1721march-2025
git pull
cd Day2/ansible/dynamic-inventory

ansible all -m ping
```

## Lab - Ansible vault 
```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible/vault
cat playbook.yml
ansible-vault create mysql-login-credentials.yml
cat mysql-login-credentials.yml
ansible-vault view mysql-login-credentials.yml
cat mysql-login-credentials.yml
ansible-vault edit mysql-login-credentials.yml
cat mysql-login-credentials.yml
ansible-vault decrypt mysql-login-credentials.yml
cat mysql-login-credentials.yml

ansible-vault encrypt mysql-login-credentials.yml
cat mysql-login-credentials.yml


ansible-playbook playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/a8efe142-32bc-4b7f-9ee3-6be568746568)
![image](https://github.com/user-attachments/assets/e2b2639d-bcd6-49e7-a07c-df4add0526fc)
![image](https://github.com/user-attachments/assets/e73a2fcc-49e5-435c-a249-bc7f92c4acfa)
![image](https://github.com/user-attachments/assets/7c73fbde-39dc-42f9-96e4-83d07c1edc09)
