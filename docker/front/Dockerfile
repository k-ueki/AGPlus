FROM node:latest

#ENV CONTAINER_PATH /var/www/docker-vue
#WORKDIR $CONTAINER_PATH

COPY . .
COPY package*.json /app
COPY yarn.lock /app

RUN apt-get update
RUN apt-get install -y yarn
RUN yarn global add @vue/cli
RUN yarn global add @vue/cli-service
RUN yarn global add axios

RUN yarn install

EXPOSE 8080

CMD ["yarn", "run", "serve"]
