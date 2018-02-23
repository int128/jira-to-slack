FROM node:9
EXPOSE 3000

WORKDIR /app
COPY package.json /app
COPY package-lock.json /app
RUN ["npm", "install"]

COPY src /app/src

USER node
CMD ["npm", "start"]
