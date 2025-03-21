# Day 5

## Lab - Developing our own custom Terraform file provider in Golang
Make sure the below folder is created
```
mkdir -p /home/rps/go/bin
```

Create file name .terraformrc under your home directory /home/rps/.terraformrc with the below content
<pre>
provider_installation {
  dev_overrides {
      "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",
  }
  direct {}
}  
</pre>

Now you may build your 
```
cd ~\terraform-1721march-2025
git pull
cd Day5/DevOpsCICDPipeline/custom-terraform-provider/file
tree
go mod init github.com/hashicorp/terraform-provider-file
go mod tidy
go build
go install
```

Expected output
![image](https://github.com/user-attachments/assets/564594be-5580-493f-9494-31cce3114ef5)
![image](https://github.com/user-attachments/assets/f54673af-4bee-4902-a507-25abdea6cc17)
![image](https://github.com/user-attachments/assets/97e393a0-fade-42de-8628-3293560a9742)
![image](https://github.com/user-attachments/assets/cf4395d1-c186-4e8b-9c24-b125352648f7)
![image](https://github.com/user-attachments/assets/a6e923ca-7817-4081-aa51-a6056514c866)

## Lab - Using our custom file provider in a terraform 
```
cd ~\terraform-1721march-2025
git pull
cd Day5/DevOpsCICDPipeline/test-terraform-provider-file-provider
ls -l
terraform plan
terraform apply --auto-approve
ls -l
```

Once you are done with this exercise, you may dispose the resources provisioned by Terraform
```
terraform destroy --auto-approve
ls -l
```

Expected output
![image](https://github.com/user-attachments/assets/9fb97d74-22e6-4aa1-a62f-9ace1c5182ca)
![image](https://github.com/user-attachments/assets/2a0d3791-ab59-4b0e-9f7f-3f90b8fcb977)
![image](https://github.com/user-attachments/assets/d55b6218-dab9-4ae9-83e7-d6079083fb61)
![image](https://github.com/user-attachments/assets/c0623ebd-d305-4b68-bf58-28e99dc975e9)
![image](https://github.com/user-attachments/assets/f606302a-3597-4032-8763-3fd0e8a32ba8)



## Info - Infrastructure automation test cases ( Sentinel Policy Management )
https://developer.hashicorp.com/terraform/tutorials/policy/sentinel-policy

## Info - Terraform Module
<pre>
- is a collection of many terraform scripts(*.tf) files in a dedicated folder
- modules encpasulate group of terraform resources related to a single infrastructure/task
- modules allows us to resue code
- the terraform scripts kept in the top-level folder is reerred as root module
- the terraform scripts kept in the top-level sub-folder is referred as child module
- a root module may have zero to many child modules
</pre>

## Info - Terraform Module vs Terraform Provider
<pre>
- Terraform Provider is developed by a Terraform Provider developer using golang programming language
- Terraform Module is written by DevOps Engineers who use the Terraform provider
- Terraform Module is written in HCL(Hashicorp Configuration Language) will file extensions *.tf
- Terraform Module is written in declarative code
- Terraform Modules are reusable code that can be invoked from Terraform Root modules
- Terraform Modules can be reused across many Terraform Projects
- Terraform Modules depends on one or more Terraform Providers
- Terraform Projects is nothing but the top level folder that has the below
  - Terraform Providers
  - Terraform Root Module
    - Terrafoorm child Module(s) optionally
</pre>

## Info - Terraform Provider Development best practices and recommended naming conventions
<pre>
- Provider name must be terraform-provider-nameoftheprovider, must be all lower case
- Resource names must start with nameoftheprovider-resource i.e docker_container, docker is the provider name while the resource managed is container
- resource and data source name must be all lowercase separated by underscore, and recommened to restrict to 2 or 3 words at the max
</pre>

## Info - Commons causes of Configuration Drift
<pre>
- manual changes
- inconsistent and manual deployments
- external dependencies
- difference in environments
- lack of version control
- poor or lack of documentation
</pre>

## Info - Risks associated with Configuration Drifts
<pre>
- Security Vulnerabilities
- Performance issues
- Compliance issues
- Makes troubleshooting very difficult
- increased down-time
- decreased reliability
- poor user-experience
</pre>

## Info - Solution to Configuration Drifts
<pre>
- Using Version Control
- Continuous Integration
- Use Configuration Management Tools to override manual changes in continuous fashion
- Use Infrastructure as Code Tools to override manual changes
- Regularly monitor and audit infrastructure
</pre>

## Info - Monitoring
<pre>
- is the process of collecting and analyzing data from IT infrastructure, system and processes
- using the data to improve business outcomes and drive value to the organization
- collects data to help keep your infrastructure up and running without any downtime
- Tools to
  - Data Collection (Logs)
  - Alerting
  - Remediation
  - Agent based monitoring
  - Agentless monitoring
  - Collecting Metrics
- Examples
  - Dynatrace
  - DataDog
  - Splunk
  - Prometheus & Grafana
</pre>

## Continuous Integration (CI) Overview
<pre>
- it is a fail-fast engineering process to find issues early 
- when bugs are detected early during development phase, they are easy to fix
- cost of fixing the bugs is also relatively cheaper
- it is similar SCRUM daily stand-up meeting, which is an inspect and adapt meeting
- the team that follows CI/CD, the engineers would be pushing code to version control several times a day
- code would be integrated many times a day
- Jenkins or similar CI/CD server will grab the latest code, they trigger a build, as part of the build, automated test cases would be executed to verify if the new code is as expected, if the new code is breaking any existing functionality.
- the build that was triggered due to new code integration succeeds, it means no functionality is broken, everything works as expected
- continous frequent feedback is given by CI/CD builds, eventually improving the code quaility and functional quality
- CI helps confidently deliver releases in a short amount of time
- Unit and Integration is the scope of CI
</pre>

## Continuous Deployment (CD) Overview
<pre>
- Once the dev release suceeds all the automated test cases added by dev team, it is automatically promoted for QA testing
- the dev certified release binaries are deployed automatically to QA environment for further automated QA testing
- the QA would automate, end to end functionality test, regression test, smoke test, performance test, stress test, component/API test, etc
</pre>

## Continuous Delivery (CD) Overview
<pre>
- the QA certified build(release) is automatically deployed into production or pre-prod environment
- in the pre-prod environment the customer or the Operations team would verify if the new release is working as expected
- especially fintech companies, after manual approval the binaries could go live in production environment
</pre>

## What is Jenkins?
- is a build automation server
- used mainly for CI/CD Builds
- this was developed in Java by Kohsuke Kawaguchi, former employee of Sun Microsystems
- Initially it was developed as Hudson is an opensource project
- Then Oracle acquired Sun Microsystems, then part of Hudson including Kohsuke Kawaguchi had quit Oracle
  created a new branch from Hudson called Jenkins
- The other part of the Hudson team they continue to develop the product as Hudson
- There is lot common code between Hudson and Jenkins
- More than 10000 active contributors are there for Jenkins
- Many other software vendors got inspired by Jenkins similar products came out in market like Bamboo, Team City, Microsoft TFS, etc.,
- Jenkins supports CI/CD build for products built in any software stack
  
## What is Cloudbees?
- Cloudbees is the enterprise paid variant of Jenkins
- Feature wise Jenkins and Cloudbees pretty much they are same
- We get support for Cloudbees while we only get community support for Jenkins
- Cloudbees is more stable as it is a paid version
  
## Jenkins Alternatives
- Bamboo
- Team City
- Cloudbees ( Enterprise Jenkins )
- Microsoft Team Foundation Server (TFS)

## Lab - Download Jenkins war file
Download the Generic Java Package (war) file from the left side (LTS)
<pre>
cd ~/Downloads
wget https://get.jenkins.io/war-stable/2.492.1/jenkins.war
</pre>

Expected output

## Lab - Launching Jenkins from terminal
```
cd ~/Downloads
java -jar ./jenkins.war
```

Expected output

## Lab - Accessing Jenkins Dashboard from your RPS Cloud machine chrome web browser
<pre>
http://localhost:8080  
</pre>
