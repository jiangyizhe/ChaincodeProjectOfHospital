#!/bin/bash +x
#
cd fixtures
docker-compose down
docker volume ls -qf dangling=true
docker volume rm $(docker volume ls -qf dangling=true)
rm -rf /tmp/hospital-*
docker rm -f -v `docker ps -a --no-trunc | grep "hospital" | cut -d ' ' -f 1` 2>/dev/null
docker-compose up -d