#!/bin/bash
echo "`/sbin/ip route|awk '/default/ { print $3 }'` todo-app.localhost" >> /etc/hosts

cd app/ && npm run test:unit
mv ./dist ../dist
cd .. && rm -Rf app/
mv dist app
#mv dist ../ && cd .. && rm -rf app
#mkdir app/ && cp -R dist/ app/ && rm -rf dist/
#mv dist app
nginx -g "daemon off;"