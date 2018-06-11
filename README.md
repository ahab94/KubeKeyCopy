# AddMyKeyPlease

Add your key anywhere, anytime, quickly on your docker, kubernetes, Openshift deployments with one hit.
# Usage
 
 - Clone locally `$ git clone https://github.com/ahab94/AddMyKeyPlease.git`
 - Change directory `$ cd AddMyKeyPlease`
 
## Docker 
 Make sure docker, docker-compose is installed on the server.
 - Swap your base64 SSH public key value in `docker-compose.yaml` for KEY enviornment variable. 
 - Build and spawn your containers `$ docker-compose up --build`
 Or
 Simply spawn containers by doing...
 - `$ docker run -e KEY=base64Key -it ahab94/addmykeyplease:latest`
 
## Kubernetes 
 Make sure kubectl is installed on the server.
 On master node...
 - Swap your base64 SSH public key value in `kubernetes/AddMyKeyDS.yaml` for KEY enviornment variable.
 - Run daemonset by doing `$ kubectl create -f kubernetes/AddMyKeyDS.yaml`.
 - It will copy ssh key and validate after a certain duration that the public key exists. 

Any Quetion or Unexpected behavior? Please report [here](https://github.com/ahab94/AddMyKeyPlease/issues).
