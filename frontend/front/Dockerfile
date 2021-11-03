FROM node:16.13-alpine3.14 as development

ENV NODE_ENV development

WORKDIR /app

COPY package.json .

COPY yarn.lock .

RUN yarn install

COPY . .

EXPOSE 3000

CMD ["yarn", "start"]