#!/bin/sh

docker-compose build
docker-compose run app go test