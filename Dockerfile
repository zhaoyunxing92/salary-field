FROM golang:1.17-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN export GOPROXY=https://proxy.golang.com.cn,direct
RUN go build -o field app.go

# run
FROM alpine:3.15.4 AS runner
WORKDIR /field

COPY --from=builder /app/field ./field
RUN chmod 775 field

EXPOSE 9090

ENTRYPOINT ["./field"]




