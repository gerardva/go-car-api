FROM golang:1.20 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

ADD app.env .

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app

RUN apk update && apk add bash

COPY --from=builder /build/main .
COPY --from=builder /build/app.env /etc/

EXPOSE 8080
ENTRYPOINT ["/app/main"]
# ENTRYPOINT ["tail", "-f", "/dev/null"]