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


