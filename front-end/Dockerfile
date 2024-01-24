# Stage 1: Build the React application
FROM node:14 as build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# Stage 2: Serve the application from Nginx
FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf

ENV HOST 0.0.0.0
EXPOSE 80

CMD sh -c "nginx -g 'daemon off;'"
