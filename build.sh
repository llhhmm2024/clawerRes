#!/bin/bash

#build server
echo "build go crawler server"
go build -ldflags "-s -w" -o server main.go

docker compose build 

docker compose down
docker compose up -d

