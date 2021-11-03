FROM node:17-alpine as development

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./

COPY package-lock.json ./

RUN npm install && npm install react-scripts@3.4.1 -g

COPY . .

CMD ["npm", "start"]