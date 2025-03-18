# Day 2

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
ansible -i hosts -m shell -a "service nginx status"
```
![image](https://github.com/user-attachments/assets/eb06b7d6-3c48-49f9-8b33-6f4e6c230e0f)
