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

