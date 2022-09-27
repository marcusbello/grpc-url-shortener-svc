FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN cd server && go build -o ./run .

FROM alpine:latest

#RUN apk --no-cache add ca-certificates
WORKDIR /root/


#Copy executable from builder
COPY --from=builder /go/src/app/server/run .


EXPOSE 50052

CMD ["./run", "port", "50052"]