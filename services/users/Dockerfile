FROM node:6.11.3-alpine

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY index.js .

ENTRYPOINT [ "node", "index.js" ]