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
EXPOSE 6666
ENTRYPOINT ./main