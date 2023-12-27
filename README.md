# About
In this project I've tried setting up a CI/CD pipeline. I've created a basic web app which exposes an api(login), whenever there'a a request, it sends a name and details in json format.
Here'a a list of tools which I used:
- Docker: For contenatization of applications.
- Minikube: For setting up local K8s cluster.
- AWS EC2: For setting up a Jenkins server.
- Jenkins: For setting up CI part.
- ArgoCD: For the CD part following GitOps principles.

## Step 1
Containerize the app using Dockerfile given in the project repo.
```
docker build -t prady0t/pipeline .
```
Now try to run it.
```
docker run -it prady0t/pipeline
```
App should be running at port 3001.

## Step 2
Push the image to the Dockerhub registery. (if not authenticated, first do docker login)
```
docker push prady0t/pipeline
```
## Step 3
Setup a Jenkins server on AWS EC2. Set inbould traffic to `Anywhere`

<img width="1325" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/c440ae5a-78ff-492f-8a8e-c4c07b02598a">


Login to your AWS account and create a new EC2 instance. Now follow these steps to setup Jenkins on EC2 instance : https://www.jenkins.io/doc/book/installing/linux/#debianubuntu.
(You would have to install Java first). Now in the terminal
```
systemctl status jenkins
```
Check Jenkins should be in the running state.
Copy paste IP of the instance followed by :8080 (default port at which jenkins run) in your browser. You should have a unlock Jenkins page.

<img width="1013" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/3241fba5-ba29-4f93-8c29-3759ac401dd9">

Now go to your instance's terminal and do :
```
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```
Use this as a password to unlock Jenkins. Download suggested pluggins.

## Step 4
Setup credentials for Github (using secret text, Github token as secret) and Dockerhub (using username and password).

<img width="1213" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/873afeaa-394b-4526-a25b-67da918363be">


Create a pipeline project in Jenkins.Select these as per given below:

<img width="578" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/19b11042-c636-46ec-b039-d9b1da9810ff">

<img width="1079" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/94fe37d3-59d9-407d-b6f0-c07c253ff7b2">

<img width="945" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/2e85b123-0c01-497e-abf6-275b994667a1">

Save

## Step 5
Setting up docker on AWS instance. Why do we need that?(I learned it the hard way) We are running single machine where Jenkins master and slave are same machine. 
We want nodes to have docker installed so taht we can build images with Jenkins. We would also need a `Docker pipeline` plugin.
Go to aws console and type these commands:

```
sudo apt-get install docker.io -y
sudo systemctl start docker
```
Check status:
```
systemctl status docker
```
Now we have to give jenkins user and ubuntu user permission to acces docker daemon:
```
sudo su -
usermod -aG docker jenkins
usermod -aG docker ubuntu
systemctl restart docker
```
Also restart jenkins.

Now add webhook to github repo. A webhook is basically Jenkins server url followed by `//github-webhook/`:

![image (4)](https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/a09ef3ae-0518-40b3-be73-6583ca8ca9a5)


Now hit `build` on Jenkins:

<img width="1116" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/1dca6396-2fe2-4c74-b60d-33f87dc74279">

### A build is triggered whenever a commit is pushed to main! 

With this we are done with our CI part.

## Step 6

Setup ArgoCD cluster on minikube. Just follow all the steps from the docs -> https://argo-cd.readthedocs.io/en/stable/getting_started/
Go to `localhost:8080`

<img width="1380" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/dfad6712-770c-4a3e-bed1-9622a072e0be">

Again for initial login, follow the docs.

## Step 7

Create a new app with these settings:

<img width="1057" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/9800692d-926d-4aec-ba99-fe9236b2b15b">

<img width="1057" alt="Screenshot 2023-12-28 at 2 11 02 AM" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/d1c21a6f-39b2-4c8b-935b-ade6d271c56f">

Click `Create`

<img width="1323" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/68b0efef-b293-49ca-a2b7-5f87c4b03402">

Go to the terminal and enter:
```
kubectl get pods
```
You will see it automatically created pods which were mentioned in the manifest files!

<img width="1017" alt="image" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/608d7e1f-e5c6-4653-bf3f-0ca6fbdf761c">

With this we are also done with our CD part.
