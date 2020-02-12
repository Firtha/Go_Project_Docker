# Build command line :
# docker build -t node/custom .
FROM node:10

WORKDIR /app

COPY package.json package.json
RUN npm install --quiet


COPY . .
RUN npm install -g nodemon
CMD [ "nodemon", "index.js" ]
RUN npm install -g mongoose

RUN chmod 775 *
