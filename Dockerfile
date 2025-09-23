FROM golang:1.23.4

WORKDIR /dopamine

COPY go.mod ./

RUN go mod download

COPY . .

RUN cd cmd/dopamine && go build

EXPOSE 8080

CMD ["./cmd/dopamine/dopamine"]