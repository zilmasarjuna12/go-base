#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64

#set env
export PORT="8087"

export MYSQL_HOST="localhost"
export MYSQL_PORT="3306"
export MYSQL_USER="root"
export MYSQL_PASSWORD="password"
export MYSQL_DATABASE="base"

go run main.go
