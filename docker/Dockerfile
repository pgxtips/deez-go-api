FROM alpine:latest

# Installing go

RUN apk add --no-cache go

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading
# them in subsequent builds if they change

COPY ./app/go.mod ./app/go.sum ./
RUN go mod download && go mod verify

COPY ./app .
RUN go build -v -o /usr/local/bin/app ./cmd/deez-go-api.go


# Installing nginx
RUN apk --no-cache add nginx

# Copy nginx config
COPY ./docker/nginx/remote/site.conf /etc/nginx/conf.d/site.conf 
COPY ./docker/nginx/remote/site.conf /etc/nginx/sites-enabled/default
COPY ./docker/nginx/nginx.conf /etc/nginx

# Copy entrypoint script and give privilege to execute 
COPY ./docker/entrypoint.sh /var/app/entrypoint.sh
RUN chmod -R 777 /var/app/entrypoint.sh

# Exposing nginx port
EXPOSE 80

# Invoke entrypoint.sh 
CMD ["/var/app/entrypoint.sh"]
