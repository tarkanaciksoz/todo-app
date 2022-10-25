FROM node:lts-alpine as build-stage
ARG ENV
WORKDIR /app

RUN apk update

COPY package*.json ./
RUN npm install
COPY . .

RUN npm run build-$ENV
FROM nginx:stable-alpine as production-stage
ARG ENV
RUN mkdir /app
COPY --from=build-stage /app /app
COPY --from=build-stage /app/nginx-$ENV.conf /etc/nginx/nginx.conf

RUN apk add --update nodejs npm
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]