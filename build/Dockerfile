FROM golang:1.16-alpine as builder

WORKDIR /build

EXPOSE 8080

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./ 

RUN CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o  go-s3-sqs ./cmd/.

FROM alpine:latest
WORKDIR /app 
COPY --from=builder /build/go-s3-sqs .

CMD ["/app/go-s3-sqs"]