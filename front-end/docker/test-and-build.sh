#!/bin/bash
echo "`/sbin/ip route|awk '/default/ { print $3 }'` todo-app.localhost" >> /etc/hosts

cd app/ && npm run test:unit
find . ! -name dist -maxdepth 1 -type f -delete
find . ! -name dist -maxdepth 1 -type d -delete
rm -rf docker node_modules public src tests
mv ./dist/* ./ && rm -rf ./dist
nginx -g "daemon off;"