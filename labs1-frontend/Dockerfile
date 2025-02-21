FROM node:20 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . . 
RUN npm run build

FROM node:20 AS runner

WORKDIR /app

RUN npm install -g serve

COPY --from=builder /app/dist ./

EXPOSE 81

CMD ["serve", "-s", ".", "-l", "81"]