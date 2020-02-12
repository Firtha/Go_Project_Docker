# Go_Project_Docker
Docker repo of the project

### Host machine
1. Docker : 19.03.1 (https://docs.docker.com/docker-for-windows/release-notes/)
2. Docker Engine : > 18.xx
3. Docker-compose : 1.24.1 (see https://github.com/docker/compose/releases/)

## Instructions

### General
A script has been writen in order to manage easily deployment and destruction of the containers used (docker.sh). Instructions are given if you call the script with more or less than 2 arguments, calling syntax is:
```# ./docker.sh deploy|destroy go|front```


### Front
1. Build the image from the working folder Dockerfile_Front :
```# docker build -t node/custom .```
2. From the same folder, mount the container from the builded image node/custom:
```# docker-compose up -d```

### Go
The container for hosting a Geth is finally not used, the files are here only to keep track of the work.
The solution used is a call to infura's API for the test Ethereum network : Rinkeby.
More details on the Go repository.
