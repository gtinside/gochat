# gochat
![Docker Image CI](https://github.com/gtinside/gochat/workflows/Docker%20Image%20CI/badge.svg?branch=master)

GoChat is a web based chat application designed using Golang

#### Core Components
<hr/>

1. GoChat Server
2. [Web Server](https://github.com/gtinside/gochat-client)

This repository is for GoChat Server. To read more about the architecture, visit my blog [GoChat Architecture](https://gauravtiwari.blog/2020/05/18/gochat-yet-another-chat-application)

#### Requirements
<hr/>

1. go1.14.2 darwin/amd64
2. Docker
3. Tested on Mac OSX, Ubuntu 14.X, CentOS 6.X
4. Node.js Package Manager
5. [AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)

#### Getting Started
<hr/>

##### Before Build
If you don't have requirements prepared, install it. The installation method may vary according to your environment. Refer to the documentation accordingly.

##### Build
```
git clone https://github.com/gtinside/gochat
docker build. --file Dockerfile --tag gochat-server:latest
docker run docker run -p 8090:8090  gochat-server:latest
```
To compile and run locally, execute the following from within project directory to start the server on **port 8090**: 
```
go build -o bin/ ./cmd/...
./bin/chatserver
```
##### Start DynamoDB locally
After the server startup, navigate to **gochat/test/integration/** directory and run the following commands:
```
docker-compose build
./scripts/setup.sh
 ```
The setup script does the following:
1. Starts Dynamodb
2. Create UserDetails and MessageDetails tables. Refer to the [architecture page](https://gauravtiwari.blog/2020/05/18/gochat-yet-another-chat-application) for more details.

#### Packaging & Deployment
<hr/>

Refer to the [Workflow](https://github.com/gtinside/gochat/blob/master/.github/workflows/dockerimage.yml) file for build and deployment details. 
Following workflows are embedded in it:
1. Image build and publish to Github package registry
2. Image build and publish to [Docker public registry](https://hub.docker.com/repository/docker/gtinside/gochat-server)
3. Refresh of AWS ECS Service. I have hardcoded the ECS Cluster and Deployment Service name for now
