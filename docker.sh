#!/bin/sh

docker network create bpk-network

docker run -d -p 5432:5432 --net bpk-network --name bpk-pg \
    -e POSTGRES_DB=testdb \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    postgres

docker build -t bpk-dx

docker run -d -p 3000:3000 --net bpk-network --name bpk-c bpk-dx