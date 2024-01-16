FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go build -o main .

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"] 
