FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"] 
