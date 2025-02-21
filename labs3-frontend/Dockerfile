FROM node:latest AS builder

WORKDIR /app
COPY . .

# Define build arguments
ARG LAB3_URL

# Set environment variables inside the container
ENV LAB3_URL=${LAB3_URL}

RUN npm install && npm run build

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*

# Copy built files
COPY --from=builder /app/dist/spa .

# Copy custom NGINX config
COPY nginx.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]