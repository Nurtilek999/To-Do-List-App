FROM golang:1.19.4

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o Pet_1 ./main.go

CMD ["./Pet_1"]