#!/bin/bash
echo "`/sbin/ip route|awk '/default/ { print $3 }'` todo-app.localhost" >> /etc/hosts

cd app/ #&& npm run test:unit
mv ./dist ../dist
cd .. && rm -Rf app/
mv dist app
nginx -g "daemon off;"