FROM golang:alpine AS builder
WORKDIR /opt/src/
RUN apk update && apk add --no-cache git
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o magiccap-api

FROM scratch
COPY --from=builder /opt/src/magiccap-api /magiccap-api
ENTRYPOINT [ "/magiccap-api" ]