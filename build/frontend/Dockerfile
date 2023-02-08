FROM node:18-slim as dev

RUN mkdir frontend

WORKDIR /frontend

RUN apt-get update \
    && apt-get install -y \
    git \
    vim

EXPOSE 3000