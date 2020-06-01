# Build the binary
FROM golang:1.14-alpine3.11 as build-env
RUN apk add git gcc
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ncrypt-api

# Production image
FROM alpine:3.11
COPY --from=build-env /app/ncrypt-api .
USER 1001
EXPOSE 1990
ENTRYPOINT ["./ncrypt-api"]