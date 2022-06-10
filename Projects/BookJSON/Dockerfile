FROM golang:1.18-alpine as builder

WORKDIR /bookjson

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /bookjson/myapp .

FROM alpine:latest
COPY --from=builder /bookjson/myapp /bookjson/myapp

EXPOSE 8080
ENTRYPOINT [ "/bookjson/myapp" ]