# Day 2

## Lab - Understanding host and group variables in ansible inventory file
<pre>
- whatever variables are unique to an ansible node are called as host variable
- for example, the ansible_port variable is called a host variable as the ssh port on the local machine varies for each container
- the common variables for ubuntu1 and ubuntu2 are defined in the group level as they are common to ubuntu1 and ubuntu2. These variables are called group variables
- ansible_host, ansible_private_key_file, ansible_user are called group variables in this inventory
</pre>

```
cd ~/terraform-1721march-2025
git pull
cd Day2/ansible
cat hosts
ansible -i hosts all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/9ff10406-dc29-45d4-a696-84f4bac08854)
