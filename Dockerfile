FROM golang:1.14.2-alpine as compiler
WORKDIR /app
COPY . /app
RUN go mod download && go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app
COPY --from=compiler /app/app .
CMD ["./app"]
