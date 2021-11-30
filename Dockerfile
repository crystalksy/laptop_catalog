# FROM golang:1.17-alpine as build

# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o app

# FROM alpine:latest
# COPY --from=build /app/app /app/

# EXPOSE 8080
# CMD ["/app"]

FROM golang:1.17-alpine3.14

WORKDIR /laptop_catalog

COPY . .

RUN go mod download


RUN go build -o mainfile

EXPOSE 8080

CMD ["./mainfile"]