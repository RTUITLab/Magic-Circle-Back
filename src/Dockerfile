FROM golang:1.17 AS builder

WORKDIR /go/src/Magic-Circle/

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/Magic-Circle/main .
CMD ["./main"]