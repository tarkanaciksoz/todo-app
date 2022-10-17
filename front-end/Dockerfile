FROM node:lts-alpine as build-stage
ARG ENV
WORKDIR /app

RUN apk update

COPY package*.json ./
RUN npm install
COPY . .
RUN chmod +x docker/wait-for.sh
RUN chmod +x docker/test-and-build.sh

RUN npm run build-$ENV
FROM nginx:stable-alpine as production-stage
ARG ENV
RUN mkdir /app
RUN mkdir /docker
COPY --from=build-stage /app/docker /docker
COPY --from=build-stage /app /app
#COPY --from=build-stage /app/docker/wait-for.sh /app/wait-for.sh
#COPY --from=build-stage /app/docker/test-and-build.sh /app/test-and-build.sh
COPY --from=build-stage /app/nginx-$ENV.conf /etc/nginx/nginx.conf
#RUN chmod +x /app/wait-for.sh
#RUN chmod +x /app/test-and-build.sh
RUN apk update && apk add --update npm
EXPOSE 80
CMD sh docker/./wait-for.sh backend:9090 --timeout=30 -- sh docker/./test-and-build.sh $ENV
#CMD sh app/./wait-for.sh backend:9090 --timeout=30 -- sh app/./test-and-build.sh $ENV