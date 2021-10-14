#!/bin/sh

sudo docker network create bpk-network

sudo docker run -d -p 5432:5432 --net bpk-network --name bpk-pg \
    -e POSTGRES_DB=testdb \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    postgres

sudo docker build -t bpk-dx

sudo docker run -d -p 3000:3000 --net bpk-network --name bpk-c bpk-dx

--restart=always
sudo docker start `sudo docker ps -q -a` or `-l`