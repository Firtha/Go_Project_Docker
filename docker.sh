#!/bin/bash

if [[ "$#" -ne "2" ]]; then
    echo "Wrong number of parameters, please follow :"
    echo "./docker.sh { deploy | destroy } { go | front }"
    echo "Running containers :"
    docker ps
else
    if [[ "$1" = "deploy" ]]; then
        if [[ "$2" = "go" ]]; then
            echo "Launching Go containers..."
            cd Dockerfile_Go
            docker-compose up -d
            cd ..
        elif [[ "$2" = "front" ]]; then
            echo "Launching Front containers..."
            cd Dockerfile_Front
            docker-compose up -d
            cd ..
        fi
    elif [[ "$1" = "destroy" ]]; then
        if [[ "$2" = "go" ]]; then
            echo "Stopping Go containers..."
            cd Dockerfile_Go
            docker-compose rm -svf
            cd ..
        elif [[ "$2" = "front" ]]; then
            echo "Stopping Front containers..."
            cd Dockerfile_Front
            docker-compose rm -svf
            cd ..
        fi
    fi
fi