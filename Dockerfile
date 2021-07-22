FROM golang:1.16 as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app 

FROM alpine:latest

ENV TZ=Asia/Bangkok

WORKDIR /app/

COPY --from=builder /app/app .

EXPOSE 3000

ENTRYPOINT ["./app"]