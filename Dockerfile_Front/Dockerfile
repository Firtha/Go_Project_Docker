# Build command line :
# docker build -t node/custom .
FROM node:7.7.2-alpine

WORKDIR /usr/app

COPY package.json .
RUN npm install --quiet

COPY app.js .
COPY index.html .

RUN chmod 775 *