FROM node:22-alpine AS build

WORKDIR /app

COPY . .

RUN corepack enable

RUN pnpm install
RUN pnpm build

FROM halverneus/static-file-server

COPY --from=build /app/dist /static

ENV PORT=80
ENV FOLDER=/static

EXPOSE 80
