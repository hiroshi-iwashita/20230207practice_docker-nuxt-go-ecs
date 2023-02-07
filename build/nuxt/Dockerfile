FROM node:18-slim as dev

RUN mkdir src

RUN apt-get update \
    && apt-get install -y \
    git \
    vim

COPY . .

EXPOSE 3000