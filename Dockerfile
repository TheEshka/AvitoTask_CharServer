# build stage
FROM golang:alpine AS build
COPY go.mod /build/
WORKDIR /build
RUN apk add git && \
    go mod download
COPY . /build/
RUN ls -alR && \
    go build -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=build /build/main /app/
RUN adduser -D -S -h /app appuser
USER appuser
ENV DATABASE_IP="127.0.0.1" \
 DATABASE_PASSW="mysecretpassword" \
 DATABASE_NAME="chat_db" \
 DATABASE_USER="postgres"

ENTRYPOINT ./main -db="user=$DATABASE_USER password=$DATABASE_PASSW host=$DATABASE_IP dbname=$DATABASE_NAME sslmode=disable" -port=":6666"