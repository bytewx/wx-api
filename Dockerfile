FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata ca-certificates

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/wx-api cmd/wx-api/main.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/wx-api /app/wx-api
COPY --from=builder /build/config/config.yaml /app/config/config.yaml

EXPOSE 8080

CMD ["./wx-api"]