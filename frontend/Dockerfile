# syntax = docker/dockerfile:1.2

FROM node:22.18.0-alpine3.22 AS build

WORKDIR /app
RUN npm i -g pnpm@9.4.0

COPY package.json pnpm-lock.yaml  ./
RUN pnpm install --ignore-scripts --frozen-lockfile

COPY tsconfig.build.json tsconfig.build.json
COPY tsconfig.json tsconfig.json
COPY tsconfig.node.json tsconfig.node.json
COPY index.html index.html
COPY vite.config.ts vite.config.ts
COPY src src
COPY public public
RUN pnpm build-singleton


FROM nginx:1.29.1-alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /app


