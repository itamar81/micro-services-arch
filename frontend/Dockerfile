#docker buildx build . --tag ghcr.io/itamar81/micro-services-arch/frontend:1.1 --push --platform linux/amd64,linux/arm64
FROM node:20.5 as builder
WORKDIR /app
COPY package*.json ./
COPY . .
RUN npm ci --omit=dev

FROM node:20.5.0-alpine3.18
WORKDIR /app
COPY --from=builder /app .

CMD [ "node" , "app.js" ]