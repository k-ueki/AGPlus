FROM node:latest

COPY package*.json /app
COPY yarn.lock /app

RUN apt-get update
RUN apt-get install -y yarn 
RUN yarn global add @vue/cli
RUN yarn global add @vue/cli-service

RUN yarn install
